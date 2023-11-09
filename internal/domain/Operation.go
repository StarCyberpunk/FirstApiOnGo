package domain

import "github.com/gofrs/uuid"

type Operation struct {
	ID             uuid.UUID
	Date_op        int64
	Id_accountTO   uuid.UUID
	Id_accountFROM uuid.UUID
	Id_cardFROM    uuid.UUID
	Id_cardTO      uuid.UUID
	Total          float64
	Currency_id    int
	Description    string
}

type OperationsViewModel struct {
	Operations []Operation `json:"operations"`
	id_ba      uuid.UUID   `json:"id_ba"`
	id_card    uuid.UUID   `json:"id_card"`
}
