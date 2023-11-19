package postgres

import (
	"awesomeProject1/internal/domain"
	"database/sql"
	"log"
)

type CurrencyRepository struct {
	db_con *sql.DB
}

func NewCurrencyRepository(db_co *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{
		db_con: db_co,
	}
}

func (repostitory *CurrencyRepository) Create(cur domain.Currency) (int, error) {
	_, err := repostitory.db_con.Query("INSERT INTO bank.currency(id,name,one_to_rub) VALUES ( $1, $2, $3);", cur.ID, cur.Name, cur.OneToRub)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return cur.ID, err
}

func (repostitory *CurrencyRepository) Find(id int) (domain.Currency, error) {
	rows, err := repostitory.db_con.Query("SELECT * FROM bank.users where login=$1;", id)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	cur := domain.Currency{}
	for rows.Next() {
		rows.Scan(&cur.ID, &cur.Name, &cur.OneToRub)
	}
	if err != nil {
		return domain.Currency{}, err
	}
	return cur, err
}
