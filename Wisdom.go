package gocrud

import (
	"database/sql"
	"net/http"
	"time"
)

func WisdomHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "tyson:password@/WisdomDB")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("SELECT * FROM Wisdom")

}
