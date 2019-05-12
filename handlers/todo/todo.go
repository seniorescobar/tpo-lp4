package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	rt := r.PathPrefix("/todo/").Subrouter()

	rt.HandleFunc("/", h.list).Methods(http.MethodGet)
	rt.HandleFunc("/", h.add).Methods(http.MethodPost)
	rt.HandleFunc("/{id}", h.edit).Methods(http.MethodPut)
	rt.HandleFunc("/", h.del).Methods(http.MethodDelete)
}

func (h *TodoHandler) list(w http.ResponseWriter, req *http.Request) {
	list, err := h.todoService.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(list); err != nil {
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

func (h *TodoHandler) edit(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := new(entities.Todo)
	if err := json.NewDecoder(req.Body).Decode(t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newTodo, err := h.todoService.Edit(id, t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(newTodo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO
func (h *TodoHandler) del(w http.ResponseWriter, req *http.Request) {}
