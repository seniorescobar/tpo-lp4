package event

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	"bitbucket.org/aj5110/tpo-lp4/api/container"
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/helpers"
)

func List(w http.ResponseWriter, req *http.Request) {
	list, err := container.EventService.List()
	if err != nil {
		log.WithField("err", err).Error("error listing events")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	listJ, err := json.Marshal(list)
	if err != nil {
		log.WithField("err", err).Error("error encoding events")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(listJ)
}

func Add(w http.ResponseWriter, req *http.Request) {
	e := new(entities.Event)
	if err := json.NewDecoder(req.Body).Decode(e); err != nil {
		log.WithField("err", err).Error("error decoding event to add")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	eNew, err := container.EventService.Add(e)
	if err != nil {
		log.WithFields(log.Fields{
			"err":   err,
			"event": e,
		}).Error("error adding event")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	eNewJ, err := json.Marshal(eNew)
	if err != nil {
		log.WithFields(log.Fields{
			"err":       err,
			"event new": eNew,
		}).Error("error encoding added event")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(eNewJ)
}

func Edit(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	oid, err := helpers.ObjectIdHex(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Error("error decoding event id to edit")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e := new(entities.Event)
	if err := json.NewDecoder(req.Body).Decode(e); err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Error("error decoding event to edit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	eNew, err := container.EventService.Edit(oid, e)
	if err == mgo.ErrNotFound {
		log.WithField("err", err).Error("error editing event")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.WithField("err", err).Error("error editing event")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	eNewJ, err := json.Marshal(eNew)
	if err != nil {
		log.WithFields(log.Fields{
			"err":          err,
			"calendar new": eNew,
		}).Error("error encoding edited event")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(eNewJ)
}

func Remove(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	oid, err := helpers.ObjectIdHex(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Error("error decoding event to remove")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := container.EventService.Remove(oid); err == mgo.ErrNotFound {
		log.WithField("err", err).Error("error removing event")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.WithField("err", err).Error("error removing event")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
