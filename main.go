package main

import (
	"database/sql"
	"log"
	"net/http"

	"bitbucket.org/aj5110/tpo-lp4/handlers/auth"
	"bitbucket.org/aj5110/tpo-lp4/repositories"
	"bitbucket.org/aj5110/tpo-lp4/services"
)

func main() {
	// db
	db, err := sql.Open("", "")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// repositories
	var (
		userRepo = repositories.NewUserRepo(db)
	)

	// services
	var (
		authService = services.NewAuthService(userRepo)
	)

	// routes
	auth.SetAuthHandler(http.DefaultServeMux, authService)

	// serve
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
