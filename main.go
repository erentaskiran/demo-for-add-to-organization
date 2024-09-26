package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

type application struct {
	Db *sql.DB
}

func main() {
	mux := NewRouter()

	db, err := newDb()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := &application{
		Db: db,
	}

	http.ListenAndServe(":3000", mux)
	fmt.Println("server started at 3000 port." + app.Db.Ping().Error())
}
