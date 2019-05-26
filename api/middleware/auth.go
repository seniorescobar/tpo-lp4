package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

var secret = []byte("secret")

func CheckUser(next http.Handler) http.Handler {
	return Protect(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := checkHelper("user", req.Context()); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, req)
	}))
}

func CheckManager(next http.Handler) http.Handler {
	return Protect(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := checkHelper("manager", req.Context()); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, req)
	}))
}

func CheckAdmin(next http.Handler) http.Handler {
	return Protect(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := checkHelper("admin", req.Context()); err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, req)
	}))
}

func checkHelper(role string, ctx context.Context) error {
	u, err := GetUser(ctx)
	if err != nil {
		return err
	}

	if u.Role != role {
		return fmt.Errorf("user's role is %s, not %s", u.Role, role)
	}

	return nil
}

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

		role, ok := claims["role"].(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), "user", entities.User{
			Id:    bson.ObjectIdHex(id),
			Email: email,
			Role:  role,
		}))

		next.ServeHTTP(w, req)
	})
}

// helpers

func GetUser(ctx context.Context) (*entities.User, error) {
	v := ctx.Value("user")
	u, ok := v.(entities.User)
	if !ok {
		return nil, errors.New("error asserting ctx user")
	}

	return &u, nil
}

func GetUID(ctx context.Context) (bson.ObjectId, error) {
	u, err := GetUser(ctx)
	if err != nil {
		return "", err
	}
	return u.Id, nil
}
