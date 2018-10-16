package model

type Account struct {
	// Ten bang
	TableName struct{}	`json:"table_name" sql:"bank.accounts"`

	// Ma account (chuoi ngau nhien duy nhat)
	Id string `json:"id"`

	// Ten hien thi
	AccountName string `json:"account_name" valid:"required~Ten tai khoan khong duoc de trong, runelength(6|100)"`

	// So tien
	Amount int64 `json:"amount"`
}

type CreateAccount struct {
	// Ma account (chuoi ngau nhien duy nhat)
	Id string `json:"id"`

	// Ten hien thi
	AccountName string `json:"account_name" valid:"required~Ten tai khoan khong duoc de trong, runelength(6|100)"`

	// So tien
	Amount int64 `json:"amount"`
} 

type GetAmount struct {
	// Ma account (chuoi ngau nhien duy nhat)
	Id string `json:"id"`

	// Ten hien thi
	AccountName string `json:"account_name" valid:"required~Ten tai khoan khong duoc de trong, runelength(6|100)"`

	// So tien
	Amount int64 `json:"amount"`
}

type AddFund struct {
	// Ma account (chuoi ngau nhien duy nhat)
	Id string `json:"id"`

	// Ten hien thi
	AccountName string `json:"account_name" valid:"required~Ten tai khoan khong duoc de trong, runelength(6|100)"`

	// So tien
	Amount int64 `json:"amount"`
}

type WithDrawMoney struct {
	// Ma account (chuoi ngau nhien duy nhat)
	Id string `json:"id"`

	// Ten hien thi
	AccountName string `json:"account_name" valid:"required~Ten tai khoan khong duoc de trong, runelength(6|100)"`

	// So tien
	Amount int64 `json:"amount"`
}


