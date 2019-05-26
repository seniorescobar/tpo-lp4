package container

import "bitbucket.org/aj5110/tpo-lp4/api/repositories"

var (
	UserRepo          repositories.IUserRepo
	TodoRepo          repositories.ITodoRepo
	CalendarEventRepo repositories.ICalendarEventRepo
	CourseRepo        repositories.ICourseRepo
	EventRepo         repositories.IEventRepo
)

func initRepos() {
	UserRepo = repositories.NewUserRepo(DB)
	TodoRepo = repositories.NewTodoRepo(DB)
	CalendarEventRepo = repositories.NewCalendarEventRepo(DB)
	CourseRepo = repositories.NewCourseRepo(DB)
	EventRepo = repositories.NewEventRepo(DB)
}
