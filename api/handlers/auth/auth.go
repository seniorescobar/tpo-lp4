package auth

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"bitbucket.org/aj5110/tpo-lp4/api/container"
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secret = []byte("secret")

func SetAuthHandler(r *mux.Router) {
	rt := r.PathPrefix("/auth/").Subrouter()

	rt.HandleFunc("/register/", register).Methods(http.MethodPost)
	rt.HandleFunc("/signin/", signin).Methods(http.MethodPost)
}

func register(w http.ResponseWriter, req *http.Request) {
	u := new(entities.User)
	if err := json.NewDecoder(req.Body).Decode(u); err != nil {
		log.WithField("err", err).Error("error decoding user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uNew, err := container.AuthService.Register(u)
	if err != nil {
		log.WithField("err", err).Error("error registering user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenString, err := generateToken(uNew)
	if err != nil {
		log.WithField("err", err).Error("error signing token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tNew, err := json.Marshal(struct {
		Token string `json:"token"`
	}{tokenString})

	if err != nil {
		log.WithField("err", err).Error("error encoding token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tNew)
}

func signin(w http.ResponseWriter, req *http.Request) {
	a := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(req.Body).Decode(a); err != nil {
		log.WithField("err", err).Error("error decoding auth")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := container.AuthService.Signin(a.Email, a.Password)
	if err != nil {
		log.WithField("err", err).Error("error signing in")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString, err := generateToken(user)
	if err != nil {
		log.WithField("err", err).Error("error signing token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tNew, err := json.Marshal(struct {
		Token string `json:"token"`
	}{tokenString})

	if err != nil {
		log.WithField("err", err).Error("error encoding token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tNew)
}

func generateToken(u *entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.Id,
		"email": u.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(secret)
}
