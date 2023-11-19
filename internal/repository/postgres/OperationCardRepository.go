package postgres

import (
	"awesomeProject1/internal/domain"
	"database/sql"
	"github.com/gofrs/uuid"
	"log"
)

type OperationCardRepository struct {
	db_con *sql.DB
}

func NewOperationCardRepository(db_co *sql.DB) *OperationCardRepository {
	return &OperationCardRepository{
		db_con: db_co,
	}
}

func (repostitory *OperationCardRepository) Create(oc domain.OperationCard) (uuid.UUID, error) {
	_, err := repostitory.db_con.Query("INSERT INTO bank.operations_card(id_op, date_op, id_card_from, id_card_to, total, id_cur, description ) VALUES ( $1, $2, $3, $4,$5,$6,$7);", oc.ID, oc.Date_op, oc.Id_cardFROM, oc.Id_cardTO, oc.Total, oc.Currency_id, oc.Description)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return oc.ID, err
}

func (repostitory *OperationCardRepository) Find(oc domain.OperationCard) (domain.OperationCard, error) {
	rows, err := repostitory.db_con.Query("SELECT * FROM bank.operations_card where id_op=$1;", oc.ID)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	opc := domain.OperationCard{}
	for rows.Next() {
		rows.Scan(&oc.ID, &oc.Date_op, &oc.Id_cardFROM, &oc.Id_cardTO, &oc.Total, &oc.Currency_id, &oc.Description)
	}
	if err != nil {
		return domain.OperationCard{}, err
	}
	return opc, err
}
