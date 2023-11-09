package domain

import "github.com/gofrs/uuid"

type Card struct {
	ID                  uuid.UUID `json:"id_card"`
	Type_card           int       `json:"type_card_id"`
	Cash                float64   `json:"cash"`
	Number_card         int64     `json:"number_card"`
	Valid_date          string    `json:"valid_date"`
	CVV                 int16     `json:"cvv"`
	Block               bool      `json:"block"`
	Currency_of_card_Id int       `json:"id_currency"`
	Id_ba               uuid.UUID `json:"id_ba"`
}

type CardViewModel struct {
	Cards []Card    `json:"cards"`
	Id_ba uuid.UUID `json:"id_ba"`
}
