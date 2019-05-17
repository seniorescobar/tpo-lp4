package todo

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/helpers"
	"bitbucket.org/aj5110/tpo-lp4/api/services"

	log "github.com/sirupsen/logrus"
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
		log.WithField("err", err).Error("error listing todos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(list); err != nil {
		log.WithField("err", err).Error("error encoding todos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandler) add(w http.ResponseWriter, req *http.Request) {
	t := new(entities.Todo)
	if err := json.NewDecoder(req.Body).Decode(t); err != nil {
		log.WithField("err", err).Error("error decoding todo to add")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tNew, err := h.todoService.Add(t)
	if err != nil {
		log.WithFields(log.Fields{
			"err":  err,
			"todo": t,
		}).Error("error adding todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tNew); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"todo new": tNew,
		}).Error("error encoding added todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandler) edit(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	oid, err := helpers.ObjectIdHex(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Error("error decoding todo id to edit")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := new(entities.Todo)
	if err := json.NewDecoder(req.Body).Decode(t); err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Error("error decoding todo to edit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tNew, err := h.todoService.Edit(oid, t)
	if err == mgo.ErrNotFound {
		log.WithField("err", err).Error("error editing todo")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.WithField("err", err).Error("error editing todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tNew); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"todo new": tNew,
		}).Error("error encoding edited todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandler) remove(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	oid, err := helpers.ObjectIdHex(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Error("error decoding todo to remove")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.todoService.Remove(oid); err == mgo.ErrNotFound {
		log.WithField("err", err).Error("error removing todo")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.WithField("err", err).Error("error removing todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
