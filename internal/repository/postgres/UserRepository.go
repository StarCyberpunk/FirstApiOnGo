package postgres

import (
	"awesomeProject1/internal/domain"
	"database/sql"
	"github.com/gofrs/uuid"
	"log"
)

type UserRepository struct {
	db_con *sql.DB
}

func NewUserRepository(db_co *sql.DB) *UserRepository {
	return &UserRepository{
		db_con: db_co,
	}
}

func (repostitory *UserRepository) CreateUser(us domain.User) {
	_, err := repostitory.db_con.Query("INSERT INTO public.users( login, password, id_role, email,id_ba) VALUES ( $1, $2, $3, $4,$5);", us.Login, us.Password, us.Role_Id, us.Email, us.Bank_account_ID)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
}
func (repostitory *UserRepository) ReadUser(usid uuid.UUID) domain.User {
	rows, err := repostitory.db_con.Query("select * from public.users where id = $1;", usid)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	var user domain.User
	for rows.Next() {
		rows.Scan(&user)
	}
	return user
}
