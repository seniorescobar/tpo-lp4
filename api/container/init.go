package container

import (
	"log"
)

func init() {
	// db
	log.Println("initializing db")
	initDB()

	// repositories
	log.Println("initializing repositories")
	initRepos()

	// services
	log.Println("initializing services")
	initServices()
}
