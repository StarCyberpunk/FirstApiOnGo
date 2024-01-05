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
	_, err := repostitory.db_con.Exec(ctx, "INSERT INTO bank.bank_account( id_ba,name,pass_serial, pass_number, cash_total,id_user) VALUES ( $1, $2, $3,$4,$5,$6);", ba.ID, ba.Name, ba.PassSerial, ba.PassNumber, 0, ba.IdUser)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return ba.ID, err
}

func (repostitory *BankAccountRepository) FindBankAccountById(ctx context.Context, id_ba uuid.UUID) (*domain.Bank_account, error) {
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

func (repostitory *BankAccountRepository) FindBankAccountByNameId(ctx context.Context, id_ba uuid.UUID, name string) (*domain.Bank_account, error) {
	rows := repostitory.db_con.QueryRow(ctx, "select * from bank.bank_account where id_ba = $1 and name=$2;", id_ba, name)
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

func (repostitory *BankAccountRepository) FindByIDForUpdate(ctx context.Context, id uuid.UUID) (*domain.Bank_account, error) {
	rows := repostitory.db_con.QueryRow(ctx, "SELECT account_id, name, user_id, balance FROM bank.account WHERE user_id = ($1) FOR UPDATE", id)
	ba := domain.Bank_account{}
	err := rows.Scan(&ba)
	if err != nil {
		return nil, err
	}
	return &ba, nil
}

func (repostitory *BankAccountRepository) UpdateAccountBalance(ctx context.Context, id uuid.UUID, deposit int) error {
	_, err := repostitory.db_con.Exec(ctx, "UPDATE bank.account SET balance = ($2) WHERE account_id=$1", id, deposit)
	if err != nil {
		return err
	}
	return err
}

func (repostitory *BankAccountRepository) FindUserAccountsILike(ctx context.Context, name string, offset int, limit int, userID uuid.UUID) ([]domain.Bank_account, error) {
	query := "SELECT account_id, name, balance, user_id FROM bank.account WHERE user_id = $1 AND name ILIKE $2 LIMIT $3 OFFSET $4;"
	rows, err := repostitory.db_con.Query(ctx, query, userID, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []domain.Bank_account
	ba := domain.Bank_account{}
	for rows.Next() {
		err = rows.Scan(ba)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, ba)
	}
	return accounts, err
}
