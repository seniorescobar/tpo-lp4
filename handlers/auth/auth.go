package auth

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/aj5110/tpo-lp4/container"
)

func Routes() {
	http.HandleFunc("/register/", register)
}

func register(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var rf struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(req.Body).Decode(&rf); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := container.GetAuthService().RegisterUser(rf.Email, rf.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
