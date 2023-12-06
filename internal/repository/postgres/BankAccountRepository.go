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
	_, err := repostitory.db_con.Exec("INSERT INTO bank.bank_account( id_ba,pass_serial, pass_number, cash_total,id_user) VALUES ( $1, $2, $3,$4,$5);", ba.ID, ba.PassSerial, ba.PassNumber, 0, ba.IdUser)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return ba.ID, err
}
func (repostitory *BankAccountRepository) ReadBankAccount(id_ba uuid.UUID) (domain.Bank_account, error) {
	rows := repostitory.db_con.QueryRow("select * from bank.bank_account where id_ba = $1;", id_ba)
	err := rows.Err()
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	var ba domain.Bank_account
	rows.Scan(&ba)
	return ba, err
}
