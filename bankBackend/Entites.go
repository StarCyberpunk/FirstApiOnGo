package bankBackend

type User struct {
	ID              int64  `json:"id"`
	Login           string `json:"login"`
	Password        string `json:"password"`
	Role_Id         int    `json:"role"`
	Email           string `json:"email"`
	Bank_account_ID int64  `json:"Bank_Account_ID"`
}
type UserAuthModel struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserRegisterModel struct {
	Login           string  `json:"login"`
	Password        string  `json:"password"`
	Role_Id         int     `json:"role"`
	Email           string  `json:"email"`
	Bank_account_ID int64   `json:"id_ba"`
	PassSerial      int     `json:"pass_serial"`
	PassNumber      int     `json:"pass_number"`
	CashTotal       float64 `json:"cash_total"`
}
type Bank_account struct {
	ID         int64   `json:"id_ba"`
	PassSerial int     `json:"pass_serial"`
	PassNumber int     `json:"pass_number"`
	CashTotal  float64 `json:"cash_total"`
}
type OperationsViewModel struct {
	Operations []Operation `json:"operations"`
	id_ba      int64       `json:"id_ba"`
	id_card    int64       `json:"id_card"`
}
type Operation struct {
	ID             int64
	Date_op        int64
	Id_accountTO   int64
	Id_accountFROM int64
	Id_cardFROM    int64
	Id_cardTO      int64
	Total          float64
	Currency_id    int64
	Description    string
}
type CardViewModel struct {
	Cards []Card `json:"cards"`
	Id_ba int64  `json:"id_ba"`
}
type Card struct {
	ID                  int64   `json:"id_card"`
	Type_card           int     `json:"type_card_id"`
	Cash                float64 `json:"cash"`
	Number_card         int64   `json:"number_card"`
	Valid_date          string  `json:"valid_date"`
	CVV                 int16   `json:"cvv"`
	Block               bool    `json:"block"`
	Currency_of_card_Id int     `json:"id_currency"`
	Id_ba               int64   `json:"id_ba"`
}
type Currency struct {
	ID       int
	Name     string
	OneToRub float64
}

var currency = map[int]Currency{
	1: {
		ID:       1,
		Name:     "USD",
		OneToRub: 100,
	},
	2: {ID: 2, Name: "EUR", OneToRub: 120},
	3: {ID: 3, Name: "RUB", OneToRub: 1},
}
