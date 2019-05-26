package routes

import (
	"net/http"

	"bitbucket.org/aj5110/tpo-lp4/api/handlers/auth"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/calendar"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/event"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/timetable"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/todo"
	"bitbucket.org/aj5110/tpo-lp4/api/middleware"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	initStaticRoutes(r)
	initApiRoutes(r)

	return r
}

func initStaticRoutes(r *mux.Router) {
	r.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods(http.MethodGet)
	r.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./client/dist/js/"))))
	r.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./client/dist/css/"))))
}

func initApiRoutes(r *mux.Router) {
	apiRouter := r.PathPrefix("/api/").Subrouter()

	// auth
	authRouter := apiRouter.PathPrefix("/auth/").Subrouter()
	authRouter.HandleFunc("/register/", auth.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/signin/", auth.Signin).Methods(http.MethodPost)
	authRouter.Handle("/admin/register/", middleware.CheckAdmin(http.HandlerFunc(auth.AdminRegister))).Methods(http.MethodPost)

	// todo
	todoRouter := r.PathPrefix("/todo/").Subrouter()
	todoRouter.Use(middleware.CheckUser)
	todoRouter.HandleFunc("/", todo.List).Methods(http.MethodGet)
	todoRouter.HandleFunc("/", todo.Add).Methods(http.MethodPost)
	todoRouter.HandleFunc("/{id}", todo.Edit).Methods(http.MethodPut)
	todoRouter.HandleFunc("/{id}", todo.Remove).Methods(http.MethodDelete)

	// calendar
	calendarRouter := r.PathPrefix("/calendar-event/").Subrouter()
	calendarRouter.Use(middleware.CheckUser)
	calendarRouter.HandleFunc("/", calendar.List).Methods(http.MethodGet)
	calendarRouter.HandleFunc("/", calendar.Add).Methods(http.MethodPost)
	calendarRouter.HandleFunc("/{id}", calendar.Edit).Methods(http.MethodPut)
	calendarRouter.HandleFunc("/{id}", calendar.Remove).Methods(http.MethodDelete)

	// timetable
	timetableRouter := r.PathPrefix("/course/").Subrouter()
	timetableRouter.Use(middleware.CheckUser)
	timetableRouter.HandleFunc("/", timetable.List).Methods(http.MethodGet)
	timetableRouter.HandleFunc("/", timetable.Add).Methods(http.MethodPost)
	timetableRouter.HandleFunc("/{id}", timetable.Edit).Methods(http.MethodPut)
	timetableRouter.HandleFunc("/{id}", timetable.Remove).Methods(http.MethodDelete)

	// event
	eventRouter := apiRouter.PathPrefix("/event/").Subrouter()
	eventRouter.Handle("/", middleware.Protect(http.HandlerFunc(event.List))).Methods(http.MethodGet)
	eventRouter.Handle("/", middleware.CheckManager(http.HandlerFunc(event.Add))).Methods(http.MethodPost)
	eventRouter.Handle("/{id}", middleware.CheckManager(http.HandlerFunc(event.Edit))).Methods(http.MethodPut)
	eventRouter.Handle("/{id}", middleware.CheckManager(http.HandlerFunc(event.Remove))).Methods(http.MethodDelete)

}
