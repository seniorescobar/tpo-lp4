package main

import (
	"log"
	"net/http"
	"os"

	"bitbucket.org/aj5110/tpo-lp4/api/routes"
)

func main() {
	// serve
	port := os.Getenv("PORT")
	log.Println("listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, routes.InitRoutes()))
}
