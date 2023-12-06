package postgres

import (
	"awesomeProject1/internal/domain"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
)

type UserRepository struct {
	db_con *sql.DB
}

func NewUserRepository(db_co *sql.DB) *UserRepository {
	return &UserRepository{
		db_con: db_co,
	}
}

func (repostitory *UserRepository) CreateUser(us domain.User) (uuid.UUID, error) {
	_, err := repostitory.db_con.Exec("INSERT INTO bank.users( id_user,login, password, id_role, email) VALUES ( $1, $2, $3, $4,$5,$6);", us.ID, us.Login, us.Password, us.Role_Id, us.Email)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error: Unable to execute query: %w", err)
	}
	return us.ID, err
}

func (repostitory *UserRepository) FindUser(us domain.UserAuthModel) (*domain.User, error) {
	var rows = repostitory.db_con.QueryRow("SELECT id_user, login, password, id_role, email FROM bank.users where login=$1;", us.Login)
	var user domain.User
	err := rows.Err()
	if err != nil {
		return nil, fmt.Errorf("Error: Unable to execute query: %w", err)
	}
	rows.Scan(&user.ID, &user.Login, &user.Password, &user.Role_Id, &user.Email)
	return &user, err
}
