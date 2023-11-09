package OLD

/*
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

 User
func InitServer(hand *http.ServeMux) {

	handler := hand
	handler.HandleFunc("/auth", func(writer http.ResponseWriter, request *http.Request) {
		db, err := sql.Open("postgres", "user=postgres password=12345 port=5432 host=localhost dbname=postgres sslmode=disable")
		if err != nil {
			log.Fatalf("Error: Unable to connect to database: %v", err)
		}
		log.Println(request.Method)
		body, err := io.ReadAll(request.Body)
		switch request.Method {
		case "GET":
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("Page"))
			break
		case "POST":
			if err != nil {
				http.Error(writer, "Bad request", http.StatusBadRequest)
				return
			}
			if len(body) == 0 {
				http.Error(writer, "body is nil", http.StatusUnprocessableEntity)
				return
			}
			var us UserAuthModel
			_ = json.Unmarshal(body, &us)

			rows, err := db.Query("select id_user from users where login = $1 and passord=$2 ;", us.Login, us.Password)
			for rows.Next() {
				var id_user int64
				rows.Scan(&id_user)
				fmt.Printf(" Name: %d\n", id_user)
			}
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			defer db.Close()
			defer rows.Close()
			writer.WriteHeader(http.StatusOK)

		}
	})
	handler.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		db, err := sql.Open("postgres", "user=postgres password=12345 port=5432 host=localhost dbname=postgres sslmode=disable")
		if err != nil {
			log.Fatalf("Error: Unable to connect to database: %v", err)
		}
		log.Println(request.Method)
		body, err := io.ReadAll(request.Body)
		switch request.Method {
		case "GET":
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("Page"))
			break
		case "POST":
			if err != nil {
				http.Error(writer, "Bad request", http.StatusBadRequest)
				return
			}
			if len(body) == 0 {
				http.Error(writer, "body is nil", http.StatusUnprocessableEntity)
				return
			}
			var us UserRegisterModel
			_ = json.Unmarshal(body, &us)
			_, err := db.Query("INSERT INTO bank_account( pass_serial, pass_number, cash_total) VALUES ( $1, $2, $3);", us.PassSerial, us.PassNumber, 0)
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			rows, err := db.Query("select id_ba from bank_account where pass_serial = $1 and pass_number= $2;", us.PassSerial, us.PassNumber)
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			var id_ba int64
			for rows.Next() {
				rows.Scan(&id_ba)
			}
			_, err = db.Query("INSERT INTO public.users( login, password, id_role, email,id_ba) VALUES ( $1, $2, $3, $4,$5);", us.Login, us.Password, us.Role_Id, us.Email, id_ba)
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}

			defer db.Close()
			defer rows.Close()
			writer.WriteHeader(http.StatusOK)
		}
	})
	handler.HandleFunc("/profile", func(writer http.ResponseWriter, request *http.Request) {
		db, err := sql.Open("postgres", "user=postgres password=12345 port=5432 host=localhost dbname=postgres sslmode=disable")
		if err != nil {
			log.Fatalf("Error: Unable to connect to database: %v", err)
		}
		log.Println(request.Method)
		body, err := io.ReadAll(request.Body)
		switch request.Method {
		case "GET":
			if err != nil {
				http.Error(writer, "Bad request", http.StatusBadRequest)
				return
			}
			if len(body) == 0 {
				http.Error(writer, "body is nil", http.StatusUnprocessableEntity)
				return
			}
			var us UserRegisterModel
			_ = json.Unmarshal(body, &us)
			rows, err := db.Query("select bank_account.pass_serial,bank_account.pass_number from bank_account where bank_account.id_ba = $1 ", us.Bank_account_ID)
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			for rows.Next() {
				rows.Scan(&us.PassNumber, &us.PassSerial)
			}
			bb, _ := json.Marshal(us)
			writer.Write(bb)
			writer.WriteHeader(http.StatusOK)
			break
			case "POST":
			if err != nil {
				http.Error(writer, "Bad request", http.StatusBadRequest)
				return
			}
			if len(body) == 0 {
				http.Error(writer, "body is nil", http.StatusUnprocessableEntity)
				return
			}
			var us UserRegisterModel
			_ = json.Unmarshal(body, &us)
			_, err := db.Query("INSERT INTO bank_account( pass_serial, pass_number, cash_total) VALUES ( $1, $2, $3);", us.PassSerial, us.PassNumber, 0)
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			rows, err := db.Query("select id_ba from bank_account where pass_serial = $1 and pass_number= $2;", us.PassSerial, us.PassNumber)
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			var id_ba int64
			for rows.Next() {
				rows.Scan(&id_ba)
			}
			_, err = db.Query("INSERT INTO public.users( login, password, id_role, email,id_ba) VALUES ( $1, $2, $3, $4,$5);", us.Login, us.Password, us.Role_Id, us.Email, id_ba)
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}

			defer db.Close()
			defer rows.Close()
			writer.WriteHeader(http.StatusOK)
		}
	})
	handler.HandleFunc("/cards", func(writer http.ResponseWriter, request *http.Request) {
		db, err := sql.Open("postgres", "user=postgres password=12345 port=5432 host=localhost dbname=postgres sslmode=disable")
		if err != nil {
			log.Fatalf("Error: Unable to connect to database: %v", err)
		}
		log.Println(request.Method)
		body, err := io.ReadAll(request.Body)
		switch request.Method {
		case "GET":
			if err != nil {
				http.Error(writer, "Bad request", http.StatusBadRequest)
				return
			}
			if len(body) == 0 {
				http.Error(writer, "body is nil", http.StatusUnprocessableEntity)
				return
			}
			var us UserRegisterModel
			_ = json.Unmarshal(body, &us)
			rows, err := db.Query("select number_card,block from card  where id_ba = $1 ;", us.Bank_account_ID)
			var cards = CardViewModel{Id_ba: us.Bank_account_ID}
			for rows.Next() {
				var cardd Card
				rows.Scan(&cardd.Number_card, &cardd.Block)
				_ = append(cards.Cards, cardd)
			}
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			bb, _ := json.Marshal(cards)
			writer.Write(bb)
			defer db.Close()
			defer rows.Close()
			writer.WriteHeader(http.StatusOK)
			break
		case "POST":
			if err != nil {
				http.Error(writer, "Bad request", http.StatusBadRequest)
				return
			}
			if len(body) == 0 {
				http.Error(writer, "body is nil", http.StatusUnprocessableEntity)
				return
			}
			var us Card
			_ = json.Unmarshal(body, &us)
			_, err := db.Query("INSERT INTO card( id_currency, type_card_id, cash,number_card,cvv,block,id_ba,valid_date) VALUES ( $1, $2, $3,$4,$5,$6,$7,$8);", 1, 1, 0, us.Number_card, us.CVV, false, us.Id_ba, us.Valid_date)
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			defer db.Close()
			writer.WriteHeader(http.StatusOK)
			break
		}
	})
	handler.HandleFunc("/operations", func(writer http.ResponseWriter, request *http.Request) {
		db, err := sql.Open("postgres", "user=postgres password=12345 port=5432 host=localhost dbname=postgres sslmode=disable")
		if err != nil {
			log.Fatalf("Error: Unable to connect to database: %v", err)
		}
		log.Println(request.Method)
		body, err := io.ReadAll(request.Body)
		switch request.Method {
		case "GET":
			if err != nil {
				http.Error(writer, "Bad request", http.StatusBadRequest)
				return
			}
			if len(body) == 0 {
				http.Error(writer, "body is nil", http.StatusUnprocessableEntity)
				return
			}
			var us UserRegisterModel
			_ = json.Unmarshal(body, &us)
			rows, err := db.Query("select date_op,total,description from operations  where id_ba_to = $1 or id_ba_from=$1 ;", us.Bank_account_ID)
			var cards = OperationsViewModel{id_ba: us.Bank_account_ID}
			for rows.Next() {
				var cardd Operation
				rows.Scan(&cardd.Date_op, &cardd.Total, cardd.Description)
				_ = append(cards.Operations, cardd)
			}
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			bb, _ := json.Marshal(cards)
			writer.Write(bb)
			defer db.Close()
			defer rows.Close()
			writer.WriteHeader(http.StatusOK)
			break
		case "POST":
			if err != nil {
				http.Error(writer, "Bad request", http.StatusBadRequest)
				return
			}
			if len(body) == 0 {
				http.Error(writer, "body is nil", http.StatusUnprocessableEntity)
				return
			}
			var us Card
			_ = json.Unmarshal(body, &us)
			_, err := db.Query("INSERT INTO operations( ) VALUES ( $1, $2, $3,$4,$5,$6);", 1, 1, 0, us.Number_card, us.CVV, us.Id_ba)
			if err != nil {
				log.Fatalf("Error: Unable to execute query: %v", err)
			}
			defer db.Close()
			writer.WriteHeader(http.StatusOK)
			break
		}
	})


}*/
