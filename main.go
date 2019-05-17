package main

import (
	"log"
	"net/http"

	"bitbucket.org/aj5110/tpo-lp4/api/handlers/todo"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
	"bitbucket.org/aj5110/tpo-lp4/api/services"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("mongodb")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	if err := session.Ping(); err != nil {
		log.Fatal(err)
	}

	// db
	db := session.DB("tpo")

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

	// client static files
	r.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods(http.MethodGet)
	r.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./client/dist/js/"))))
	r.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./client/dist/css/"))))

	// api router
	apiRouter := r.PathPrefix("/api/").Subrouter()
	todo.SetTodoHandler(apiRouter, todoService)

	// serve
	log.Println("listening  on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
