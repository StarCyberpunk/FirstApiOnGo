package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

type Bank_account struct {
	ID         uuid.UUID `json:"id_ba"`
	PassSerial int       `json:"pass_serial"`
	PassNumber int       `json:"pass_number"`
	CashTotal  float64   `json:"cash_total"`
	IdUser     uuid.UUID `json:"id_user"`
}
type BankAccountRepository interface {
	CreateBankAccount(ctx context.Context, bank_account Bank_account) (uuid.UUID, error)
	ReadBankAccount(ctx context.Context, id_ba uuid.UUID) (*Bank_account, error)
}
