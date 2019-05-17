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

func (s *CalendarService) List() ([]entities.Event, error) {
	return s.calendarRepo.List()
}

func (s *CalendarService) Add(e *entities.Event) (*entities.Event, error) {
	return s.calendarRepo.Add(e)
}

func (s *CalendarService) Edit(id bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	return s.calendarRepo.Edit(id, e)
}

func (s *CalendarService) Remove(id bson.ObjectId) error {
	return s.calendarRepo.Remove(id)
}
