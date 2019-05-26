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

	todo.SetTodoHandler(apiRouter)
	calendar.SetCalendarHandler(apiRouter)
	timetable.SetTimetableHandler(apiRouter)

	// event
	eventRouter := apiRouter.PathPrefix("/event/").Subrouter()
	eventRouter.Handle("/", middleware.Protect(http.HandlerFunc(event.List))).Methods(http.MethodGet)
	eventRouter.Handle("/", middleware.CheckManager(http.HandlerFunc(event.Add))).Methods(http.MethodPost)
	eventRouter.Handle("/{id}", middleware.CheckManager(http.HandlerFunc(event.Edit))).Methods(http.MethodPut)
	eventRouter.Handle("/{id}", middleware.CheckManager(http.HandlerFunc(event.Remove))).Methods(http.MethodDelete)

}
