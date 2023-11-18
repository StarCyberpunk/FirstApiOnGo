package usecase

import (
	"awesomeProject1/internal/domain"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type CreateUserUseCase struct {
	UserRepository        domain.UserRepository
	BankAccountRepository domain.BankAccountRepository
}

func NewCreateUserUseCase(userRepository domain.UserRepository, bankRepository domain.BankAccountRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository:        userRepository,
		BankAccountRepository: bankRepository,
	}
}

func (useCase *CreateUserUseCase) Handle(user domain.UserRegisterModel) (uuid.UUID, error) {
	id, _ := uuid.NewV4()
	id_ba, _ := uuid.NewV4()
	//TODO: Проверить есть ли пользователь в бд
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	ba := domain.Bank_account{ID: id_ba, PassSerial: user.PassSerial, PassNumber: user.PassNumber, CashTotal: user.CashTotal}
	id_bba, err := useCase.BankAccountRepository.CreateBankAccount(ba)
	user_pa := domain.User{Login: user.Login, Password: hash, ID: id, Role_Id: user.Role_Id, Email: user.Email, Bank_account_ID: id_bba}
	id_us, err := useCase.UserRepository.CreateUser(user_pa)
	if err != nil {
		return uuid.Nil, err
	}
	return id_us, nil
}
func (useCase *CreateUserUseCase) Login(user domain.UserAuthModel) (string, error) {
	us, err := useCase.UserRepository.FindUser(user)
	if err != nil {
		return "", err
	}
	payload := jwt.MapClaims{
		"sub": us.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}
