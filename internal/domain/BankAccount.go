package domain

import "github.com/gofrs/uuid"

type Bank_account struct {
	ID         uuid.UUID `json:"id_ba"`
	PassSerial int       `json:"pass_serial"`
	PassNumber int       `json:"pass_number"`
	CashTotal  float64   `json:"cash_total"`
}
type BankAccountRepository interface {
	CreateBankAccount(bank_account Bank_account) (uuid.UUID, error)
	ReadBankAccount(id_ba uuid.UUID) (Bank_account, error)
}
