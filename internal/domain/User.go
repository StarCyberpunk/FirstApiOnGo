package domain

import "github.com/gofrs/uuid"

type User struct {
	ID              uuid.UUID `json:"id"`
	Login           string    `json:"login"`
	Password        []byte    `json:"password"`
	Role_Id         int       `json:"role"`
	Email           string    `json:"email"`
	Bank_account_ID uuid.UUID `json:"ba_ID"`
}
type UserAuthModel struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserRegisterModel struct {
	Id              uuid.UUID `json:"id"`
	Login           string    `json:"login"`
	Password        string    `json:"password"`
	Role_Id         int       `json:"role"`
	Email           string    `json:"email"`
	Bank_account_ID uuid.UUID `json:"id_ba"`
	PassSerial      int       `json:"pass_serial"`
	PassNumber      int       `json:"pass_number"`
	CashTotal       float64   `json:"cash_total"`
}

type UserRepository interface {
	CreateUser(user User) (uuid.UUID, error)
	FindUser(us UserAuthModel) (User, error)
}
