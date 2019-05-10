package main

import (
	"net/http"

	"bitbucket.org/aj5110/tpo-lp4/handlers/auth"
)

func main() {
	auth.Routes()

	http.ListenAndServe(":8080", nil)
}
