package container

import (
	"bitbucket.org/aj5110/tpo-lp4/api/services"
)

var (
	AuthService      *services.AuthService
	TodoService      *services.TodoService
	CalendarService  *services.CalendarService
	TimetableService *services.TimetableService
	EventService     *services.EventService
)

func initServices() {
	AuthService = services.NewAuthService(UserRepo)
	TodoService = services.NewTodoService(TodoRepo)
	CalendarService = services.NewCalendarService(CalendarEventRepo)
	TimetableService = services.NewTimetableService(CourseRepo)
	EventService = services.NewEventService(EventRepo)
}
