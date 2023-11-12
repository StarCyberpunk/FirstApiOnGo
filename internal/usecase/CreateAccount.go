package usecase

import (
	"awesomeProject1/internal/domain"
	"github.com/gofrs/uuid"
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
	user_pa := domain.User{Login: user.Login, Password: user.Password, ID: id, Role_Id: user.Role_Id, Email: user.Email, Bank_account_ID: user.Bank_account_ID}
	id_us, err := useCase.UserRepository.CreateUser(user_pa)
	if err != nil {
		return uuid.Nil, err
	}
	return id_us, nil
}
