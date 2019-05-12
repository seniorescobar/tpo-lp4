package todo

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"bitbucket.org/aj5110/tpo-lp4/entities"
	"bitbucket.org/aj5110/tpo-lp4/services"
)

type TodoHandler struct {
	todoService *services.TodoService
}

func SetTodoHandler(r *mux.Router, todoService *services.TodoService) {
	h := TodoHandler{
		todoService,
	}

	r.HandleFunc("/todo/", h.list).Methods(http.MethodGet)
	r.HandleFunc("/todo/", h.add).Methods(http.MethodPost)
	r.HandleFunc("/todo/", h.edit).Methods(http.MethodPut)
	r.HandleFunc("/todo/", h.del).Methods(http.MethodDelete)
}

func (h *TodoHandler) list(w http.ResponseWriter, req *http.Request) {
	list, err := h.todoService.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(w)
	if err := e.Encode(list); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) add(w http.ResponseWriter, req *http.Request) {
	t := new(entities.Todo)
	if err := json.NewDecoder(req.Body).Decode(t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.todoService.Add(t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO
func (h *TodoHandler) edit(w http.ResponseWriter, req *http.Request) {}

// TODO
func (h *TodoHandler) del(w http.ResponseWriter, req *http.Request) {}
