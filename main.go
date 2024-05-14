package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type Wisdom struct {
	Source   string
	Wisdom   string
	Datetime string
}

func main() {

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("Current date and time:", currentTime)

	fmt.Println()

	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1)/WisdomDB")
	if err != nil {
		panic(err)
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
