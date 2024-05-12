package main

import (
	"fmt"
	"net/http"
	"time"

	// the http handler I am using
	"github.com/gorilla/mux"
)

// this crud will allow you to add useful tips
func main() {
	print("test")

	// This will hold the data of our data
	type Wisdom struct {
		Source string
		Wisdom string
		date   time.Time
	}

	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("Wisdom", WisdomGetter).Methods("GET")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Handler")
}
