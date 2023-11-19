package domain

import "github.com/gofrs/uuid"

type OperationBankAccount struct {
	ID             uuid.UUID
	Date_op        int64
	Id_accountTO   uuid.UUID
	Id_accountFROM uuid.UUID
	Total          float64
	Currency_id    int
	Description    string
}
