package usecase

import (
	"awesomeProject1/internal/domain"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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
	userauth := domain.UserAuthModel{Login: user.Login}
	us, err := useCase.UserRepository.FindUser(userauth)
	if err != nil {
		return uuid.Nil, err
	}
	//возващать нужно новые ошибки
	if us != nil {
		return uuid.Nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error: Password is not valid: %w", err)
	}
	user_pa := domain.User{Login: user.Login, Password: hash, ID: id, Role_Id: user.Role_Id, Email: user.Email}
	ba := domain.Bank_account{ID: id_ba, PassSerial: user.PassSerial, PassNumber: user.PassNumber, CashTotal: user.CashTotal, IdUser: user.Id}
	_, err = useCase.UserRepository.CreateUser(user_pa)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error: User dont created: %w", err)
	}
	_, err = useCase.BankAccountRepository.CreateBankAccount(ba)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error: Bank_Account dont created: %w", err)
	}
	return id, nil
}
func (useCase *CreateUserUseCase) Login(user domain.UserAuthModel) (string, error) {
	us, err := useCase.UserRepository.FindUser(user)
	if err != nil {
		return "", fmt.Errorf("Not found: %w", err)
	}
	err = bcrypt.CompareHashAndPassword(us.Password, []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("Password is wrong: %w", err)
	}
	// Вынести генерацию в отдельный метод
	payload := jwt.MapClaims{
		"sub": us.Email,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	rr := context.Background().Value("secretKey")
	fmt.Println(rr)
	// вынести в Di
	t, err := token.SignedString([]byte("2"))
	if err != nil {
		return "", err
	}
	return t, nil
}
