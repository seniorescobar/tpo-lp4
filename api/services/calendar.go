package services

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
	"gopkg.in/mgo.v2/bson"
)

type CalendarService struct {
	calendarRepo repositories.ICalendarRepo
}

func NewCalendarService(calendarRepo repositories.ICalendarRepo) *CalendarService {
	return &CalendarService{calendarRepo}
}

func (s *CalendarService) List(uid bson.ObjectId) ([]entities.Event, error) {
	return s.calendarRepo.List(uid)
}

func (s *CalendarService) Add(uid bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	return s.calendarRepo.Add(uid, e)
}

func (s *CalendarService) Edit(uid bson.ObjectId, id bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	return s.calendarRepo.Edit(uid, id, e)
}

func (s *CalendarService) Remove(uid bson.ObjectId, id bson.ObjectId) error {
	return s.calendarRepo.Remove(uid, id)
}
