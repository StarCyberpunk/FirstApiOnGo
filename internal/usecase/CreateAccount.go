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
	UserRepository domain.UserRepository
}

func NewCreateUserUseCase(userRepository domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (useCase *CreateUserUseCase) Handle(user domain.UserRegisterModel) (uuid.UUID, error) {
	id, _ := uuid.NewV4()
	//TODO: Проверить есть ли пользователь в бд
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user_pa := domain.User{Login: user.Login, Password: hash, ID: id, Role_Id: user.Role_Id, Email: user.Email, Bank_account_ID: user.Bank_account_ID}
	id_us, err := useCase.UserRepository.CreateUser(user_pa)
	if err != nil {
		return uuid.Nil, err
	}
	return id_us, nil
}
func (useCase *CreateUserUseCase) Login(user domain.UserAuthModel) (string, error) {
	//hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	//user_pa := domain.User{Login: user.Login, Password: hash}
	us, err := useCase.UserRepository.FindUser(user)
	if err != nil {
		return "", err
	}
	payload := jwt.MapClaims{
		"sub": us.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}
