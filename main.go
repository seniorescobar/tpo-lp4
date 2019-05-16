package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"bitbucket.org/aj5110/tpo-lp4/api/handlers/todo"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
	"bitbucket.org/aj5110/tpo-lp4/api/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// db
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	db := client.Database("tpo")

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
