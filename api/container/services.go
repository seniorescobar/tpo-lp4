package container

import (
	"bitbucket.org/aj5110/tpo-lp4/api/services"
)

var (
	TodoService      *services.TodoService
	CalendarService  *services.CalendarService
	TimetableService *services.TimetableService
	AuthService      *services.AuthService
)

func initServices() {
	TodoService = services.NewTodoService(TodoRepo)
	CalendarService = services.NewCalendarService(EventRepo)
	TimetableService = services.NewTimetableService(CourseRepo)
	AuthService = services.NewAuthService(UserRepo)
}
