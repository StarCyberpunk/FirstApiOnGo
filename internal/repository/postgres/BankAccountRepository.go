package postgres

import (
	"awesomeProject1/internal/domain"
	"database/sql"
	"github.com/gofrs/uuid"
)

type BankAccountRepository struct {
	db_con *sql.DB
}

func NewBankAccountRepository(db_co *sql.DB) *BankAccountRepository {
	return &BankAccountRepository{
		db_con: db_co,
	}
}

func (repostitory *BankAccountRepository) CreateBankAccount(ba domain.Bank_account) {
	repostitory.db_con.Query("INSERT INTO public.bank_account( pass_serial, pass_number, cash_total) VALUES ( $1, $2, $3);", ba.PassSerial, ba.PassNumber, 0)
}
func (repostitory *BankAccountRepository) ReadBankAccount(baid uuid.UUID) {
	repostitory.db_con.Query("select * from public.bank_account where id = $1;", baid)
}
