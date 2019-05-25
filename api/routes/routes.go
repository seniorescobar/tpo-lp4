package routes

import (
	"net/http"

	"bitbucket.org/aj5110/tpo-lp4/api/handlers/auth"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/calendar"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/timetable"
	"bitbucket.org/aj5110/tpo-lp4/api/handlers/todo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// temporary workaround
	r.Use(handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}),
		handlers.AllowedOrigins([]string{"*"}),
	))

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

	todo.SetTodoHandler(apiRouter)
	calendar.SetCalendarHandler(apiRouter)
	timetable.SetTimetableHandler(apiRouter)
	auth.SetAuthHandler(apiRouter)
}
