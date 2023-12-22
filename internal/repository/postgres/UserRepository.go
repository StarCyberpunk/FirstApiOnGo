package postgres

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/internal/pkg/persistence"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
)

type UserRepository struct {
	db_con persistence.Connection
}

func NewUserRepository(connection persistence.Connection) *UserRepository {
	return &UserRepository{
		db_con: connection,
	}
}

func (repostitory *UserRepository) CreateUser(ctx context.Context, us domain.User) (uuid.UUID, error) {
	_, err := repostitory.db_con.Exec(ctx, "INSERT INTO bank.users( id_user,login, password, id_role, email) VALUES ( $1, $2, $3, $4,$5);", us.ID, us.Login, us.Password, us.Role_Id, us.Email)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error: Unable to execute query: %w", err)
	}
	return us.ID, err
}

func (repostitory *UserRepository) FindUser(ctx context.Context, us domain.UserAuthModel) (*domain.User, error) {
	var rows = repostitory.db_con.QueryRow(ctx, "SELECT id_user, login, password, id_role, email FROM bank.users where login=$1;", us.Login)
	user := domain.User{}
	err := rows.Scan(&user.ID, &user.Login, &user.Password, &user.Role_Id, &user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, err
}
