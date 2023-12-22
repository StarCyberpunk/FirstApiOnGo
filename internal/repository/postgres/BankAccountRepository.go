package postgres

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/internal/pkg/persistence"
	"context"
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"log"
)

type BankAccountRepository struct {
	db_con persistence.Connection
}

func NewBankAccountRepository(connection persistence.Connection) *BankAccountRepository {
	return &BankAccountRepository{
		db_con: connection,
	}
}

func (repostitory *BankAccountRepository) CreateBankAccount(ctx context.Context, ba domain.Bank_account) (uuid.UUID, error) {
	_, err := repostitory.db_con.Exec(ctx, "INSERT INTO bank.bank_account( id_ba,pass_serial, pass_number, cash_total,id_user) VALUES ( $1, $2, $3,$4,$5);", ba.ID, ba.PassSerial, ba.PassNumber, 0, ba.IdUser)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return ba.ID, err
}
func (repostitory *BankAccountRepository) ReadBankAccount(ctx context.Context, id_ba uuid.UUID) (*domain.Bank_account, error) {
	rows := repostitory.db_con.QueryRow(ctx, "select * from bank.bank_account where id_ba = $1;", id_ba)
	ba := domain.Bank_account{}
	err := rows.Scan(&ba)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &ba, err
}
