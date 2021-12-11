package loan

const StoreLoan = `
	INSERT INTO 
		loans(
			full_name, 
			gender, 
			ktp_number, 
			image_of_ktp, 
			image_of_selfie, 
			date_of_birth, 
			address, 
			address_province, 
			phone_number, 
			email,
			nationality,
			loan_amount, 
			tenor,
			status,
			created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id;`

const ListLoan = `SELECT 
		id, 
		full_name, 
		gender, 
		ktp_number, 
		image_of_ktp, 
		image_of_selfie, 
		date_of_birth, 
		address, 
		address_province, 
		phone_number, 
		email,
		nationality,
		loan_amount, 
		tenor,
		status,
		created_at
	FROM loans;`

const getById = `SELECT 
		id, 
		full_name, 
		gender, 
		ktp_number, 
		image_of_ktp, 
		image_of_selfie, 
		date_of_birth, 
		address, 
		address_province, 
		phone_number, 
		email,
		nationality,
		loan_amount, 
		tenor,
		status,
		created_at
	FROM loans lp 
	WHERE lp.id=$1;`

const getLoansByKtpNumber = `SELECT count(*) FROM loans WHERE loans.ktp_number=$1;`

const CountLoan = `SELECT count(*) FROM loans;`

const StoreInstallment = `
	INSERT INTO 
		payment_installments(
			loan_id, 
			all_paid_off, 
			tenor_remaining, 
			total_tenor, 
			installment_amount, 
			total_installment_amount, 
			created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

const getInstallmentListByLoanId = `
	SELECT 
		id, 
		loan_id, 
		all_paid_off, 
		tenor_remaining, 
		total_tenor, 
		installment_amount, 
		total_installment_amount,
		created_at
	FROM payment_installments pi 
	WHERE pi.loan_id=$1;`
