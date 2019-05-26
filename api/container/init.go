package container

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	log "github.com/sirupsen/logrus"
)

func init() {
	// db
	log.Info("initializing db")
	initDB()

	// repositories
	log.Info("initializing repositories")
	initRepos()

	// services
	log.Info("initializing services")
	initServices()

	// seed db
	for _, u := range []*entities.User{
		&entities.User{Email: "admin@tpo.com", Password: "password123", Role: "admin"},
		&entities.User{Email: "manager@tpo.com", Password: "password123", Role: "manager"},
		&entities.User{Email: "user@tpo.com", Password: "password123", Role: "user"},
	} {
		uNew, err := AuthService.Register(u)
		if err != nil {
			log.WithField("err", err).Error("error registering user")
			continue
		}

		log.Info("added ", uNew.Email)
	}
}
