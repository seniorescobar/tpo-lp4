package services

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
	"gopkg.in/mgo.v2/bson"
)

type EventService struct {
	eventRepo repositories.IEventRepo
}

func NewEventService(eventRepo repositories.IEventRepo) *EventService {
	return &EventService{eventRepo}
}

func (s *EventService) List() ([]entities.Event, error) {
	return s.eventRepo.List()
}

func (s *EventService) Add(e *entities.Event) (*entities.Event, error) {
	return s.eventRepo.Add(e)
}

func (s *EventService) Edit(id bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	return s.eventRepo.Edit(id, e)
}

func (s *EventService) Remove(id bson.ObjectId) error {
	return s.eventRepo.Remove(id)
}
