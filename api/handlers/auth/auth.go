package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/services"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secret = []byte("secret")

type AuthHandler struct {
	authService *services.AuthService
}

func SetAuthHandler(r *mux.Router, authService *services.AuthService) {
	h := AuthHandler{
		authService,
	}

	rt := r.PathPrefix("/auth/").Subrouter()

	rt.HandleFunc("/signin/", h.signin).Methods(http.MethodPost)
}

func (h *AuthHandler) signin(w http.ResponseWriter, req *http.Request) {
	a := new(entities.Auth)
	if err := json.NewDecoder(req.Body).Decode(a); err != nil {
		log.WithField("err", err).Error("error decoding auth")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.authService.Signin(a.Email, a.Password)
	if err != nil {
		log.WithField("err", err).Error("error signing in")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.WithField("err", err).Error("error signing token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	w.WriteHeader(http.StatusOK)
}
