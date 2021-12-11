package loan

import (
	"database/sql"
	"time"
)

type Loan struct {
	Id              int64        `json:"id"`
	FullName        string       `json:"full_name"`
	Gender          string       `json:"gender"`
	KTPNumber       string       `json:"ktp_number"`
	ImageOfKTP      string       `json:"image_of_ktp"`
	ImageOfSelfie   string       `json:"image_of_selfie"`
	DateOfBirth     string       `json:"date_of_birth"`
	Address         string       `json:"address"`
	AddressProvince string       `json:"address_province"`
	PhoneNumber     string       `json:"phone_number"`
	Email           string       `json:"email"`
	Nationality     string       `json:"nationality"`
	LoanAmount      string       `json:"loan_amount"`
	Status          string       `json:"status"`
	Tenor           string       `json:"tenor"`
	CreatedAt       time.Time    `json:"created_at"`
	Installment     *Installment `json:"installment"`
}

type Installment struct {
	Id                     int64     `json:"id"`
	LoanId                 int64     `json:"loan_id"`
	AllPaidOff             bool      `json:"all_paid_off"`
	TenorRemaining         string    `json:"tenor_remaining"`
	TotalTenor             string    `json:"total_tenor"`
	InstallmentAmount      string    `json:"installment_amount"`
	TotalInstallmentAmount string    `json:"total_installment_amount"`
	CreatedAt              time.Time `json:"created_at"`
}

type Scan struct {
	Id              sql.NullInt64
	FullName        sql.NullString
	Gender          sql.NullString
	KTPNumber       sql.NullString
	ImageOfKTP      sql.NullString
	ImageOfSelfie   sql.NullString
	DateOfBirth     sql.NullString
	Address         sql.NullString
	AddressProvince sql.NullString
	PhoneNumber     sql.NullString
	Email           sql.NullString
	Nationality     sql.NullString
	LoanAmount      sql.NullString
	Status          sql.NullString
	Tenor           sql.NullString
	CreatedAt       sql.NullTime
}

type ScanInstallment struct {
	Id                     sql.NullInt64
	LoanId                 sql.NullInt64
	AllPaidOff             sql.NullBool
	TenorRemaining         sql.NullString
	TotalTenor             sql.NullString
	InstallmentAmount      sql.NullString
	TotalInstallmentAmount sql.NullString
	CreatedAt              sql.NullTime
}

func (y *Loan) FromScan(s Scan) *Loan {
	y.Id = s.Id.Int64
	y.FullName = s.FullName.String
	y.Gender = s.Gender.String
	y.KTPNumber = s.KTPNumber.String
	y.ImageOfKTP = s.ImageOfKTP.String
	y.ImageOfSelfie = s.ImageOfSelfie.String
	y.DateOfBirth = s.DateOfBirth.String
	y.Address = s.Address.String
	y.AddressProvince = s.AddressProvince.String
	y.PhoneNumber = s.PhoneNumber.String
	y.Email = s.Email.String
	y.Nationality = s.Nationality.String
	y.LoanAmount = s.LoanAmount.String
	y.Status = s.Status.String
	y.Tenor = s.Tenor.String
	y.CreatedAt = s.CreatedAt.Time
	return y
}

func (w *Installment) FromScan(s ScanInstallment) *Installment {
	w.Id = s.Id.Int64
	w.LoanId = s.LoanId.Int64
	w.AllPaidOff = s.AllPaidOff.Bool
	w.TenorRemaining = s.TenorRemaining.String
	w.TotalTenor = s.TotalTenor.String
	w.InstallmentAmount = s.InstallmentAmount.String
	w.TotalInstallmentAmount = s.TotalInstallmentAmount.String
	w.CreatedAt = s.CreatedAt.Time
	return w
}
