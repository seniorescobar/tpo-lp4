package container

import "bitbucket.org/aj5110/tpo-lp4/services"

func GetAuthService() *services.AuthService {
	return services.NewAuthService(GetUserRepo())
}
