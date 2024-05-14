package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {

	// run after main to catch any panics
	// panics go up the call stack until they find a deferred recover function
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("A fatal error has occurred")
	// 	}
	// }()

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("Current date and time:", currentTime)

	fmt.Println()

	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1)/WisdomDB")
	if err != nil {
		panic("Failed to open the DataBase")
	}

	if err := db.Ping(); err != nil {
		log.Fatal("The database is not accessible")
	}

	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/Wisdom", WisdomPoster).Methods("POST")
	router.HandleFunc("/Wisdom", WisdomGetter).Methods("GET")
	router.HandleFunc("/Wisdom/{WisdomID}", GetWisdomByID).Methods("GET")

	// Start the server after all routes are configured
	http.ListenAndServe(":3000", router)
	print("finished")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Handler")
}
