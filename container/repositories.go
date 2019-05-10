package container

import "bitbucket.org/aj5110/tpo-lp4/repositories"

var GetUserRepo = func() repositories.IUserRepo {
	return new(repositories.UserRepo)
}
