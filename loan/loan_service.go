package loan

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/go-gin-example/utils"
)

type Service struct {
	*sql.DB
}

func NewLoanService(db *sql.DB) Repository {
	return Service{db}
}

func (p Service) Store(req CreateLoanRequest) (*Loan, error) {
	status := utils.RandomizeStatus()
	province := strings.ToUpper(req.AddressProvince)
	gender := strings.ToUpper(req.Gender)
	nationality := strings.ToUpper(req.Nationality)
	row := p.DB.QueryRow(StoreLoan,
		req.FullName,
		gender,
		req.KTPNumber,
		req.ImageOfKTP,
		req.ImageOfSelfie,
		req.DateOfBirth,
		req.Address,
		province,
		req.PhoneNumber,
		req.Email,
		nationality,
		req.LoanAmount,
		req.Tenor,
		status,
		time.Now())

	var id int64
	err := row.Scan(&id)
	//select the last inserted id and get the data
	loan, err := p.findLoanById(id)
	if err != nil {
		return nil, err
	}
	//insert into installment if the status is accepted
	if loan.Status == utils.Accepted {
		installment, err := p.insertInstallment(id, req)
		if err != nil {
			return nil, err
		}
		loan.Installment = installment
		return loan, nil
	}
	return loan, nil
}

func (p Service) List() ([]*Loan, error) {
	rows, err := p.DB.Query(ListLoan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*Loan
	for rows.Next() {
		var s Scan
		err = rows.Scan(
			&s.Id,
			&s.FullName,
			&s.Gender,
			&s.KTPNumber,
			&s.ImageOfKTP,
			&s.ImageOfSelfie,
			&s.DateOfBirth,
			&s.Address,
			&s.AddressProvince,
			&s.PhoneNumber,
			&s.Email,
			&s.Nationality,
			&s.LoanAmount,
			&s.Tenor,
			&s.Status,
			&s.CreatedAt)
		if err != nil {
			return nil, err
		}
		data := &Loan{}
		data = data.FromScan(s)
		installment, _ := p.getInstallment(data.Id)
		data.Installment = installment
		items = append(items, data)
	}
	return items, nil
}

func (p Service) isSurpassLimit() bool {
	row := p.DB.QueryRow(CountLoan)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	if count >= utils.MaxLimitLoanProcessPerDay {
		return true
	}
	return false
}

func (p Service) IsKtpNumberExist(ktpNumber string) bool {
	row := p.DB.QueryRow(getLoansByKtpNumber, ktpNumber)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

func (p Service) findLoanById(id int64) (*Loan, error) {
	row := p.DB.QueryRow(getById, id)
	var s Scan
	err := row.Scan(
		&s.Id,
		&s.FullName,
		&s.Gender,
		&s.KTPNumber,
		&s.ImageOfKTP,
		&s.ImageOfSelfie,
		&s.DateOfBirth,
		&s.Address,
		&s.AddressProvince,
		&s.PhoneNumber,
		&s.Email,
		&s.Nationality,
		&s.LoanAmount,
		&s.Tenor,
		&s.Status,
		&s.CreatedAt)
	if err != nil {
		return nil, err
	}
	data := &Loan{}
	data = data.FromScan(s)
	return data, nil
}

func (p Service) insertInstallment(idLoans int64, req CreateLoanRequest) (*Installment, error) {
	intLoanAmount, _ := strconv.Atoi(req.LoanAmount)
	intTenor, _ := strconv.Atoi(req.Tenor)
	paymentAmount := intLoanAmount / intTenor
	paymentAmountFloat := float64(paymentAmount)
	paymentAmountFinalize := paymentAmountFloat + (paymentAmountFloat * utils.InterestRate)
	paymentAmountFinalizeString := strconv.FormatFloat(paymentAmountFinalize, 'f', 2, 64)
	row := p.DB.QueryRow(StoreInstallment,
		idLoans,
		false,
		req.Tenor,
		req.Tenor,
		paymentAmountFinalizeString,
		req.LoanAmount,
		time.Now())
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	installment, err := p.getInstallment(idLoans)
	if err != nil {
		return nil, err
	}
	return installment, nil
}

func (p Service) getInstallment(loanId int64) (*Installment, error) {
	rows := p.DB.QueryRow(getInstallmentListByLoanId, loanId)
	var s ScanInstallment
	err := rows.Scan(
		&s.Id,
		&s.LoanId,
		&s.AllPaidOff,
		&s.TenorRemaining,
		&s.TotalTenor,
		&s.InstallmentAmount,
		&s.TotalInstallmentAmount,
		&s.CreatedAt)
	if err != nil {
		return nil, err
	}
	data := &Installment{}
	data = data.FromScan(s)
	return data, nil
}
