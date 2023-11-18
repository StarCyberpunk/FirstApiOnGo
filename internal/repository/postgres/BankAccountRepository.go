package postgres

import (
	"awesomeProject1/internal/domain"
	"database/sql"
	"github.com/gofrs/uuid"
	"log"
)

type BankAccountRepository struct {
	db_con *sql.DB
}

func NewBankAccountRepository(db_co *sql.DB) *BankAccountRepository {
	return &BankAccountRepository{
		db_con: db_co,
	}
}

func (repostitory *BankAccountRepository) CreateBankAccount(ba domain.Bank_account) (uuid.UUID, error) {
	_, err := repostitory.db_con.Query("INSERT INTO bank.bank_account( id_ba,pass_serial, pass_number, cash_total) VALUES ( $1, $2, $3,$4);", ba.ID, ba.PassSerial, ba.PassNumber, 0)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return ba.ID, err
}
func (repostitory *BankAccountRepository) ReadBankAccount(id_ba uuid.UUID) (domain.Bank_account, error) {
	rows, err := repostitory.db_con.Query("select * from bank.bank_account where id_ba = $1;", id_ba)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	var ba domain.Bank_account
	for rows.Next() {
		rows.Scan(&ba)
	}
	return ba, err
}
