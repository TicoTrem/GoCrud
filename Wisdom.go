package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	// GetWisdomByID(w, r)
	rows, err := db.Query("SELECT * FROM wisdom")
	if err != nil {
		panic(err)
	}

	var id int
	var source string
	var myWisdom string
	var date string
	for rows.Next() {
		err = rows.Scan(&id, &source, &myWisdom, &date)
		if err != nil {
			panic(err)
		}
		print("Source is: ", source, "\nWisdom is: ", myWisdom, "\n", "Date is: ", date, "\n")
	}

}

func GetWisdomByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["WisdomID"]

	myWisdom := Wisdom{}

	var myQuery string = fmt.Sprintf("SELECT * FROM wisdom WHERE id = %v;", id)
	print(myQuery)
	rows, err := db.Query(myQuery)

	if err != nil {
		panic(err)
	}

	var wisID int
	if rows.Next() {
		err = rows.Scan(&wisID, &myWisdom.Wisdom, &myWisdom.Source, &myWisdom.Datetime)
	} else {
		log.Fatal("There were no rows returned for that ID")
	}

	if err != nil {
		panic("rows.Scan failed in GetWisdomByID:\n" + err.Error())
	}

	b, err := json.Marshal(myWisdom)

	if err != nil {
		panic("json.Marshal did not workkk:\n" + err.Error())
	}
	fmt.Println(b)
	w.Write(b)

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
