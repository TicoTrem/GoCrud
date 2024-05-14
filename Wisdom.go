package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Wisdom struct {
	Source   string
	Wisdom   string
	Datetime string
}

func WisdomGetter(w http.ResponseWriter, r *http.Request) {

	GetWisdomByID(w, r)
	rows, err := db.Query("SELECT * FROM wisdom")
	if err != nil {
		panic(err)
	}

	var id int
	var source string
	var wisdom string
	var date string
	for rows.Next() {
		err = rows.Scan(&id, &source, &wisdom, &date)
		if err != nil {
			panic(err)
		}
		print("Source is: ", source, "\nWisdom is: ", wisdom, "\n", "Date is: ", date, "\n")
	}

}

func GetWisdomByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["WisdomID"]

	myWisdom := Wisdom{}

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM wisdom WHERE id == %v;", id))
	if err != nil {
		panic(err)
	}

	var wisID int
	err = rows.Scan(&wisID, &myWisdom.Source, &myWisdom.Datetime)

	b, err := json.Marshal(myWisdom)
	w.Body = b

}

func WisdomPoster(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM wisdom")
	if err != nil {
		panic(err)
	}

	// decode json
	var json_data Wisdom
	// maps the json data to the Wisdom
	json.NewDecoder(r.Body).Decode(&json_data)

	query := fmt.Sprintf("INSERT INTO wisdom (source, wisdom, datetime) VALUES (\"%v\", \"%v\", \"%v\")", json_data.Source, json_data.Wisdom, time.Now().Format("2006-01-02 15:04:05"))
	print(query)
	// we don't use := here because we are not initializing any new variables
	_, err = db.Query(query)
	if err != nil {
		panic(err)
	}

	print(rows)
}
