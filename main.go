package main

import (
	"database/sql"
	"log"
	"net/http"

	"bitbucket.org/aj5110/tpo-lp4/handlers/todo"
	"bitbucket.org/aj5110/tpo-lp4/repositories"
	"bitbucket.org/aj5110/tpo-lp4/services"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// db
	db, err := sql.Open("postgres", "user=postgres password=docker dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// repositories
	var (
		todoRepo = repositories.NewTodoRepo(db)
	)

	// services
	var (
		todoService = services.NewTodoService(todoRepo)
	)

	// routes
	r := mux.NewRouter()
	todo.SetTodoHandler(r, todoService)

	// serve
	http.ListenAndServe(":8080", r)
}
