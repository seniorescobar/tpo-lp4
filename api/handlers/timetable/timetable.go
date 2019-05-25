package timetable

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/helpers"
	"bitbucket.org/aj5110/tpo-lp4/api/middleware"
	"bitbucket.org/aj5110/tpo-lp4/api/services"
)

type TimetableHandler struct {
	timetableService *services.TimetableService
}

func SetTimetableHandler(r *mux.Router, timetableService *services.TimetableService) {
	h := TimetableHandler{
		timetableService,
	}

	rt := r.PathPrefix("/course/").Subrouter()

	rt.Use(middleware.Protect)

	rt.HandleFunc("/", h.list).Methods(http.MethodGet)
	rt.HandleFunc("/", h.add).Methods(http.MethodPost)
	rt.HandleFunc("/{id}", h.edit).Methods(http.MethodPut)
	rt.HandleFunc("/{id}", h.remove).Methods(http.MethodDelete)
}

func (h *TimetableHandler) list(w http.ResponseWriter, req *http.Request) {
	list, err := h.timetableService.List(req.Context())
	if err != nil {
		log.WithField("err", err).Error("error listing courses")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	listJ, err := json.Marshal(list)
	if err != nil {
		log.WithField("err", err).Error("error encoding courses")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(listJ)
}

func (h *TimetableHandler) add(w http.ResponseWriter, req *http.Request) {
	e := new(entities.Course)
	if err := json.NewDecoder(req.Body).Decode(e); err != nil {
		log.WithField("err", err).Error("error decoding course to add")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	eNew, err := h.timetableService.Add(req.Context(), e)
	if err != nil {
		log.WithFields(log.Fields{
			"err":    err,
			"course": e,
		}).Error("error adding course")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	eNewJ, err := json.Marshal(eNew)
	if err != nil {
		log.WithFields(log.Fields{
			"err":        err,
			"course new": eNew,
		}).Error("error encoding added course")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(eNewJ)
}

func (h *TimetableHandler) edit(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	oid, err := helpers.ObjectIdHex(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Error("error decoding course id to edit")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e := new(entities.Course)
	if err := json.NewDecoder(req.Body).Decode(e); err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Error("error decoding course to edit")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	eNew, err := h.timetableService.Edit(req.Context(), oid, e)
	if err == mgo.ErrNotFound {
		log.WithField("err", err).Error("error editing course")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.WithField("err", err).Error("error editing course")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	eNewJ, err := json.Marshal(eNew)
	if err != nil {
		log.WithFields(log.Fields{
			"err":           err,
			"timetable new": eNew,
		}).Error("error encoding edited course")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(eNewJ)
}

func (h *TimetableHandler) remove(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	oid, err := helpers.ObjectIdHex(id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"id":  id,
		}).Error("error decoding course to remove")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.timetableService.Remove(req.Context(), oid); err == mgo.ErrNotFound {
		log.WithField("err", err).Error("error removing course")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.WithField("err", err).Error("error removing course")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
