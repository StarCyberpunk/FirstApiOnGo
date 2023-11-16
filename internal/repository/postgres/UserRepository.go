package postgres

import (
	"awesomeProject1/internal/domain"
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (repostitory *UserRepository) CreateUser(us domain.User) (uuid.UUID, error) {
	_, err := repostitory.db_con.Query("INSERT INTO public.users( id,login, password, id_role, email,id_ba) VALUES ( $1, $2, $3, $4,$5,$6);", us.ID, us.Login, us.Password, us.Role_Id, us.Email, us.Bank_account_ID)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return us.ID, err
}

var ErrUnauthorized = errors.New("Unauthorized")

func (repostitory *UserRepository) FindUser(us domain.UserAuthModel) (domain.User, error) {
	rows, err := repostitory.db_con.Query("SELECT id, login, id_role, email, id_ba, password FROM public.users where login=$1;", us.Login)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	user := domain.User{}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Login, &user.Role_Id, &user.Email, &user.Bank_account_ID, &user.Password)
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(us.Password))
	if err != nil {
		return domain.User{}, err
	}
	return user, err
}
