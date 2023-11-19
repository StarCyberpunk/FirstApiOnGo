package domain

import "github.com/gofrs/uuid"

type OperationCard struct {
	ID          uuid.UUID
	Date_op     int64
	Id_cardFROM uuid.UUID
	Id_cardTO   uuid.UUID
	Total       float64
	Currency_id int
	Description string
}

type OperationsViewModel struct {
	OperationsCard        []OperationCard        `json:"operations_card"`
	OperationsBankAccount []OperationBankAccount `json:"operations_bank_account"`
	id_ba                 uuid.UUID              `json:"id_ba"`
	id_card               uuid.UUID              `json:"id_card"`
}
