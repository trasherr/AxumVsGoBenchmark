package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB
var userModel *UserModel

func main() {
	var err error

	connStr = "postgres://postgres:postgres@localhost:5432/benchmark_db?sslmode=disable"
	
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	userModel = &UserModel{DB: db}


	mux := http.NewServeMux()
	mux.HandleFunc("GET /", getAllUsers)
	mux.HandleFunc("GET /user", login)
	mux.HandleFunc("POST /user", register)
	mux.HandleFunc("PUT /user", authMiddleware(updateUser))
	mux.HandleFunc("DELETE /user", authMiddleware(deleteUser))

	http.ListenAndServe("localhost:3000", mux)
}
