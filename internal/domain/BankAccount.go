package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

type Bank_account struct {
	ID         uuid.UUID `json:"id_ba"`
	Name       string    `json:"name"`
	Balance    int       `json:"balance"`
	PassSerial int       `json:"pass_serial"`
	PassNumber int       `json:"pass_number"`
	CashTotal  float64   `json:"cash_total"`
	IdUser     uuid.UUID `json:"id_user"`
}
type BankAccountRepository interface {
	CreateBankAccount(ctx context.Context, bank_account Bank_account) (uuid.UUID, error)
	FindBankAccountById(ctx context.Context, id_ba uuid.UUID) (*Bank_account, error)
	FindBankAccountByNameId(ctx context.Context, id_ba uuid.UUID, name string) (*Bank_account, error)
	FindByIDForUpdate(ctx context.Context, id uuid.UUID) (*Bank_account, error)
	UpdateAccountBalance(ctx context.Context, id uuid.UUID, deposit int) error
	FindUserAccountsILike(ctx context.Context, name string, offset int, limit int, userID uuid.UUID) ([]Bank_account, error)
}
