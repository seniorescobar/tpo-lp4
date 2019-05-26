package todo

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	"bitbucket.org/aj5110/tpo-lp4/api/container"
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/helpers"
	"bitbucket.org/aj5110/tpo-lp4/api/middleware"
)

func List(w http.ResponseWriter, req *http.Request) {
	uid, err := middleware.GetUID(req.Context())
	if err != nil {
		log.WithField("err", err).Error("error getting uid from req ctx")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	list, err := container.TodoService.List(uid)
	if err != nil {
		log.WithField("err", err).Error("error listing todos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	listJ, err := json.Marshal(list)
	if err != nil {
		log.WithField("err", err).Error("error encoding todos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(listJ)
}

func Add(w http.ResponseWriter, req *http.Request) {
	uid, err := middleware.GetUID(req.Context())
	if err != nil {
		log.WithField("err", err).Error("error getting uid from req ctx")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := new(entities.Todo)
	if err := json.NewDecoder(req.Body).Decode(t); err != nil {
		log.WithField("err", err).Error("error decoding todo to add")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tNew, err := container.TodoService.Add(uid, t)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "todo": t}).Error("error adding todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tNewJ, err := json.Marshal(tNew)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "todo new": tNew}).Error("error encoding added todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tNewJ)
}

func Edit(w http.ResponseWriter, req *http.Request) {
	var err error

	t := new(entities.Todo)
	if err := json.NewDecoder(req.Body).Decode(t); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("error decoding todo to edit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Id, err = helpers.ObjectIdHex(mux.Vars(req)["id"])
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("error decoding todo id to edit")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.UserId, err = middleware.GetUID(req.Context())
	if err != nil {
		log.WithField("err", err).Error("error getting uid from req ctx")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := container.TodoService.Edit(t); err == mgo.ErrNotFound {
		log.WithField("err", err).Error("error editing todo")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.WithField("err", err).Error("error editing todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	tJ, err := json.Marshal(t)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "todo": t}).Error("error encoding edited todo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tJ)
}

func Remove(w http.ResponseWriter, req *http.Request) {
	uid, err := middleware.GetUID(req.Context())
	if err != nil {
		log.WithField("err", err).Error("error getting uid from req ctx")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := mux.Vars(req)["id"]

	oid, err := helpers.ObjectIdHex(id)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("error decoding todo to remove")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := container.TodoService.Remove(uid, oid); err == mgo.ErrNotFound {
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
