package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"bitbucket.org/aj5110/tpo-lp4/handlers/todo"
	"bitbucket.org/aj5110/tpo-lp4/repositories"
	"bitbucket.org/aj5110/tpo-lp4/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// db
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	log.Fatal()

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

	// api router
	apiRouter := r.PathPrefix("/api/").Subrouter()
	todo.SetTodoHandler(apiRouter, todoService)

	// serve
	log.Println("listening  on port 8080")
	http.ListenAndServe(":8080", r)
}
