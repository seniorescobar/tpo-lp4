package main

import (
	"log"
	"net/http"

	"bitbucket.org/aj5110/tpo-lp4/api/handlers/auth"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/calendar"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/timetable"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/todo"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
	"bitbucket.org/aj5110/tpo-lp4/api/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")
	// session, err := mgo.Dial("mongodb")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	if err := session.Ping(); err != nil {
		log.Fatal(err)
	}

	// db
	db := session.DB("tpo")

	// repositories
	var (
		todoRepo   = repositories.NewTodoRepo(db)
		eventRepo  = repositories.NewEventRepo(db)
		courseRepo = repositories.NewCourseRepo(db)
		userRepo   = repositories.NewUserRepo(db)
	)

	// services
	var (
		todoService      = services.NewTodoService(todoRepo)
		calendarService  = services.NewCalendarService(eventRepo)
		timetableService = services.NewTimetableService(courseRepo)
		authService      = services.NewAuthService(userRepo)
	)

	// routes
	r := mux.NewRouter()

	// temporary workaround
	r.Use(handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}),
		handlers.AllowedOrigins([]string{"*"}),
	))

	// client static files
	r.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods(http.MethodGet)
	r.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./client/dist/js/"))))
	r.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./client/dist/css/"))))

	// api router
	apiRouter := r.PathPrefix("/api/").Subrouter()
	todo.SetTodoHandler(apiRouter, todoService)
	calendar.SetCalendarHandler(apiRouter, calendarService)
	timetable.SetTimetableHandler(apiRouter, timetableService)
	auth.SetAuthHandler(apiRouter, authService)

	// serve
	log.Println("listening  on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
