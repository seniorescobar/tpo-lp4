package auth

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"bitbucket.org/aj5110/tpo-lp4/api/container"
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

var secret = []byte("secret")

func Register(w http.ResponseWriter, req *http.Request) {
	u := new(entities.User)
	if err := json.NewDecoder(req.Body).Decode(u); err != nil {
		log.WithField("err", err).Error("error decoding user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u.Role = "user"

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

func Signin(w http.ResponseWriter, req *http.Request) {
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

func ChangePassword(w http.ResponseWriter, req *http.Request) {
	user, err := middleware.GetUser(req.Context())
	if err != nil {
		log.WithField("err", err).Error("error getting user from ctx")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	p := struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}{}

	if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
		log.WithField("err", err).Error("error decoding passwords")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := container.AuthService.ChangePassword(user.Email, p.OldPassword, p.NewPassword); err != nil {
		log.WithField("err", err).Error("error changing password")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func generateToken(u *entities.User) (string, error) {
	// TODO map -> struct
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.Id,
		"email": u.Email,
		"role":  u.Role,
		// "exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(secret)
}
