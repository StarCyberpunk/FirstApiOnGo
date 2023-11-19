package postgres

import (
	"awesomeProject1/internal/domain"
	"database/sql"
	"github.com/gofrs/uuid"
	"log"
)

type CardRepository struct {
	db_con *sql.DB
}

func NewCardRepository(db_co *sql.DB) *CardRepository {
	return &CardRepository{
		db_con: db_co,
	}
}

func (repostitory *CardRepository) Create(ca domain.Card) (uuid.UUID, error) {
	_, err := repostitory.db_con.Query("INSERT INTO bank.cards( id_card, id_currency, type_card_id, cash, number_card, valid_date, cvv, block, id_ba)  VALUES ( $1, $2, $3, $4,$5,$6,$7,$8,$9);", ca.ID, ca.Currency_of_card_Id, ca.Type_card, ca.Cash, ca.Number_card, ca.Valid_date, ca.CVV, ca.Block, ca.Id_ba)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return ca.ID, err
}

func (repostitory *CardRepository) Find(id uuid.UUID) (domain.Card, error) {
	rows, err := repostitory.db_con.Query("SELECT * FROM bank.cards where id_card=$1;", id)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	ca := domain.Card{}
	for rows.Next() {
		rows.Scan(&ca.ID, &ca.Currency_of_card_Id, &ca.Type_card, &ca.Cash, &ca.Number_card, &ca.Valid_date, &ca.CVV, &ca.Block, &ca.Id_ba)
	}
	if err != nil {
		return domain.Card{}, err
	}
	return ca, err
}
