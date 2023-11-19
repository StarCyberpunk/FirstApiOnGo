package postgres

import (
	"awesomeProject1/internal/domain"
	"database/sql"
	"github.com/gofrs/uuid"
	"log"
)

type OperationBARepository struct {
	db_con *sql.DB
}

func NewOperationBARepository(db_co *sql.DB) *OperationBARepository {
	return &OperationBARepository{
		db_con: db_co,
	}
}

func (repostitory *OperationBARepository) Create(oc domain.OperationBankAccount) (uuid.UUID, error) {
	_, err := repostitory.db_con.Query("INSERT INTO bank.operations_bank_account(id_op, date_op, id_ba_from, id_ba_to, total, id_cur, description ) VALUES ( $1, $2, $3, $4,$5,$6,$7);", oc.ID, oc.Date_op, oc.Id_accountFROM, oc.Id_accountTO, oc.Total, oc.Currency_id, oc.Description)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return oc.ID, err
}

func (repostitory *OperationBARepository) Find(oc domain.OperationBankAccount) (domain.OperationBankAccount, error) {
	rows, err := repostitory.db_con.Query("SELECT * FROM bank.operations_bank_account where id_op=$1;", oc.ID)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	opc := domain.OperationBankAccount{}
	for rows.Next() {
		rows.Scan(&oc.ID, &oc.Date_op, &oc.Id_accountFROM, &oc.Id_accountTO, &oc.Total, &oc.Currency_id, &oc.Description)
	}
	if err != nil {
		return domain.OperationBankAccount{}, err
	}
	return opc, err
}
