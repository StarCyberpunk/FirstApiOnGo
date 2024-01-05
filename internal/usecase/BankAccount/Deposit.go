package BankAccount

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/internal/pkg/persistence"
	"context"
	"github.com/gofrs/uuid"
)

type DepositAccountUseCase struct {
	bankAccountRepository domain.BankAccountRepository
	transactionManager    persistence.TransactionManager
}

type DepositAccountCommand struct {
	UserID  uuid.UUID
	Deposit int
}

func NewDepositAccountUseCase(bankAccountRepository domain.BankAccountRepository, transactionManager persistence.TransactionManager) *DepositAccountUseCase {
	return &DepositAccountUseCase{
		bankAccountRepository: bankAccountRepository,
		transactionManager:    transactionManager,
	}
}

func (useCase DepositAccountUseCase) DepositAccountHandler(ctx context.Context, command *DepositAccountCommand) error {
	return useCase.transactionManager.Do(ctx, func(ctx context.Context) error {
		bankAccount, err := useCase.bankAccountRepository.FindByIDForUpdate(ctx, command.UserID)
		if err != nil {
			return err
		}
		bankAccount.Balance += command.Deposit
		err = useCase.bankAccountRepository.UpdateAccountBalance(ctx, bankAccount.IdUser, bankAccount.Balance)
		if err != nil {
			return err
		}

		err = useCase.bankAccountRepository.UpdateAccountBalance(ctx, command.UserID, command.Deposit)
		return err
	})

}
