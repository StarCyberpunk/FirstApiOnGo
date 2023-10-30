package and

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type Person struct {
	FName, SName    string
	Login           string
	Password        string
	Role            string
	Bank_account_ID int64
}
type Answer struct {
	person Person
	code   bool
}

var NameDB = "hello.txt"
var Roles = [3]string{"Admin", "User", "Guest"}
var Auth = false
var User Person

type Bank_account struct {
	ID           int64
	PassSerial   int16
	PassNumber   int16
	CashTotal    float64
	Cards        []Card
	Operations   []Operation
	Ogranichenie []Ogranichenie
}
type Operation struct {
	ID             int64
	Date           int64 //UNIX
	Id_accountTO   int64
	Id_accountFROM int64
	Id_cardFROM    int64
	Id_cardTO      int64
	Total          float64
	Currency_op    Currency
	Comission      float64
	Description    string
}
type Card struct {
	ID               int64
	Type_card        int
	Id_client        int64
	Cash             float64
	Number           int64
	Valid_date       int64
	CVV              int32
	Block            bool
	Currency_of_card int
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

type Ogranichenie struct {
	ID   int
	Name string
}
type BASS struct {
	Persons []Person
	BASS    []Bank_account
}

var BAS BASS

func createPerson() Person {
	var person = Person{}
	fmt.Println("Регистрация")
	fmt.Println("Введите имя")
	fmt.Scan(&person.FName)
	fmt.Println("Введите фамилия")
	fmt.Scan(&person.SName)
	for {
		fmt.Println("Введите логин")
		fmt.Scan(&person.Login)
		if !haveLogin(person.Login) {
			break
		} else {
			fmt.Println("Логин занят")
		}
	}
	fmt.Println("Введите пароль")
	fmt.Scan(&person.Password)
	person.Role = Roles[1]
	var seria int16
	var numb int16
	fmt.Println("Введите серию паспорта")
	fmt.Scan(&seria)
	fmt.Println("Введите номер паспорта")
	fmt.Scan(&numb)
	person.Bank_account_ID = create_Bank_Account(seria, numb, person).ID
	save_pers(person)
	return person
}
func haveLogin(login string) bool {
	for _, el := range BAS.Persons {
		if el.Login == login {
			return true
		}
	}
	return false
}
func findPerson(login, password string) Answer {
	for _, el := range BAS.Persons {
		if el.Login == login && el.Password == password {
			return Answer{person: el, code: true}
		}
	}
	return Answer{person: Person{}, code: false}
}
func auth() {
	if Auth {
		fmt.Println("Вы авторизованы")
		return
	}
	var login, password string
	fmt.Println("Введите логин")
	fmt.Scan(&login)
	fmt.Println("Введите пароль")
	fmt.Scan(&password)
	var res = findPerson(login, password)
	if res.code {
		fmt.Println("Успешно авторизован " + res.person.SName + " " + res.person.FName + "!")
		Auth = true
		User = res.person
	} else {
		fmt.Println("Авторизируйтесь заново")
	}
}
func save_pers(person Person) {
	jsonFile, err := os.Open("BAS.json")
	byteValue, err := io.ReadAll(jsonFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(byteValue, &BAS)
	BAS.Persons = append(BAS.Persons, person)
	defer jsonFile.Close()

	file, err := json.Marshal(BAS)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("BAS.json", file, 0644)
	if err != nil {
		panic(err)
	}
}
func LoadDB() {
	jsonFile, err := os.Open("BAS.json")
	byteValue, err := io.ReadAll(jsonFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(byteValue, &BAS)
	defer jsonFile.Close()

}
func UpdateDB() {
	jsonFile, err := os.Open("BAS.json")

	// if we os.Open returns an error then handle it
	defer jsonFile.Close()
	file, err := json.Marshal(BAS)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("BAS.json", file, 0644)
	if err != nil {
		panic(err)
	}
}
func saveBASinJSON(nw Bank_account) {
	jsonFile, err := os.Open("BAS.json")
	byteValue, err := io.ReadAll(jsonFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(byteValue, &BAS)
	BAS.BASS = append(BAS.BASS, nw)
	defer jsonFile.Close()

	file, err := json.Marshal(BAS)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("BAS.json", file, 0644)
	if err != nil {
		panic(err)
	}
}
func create_Bank_Account(passSerial int16, passNumber int16, person Person) Bank_account {
	if person.Bank_account_ID != 0 {
		return Bank_account{}
	}
	if find_Bank_Account(person.Bank_account_ID).ID != 0 {
		return Bank_account{}
	}
	nw := Bank_account{ID: rand.Int63n(100000),
		PassSerial: passSerial,
		PassNumber: passNumber,
		Cards:      []Card{}, Operations: []Operation{}, CashTotal: 0}
	saveBASinJSON(nw)
	return nw
}
func close_Bank_Account(id int64) {
	for _, el := range BAS.Persons {
		if el.Bank_account_ID == id {
			el.Bank_account_ID = 0
			UpdateDB()
			return
		}
	}
	for i, el := range BAS.BASS {
		if el.ID == id {
			BAS.BASS = append(BAS.BASS[:i], BAS.BASS[i+1:]...)
			UpdateDB()
			return
		}
	}

}
func find_Bank_Account(id int64) Bank_account {
	for i := 0; i < len(BAS.BASS); i++ {
		if BAS.BASS[i].ID == id {
			return BAS.BASS[i]
		}
	}
	return Bank_account{}
}
func show_Bank_Account(id int64) {
	ba := find_Bank_Account(id)
	fmt.Println("Карты:")
	for _, el := range ba.Cards {
		fmt.Print("Название " + string(el.Number))
		fmt.Printf("%d \n", el.Cash)
	}
	fmt.Println("Операции:")
	for _, el := range ba.Operations {
		fmt.Println("Название " + string(el.Date))
	}
}
func set_limit_Bank_Account(id int64, ogranichenie Ogranichenie) {
	for _, el := range BAS.BASS {
		if el.ID == id {
			_ = append(el.Ogranichenie, ogranichenie)
			UpdateDB()
			return
		}
	}
}
func set_card(id int64, card Card) {
	for _, el := range BAS.BASS {
		if el.ID == id {
			_ = append(el.Cards, card)
			UpdateDB()
			return
		}
	}
}
func create_card(id int64, typ int, cur int) {
	card := Card{ID: rand.Int63n(100000), Type_card: typ, Cash: 0,
		Number: rand.Int63n(1000000000),
		CVV:    rand.Int31n(999),
		Block:  true, Currency_of_card: cur}
	for _, el := range BAS.BASS {
		if el.ID == id {
			_ = append(el.Cards, card)
			UpdateDB()
			return
		}
	}
}
func active_block_card(id int64, card Card, act_blo bool) {
	card.Block = act_blo
	for _, el := range BAS.BASS {
		if el.ID == id {
			_ = append(el.Cards, card)
			UpdateDB()
			return
		}
	}
}
func show_card(card Card) {
	fmt.Printf("Номер: %d,Баланс: %d Статус: %b ", card.Number, card.Cash, card.Block)
}
func convert_cur(cash float64, currency2 Currency) float64 {
	return cash * currency2.OneToRub
}
func send_moneyCard(id1 int64, id2 int64, card Card, card2 Card, cash float64) {
	account := find_Bank_Account(id1)
	account2 := find_Bank_Account(id2)
	op := Operation{ID: rand.Int63n(100000000),
		Id_cardFROM:    card.ID,
		Id_cardTO:      card2.ID,
		Id_accountTO:   account2.ID,
		Id_accountFROM: account.ID,
		Currency_op:    currency[card2.Currency_of_card], Total: cash}
	account.CashTotal -= cash
	account2.CashTotal += cash
	card.Cash -= cash
	card2.Cash += cash
	_ = append(account.Operations, op)
	_ = append(account2.Operations, op)
	UpdateDB()
}
func send_moneyAc(id1 int64, id2 int64, cash float64, currency2 Currency) {
	account := find_Bank_Account(id1)
	account2 := find_Bank_Account(id2)
	op := Operation{ID: rand.Int63n(100000000),
		Id_accountTO:   account2.ID,
		Id_accountFROM: account.ID,
		Currency_op:    currency2, Total: cash}
	account.CashTotal -= cash
	account2.CashTotal += cash
	_ = append(account.Operations, op)
	_ = append(account2.Operations, op)
	UpdateDB()
}
func send_money_self(id int64, card Card, card2 Card, cash float64) {
	account := find_Bank_Account(id)
	op := Operation{ID: rand.Int63n(100000000),
		Id_cardFROM:    card.ID,
		Id_cardTO:      card2.ID,
		Id_accountTO:   account.ID,
		Id_accountFROM: account.ID,
		Currency_op:    currency[card2.Currency_of_card], Total: cash}
	card.Cash -= cash
	card2.Cash += cash
	_ = append(account.Operations, op)
	UpdateDB()
}
func convert_card_cur(card Card, currency2 Currency) {
	card.Currency_of_card = currency2.ID
	card.Cash = card.Cash * currency2.OneToRub
	UpdateDB()
}

/*
	func menu() bool {
		if !Auth {
			fmt.Println("---------------------")
			fmt.Println(strconv.Itoa(1) + " Авторизация")
			fmt.Println(strconv.Itoa(2) + " Регистрация")
			fmt.Println(strconv.Itoa(3) + " Выход")
			fmt.Println("---------------------")
			var ii int
			fmt.Fscan(os.Stdin, &ii)
			switch ii {
			case 1:
				auth()
				return true
			case 2:
				person := createPerson()
				save_pers(person)
				return true
			case 3:
				return false
			default:
				return true
			}
		} else {
			fmt.Println("---------------------")
			fmt.Println(strconv.Itoa(1) + "Банковский аккаунт")
			fmt.Println(strconv.Itoa(2) + " Карты")
			fmt.Println(strconv.Itoa(3) + " Валюты")
			fmt.Println(strconv.Itoa(4) + " Операции")
			fmt.Println(strconv.Itoa(5) + " Выход")
			fmt.Println("---------------------")
			var ii int
			fmt.Fscan(os.Stdin, &ii)
			switch ii {
			case 1:
				fmt.Println("---------------------")
				fmt.Println(strconv.Itoa(1) + "Создание/закрытие")
				fmt.Println(strconv.Itoa(2) + "Посмотреть баланс и история")
				fmt.Println(strconv.Itoa(3) + "Установка лимита")
				fmt.Println(strconv.Itoa(4) + " Операции")
				fmt.Println(strconv.Itoa(5) + " Назад")
				fmt.Println("---------------------")
				var ii int
				fmt.Fscan(os.Stdin, &ii)
				switch ii {
				case 1:
					fmt.Println("---------------------")
					fmt.Println(strconv.Itoa(1) + "Создание")
					fmt.Println(strconv.Itoa(2) + "Закрытие")
					fmt.Println(strconv.Itoa(3) + " Назад")
					fmt.Println("---------------------")
					var ii int
					fmt.Fscan(os.Stdin, &ii)
					switch ii {
					case 1:
						fmt.Println("Уже создан")
						return true
					case 2:
						close_Bank_Account()
						fmt.Println("Закрыт")
						return true
					case 3:
						return true
					default:
						return true
					}
					return true
				case 2:
					person := createPerson()
					save_pers(person)
					return true
				case 3:
					return false
				default:
					return true
				}
				return true
			case 2:

				return true
			case 3:
				return true
			case 4:
				return true
			case 5:
				return true
			default:
				return true
			}
		}
		return true
	}
*/
