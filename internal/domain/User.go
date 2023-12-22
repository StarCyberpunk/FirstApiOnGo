package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password []byte    `json:"password"`
	Role_Id  int       `json:"role"`
	Email    string    `json:"email"`
}
type UserAuthModel struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserRegisterModel struct {
	Id         uuid.UUID `json:"id"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	Role_Id    int       `json:"role"`
	Email      string    `json:"email"`
	PassSerial int       `json:"pass_serial"`
	PassNumber int       `json:"pass_number"`
	CashTotal  float64   `json:"cash_total"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (uuid.UUID, error)
	FindUser(ctx context.Context, us UserAuthModel) (*User, error)
}
