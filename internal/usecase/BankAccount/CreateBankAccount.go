package BankAccount

import (
	"awesomeProject1/internal/domain"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
)

type CreateBankAccountUseCase struct {
	UserRepository        domain.UserRepository
	BankAccountRepository domain.BankAccountRepository
}

type filterCommand struct {
	Name   string
	UserID uuid.UUID
	Offset int
	Limit  int
}

func NewCreateBankAccountUseCase(userRepository domain.UserRepository, bankRepository domain.BankAccountRepository) *CreateBankAccountUseCase {
	return &CreateBankAccountUseCase{
		UserRepository:        userRepository,
		BankAccountRepository: bankRepository,
	}
}
func (useCase *CreateBankAccountUseCase) Create(ctx context.Context, bankAccount domain.Bank_account) (*domain.Bank_account, error) {
	bafromdb, err := useCase.BankAccountRepository.FindBankAccountByNameId(ctx, bankAccount.IdUser, bankAccount.Name)
	if err != nil {
		return nil, fmt.Errorf("Error: Bank_Account dont created: %w", err)
	}
	if bafromdb != nil {
		return nil, fmt.Errorf("Error: Bank_Account is exist:")
	}
	id_ba, _ := uuid.NewV4()
	ba := domain.Bank_account{ID: id_ba, Name: bankAccount.Name, PassSerial: bankAccount.PassSerial, PassNumber: bankAccount.PassNumber, CashTotal: bankAccount.CashTotal, IdUser: bankAccount.IdUser}
	_, err = useCase.BankAccountRepository.CreateBankAccount(ctx, ba)
	if err != nil {
		return nil, fmt.Errorf("Error: Bank_Account dont created: %w", err)
	}
	return &ba, nil
}
func (useCase *CreateBankAccountUseCase) Read(ctx context.Context, bankAccount domain.Bank_account) (*domain.Bank_account, error) {
	bafromdb, err := useCase.BankAccountRepository.FindBankAccountByNameId(ctx, bankAccount.IdUser, bankAccount.Name)
	if err != nil {
		return nil, fmt.Errorf("Error: Bank_Account dont created: %w", err)
	}
	if bafromdb != nil {
		return nil, fmt.Errorf("Error: Bank_Account is exist:")
	}
	id_ba, _ := uuid.NewV4()
	ba := domain.Bank_account{ID: id_ba, Name: bankAccount.Name, PassSerial: bankAccount.PassSerial, PassNumber: bankAccount.PassNumber, CashTotal: bankAccount.CashTotal, IdUser: bankAccount.IdUser}
	_, err = useCase.BankAccountRepository.CreateBankAccount(ctx, ba)
	if err != nil {
		return nil, fmt.Errorf("Error: Bank_Account dont created: %w", err)
	}
	return &ba, nil
}
func (useCase *CreateBankAccountUseCase) ReadAccounts(ctx context.Context, command *filterCommand) ([]domain.Bank_account, error) {
	bafromdb, err := useCase.BankAccountRepository.FindUserAccountsILike(ctx, command.Name, command.Offset, command.Limit, command.UserID)
	if err != nil {
		return nil, fmt.Errorf("Error: Bank_Account dont created: %w", err)
	}
	if bafromdb != nil {
		return nil, fmt.Errorf("Error: Bank_Account is exist:")
	}
	return bafromdb, err
}
