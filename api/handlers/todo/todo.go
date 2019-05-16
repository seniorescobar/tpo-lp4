package todo

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/services"
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
	rt.HandleFunc("/{id}", h.remove).Methods(http.MethodDelete)
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
}

func (h *TodoHandler) add(w http.ResponseWriter, req *http.Request) {
	t := new(entities.Todo)
	if err := json.NewDecoder(req.Body).Decode(t); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tNew, err := h.todoService.Add(t)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tNew); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandler) edit(w http.ResponseWriter, req *http.Request) {
	id := bson.ObjectIdHex(mux.Vars(req)["id"])
	if id == "" {
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(newTodo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandler) remove(w http.ResponseWriter, req *http.Request) {
	id := bson.ObjectIdHex(mux.Vars(req)["id"])
	if err := h.todoService.Remove(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
