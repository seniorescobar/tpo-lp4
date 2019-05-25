package main

import (
	"log"
	"net/http"

	"bitbucket.org/aj5110/tpo-lp4/api/routes"
)

func main() {
	// serve
	log.Println("listening  on port 8080")
	log.Fatal(http.ListenAndServe(":8080", routes.InitRoutes()))
}
