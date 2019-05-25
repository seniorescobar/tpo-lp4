package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

var secret = []byte("secret")

func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var tknString string
		authHeader := req.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tknString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		claims := make(jwt.MapClaims)

		tkn, err := jwt.ParseWithClaims(tknString, claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err == jwt.ErrSignatureInvalid {
			log.WithField("err", err).Error("error parsing claims")
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			log.WithField("err", err).Error("error parsing claims")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		id, ok := claims["id"].(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), "user", entities.User{
			Id:    bson.ObjectIdHex(id),
			Email: email,
		}))

		next.ServeHTTP(w, req)
	})
}

// helpers

func GetUID(req *http.Request) (bson.ObjectId, error) {
	v := req.Context().Value("user")
	u, ok := v.(entities.User)
	if !ok {
		return "", errors.New("error asserting ctx user")
	}

	return u.Id, nil
}
