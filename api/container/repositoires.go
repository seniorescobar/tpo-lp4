package container

import "bitbucket.org/aj5110/tpo-lp4/api/repositories"

var (
	TodoRepo   repositories.ITodoRepo
	EventRepo  repositories.IEventRepo
	CourseRepo repositories.ICourseRepo
	UserRepo   repositories.IUserRepo
)

func initRepos() {
	TodoRepo = repositories.NewTodoRepo(DB)
	EventRepo = repositories.NewEventRepo(DB)
	CourseRepo = repositories.NewCourseRepo(DB)
	UserRepo = repositories.NewUserRepo(DB)
}
