package main

import (
	"awesomeProject1/internal/di"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	/*db, err := sql.Open("postgres", "user=postgres password=12345 port=5432 host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Error: Unable to connect to database: %v", err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT id_user,login FROM users")
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("User ID: %d, Name: %s\n", id, name)
	}
	handler := http.NewServeMux()
	OLD.InitServer(handler)
	err := http.ListenAndServe("localhost:8000", handler)
	if err != nil {
		log.Fatal(err)
	}
	*/
	cont := di.NewContainer()
	cont.InitRepository()
	cont.InitUseCases()
	err := http.ListenAndServe("localhost:8000", cont.HTTPRouter())
	if err != nil {
		log.Fatal(err)
	}

}
