package usecase

import "awesomeProject1/internal/domain"

type CreateUserUseCase struct {
	UserRepository domain.UserRepository
}

func NewCreateUserUseCase(userRepository domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (useCase *CreateUserUseCase) Handle(user *domain.User) (*domain.User, error) {
	err := useCase.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
