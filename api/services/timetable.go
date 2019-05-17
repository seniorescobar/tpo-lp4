package services

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
	"gopkg.in/mgo.v2/bson"
)

type TimetableService struct {
	calendarRepo repositories.ITimetableRepo
}

func NewTimetableService(calendarRepo repositories.ITimetableRepo) *TimetableService {
	return &TimetableService{calendarRepo}
}

func (s *TimetableService) List() ([]entities.Course, error) {
	return s.calendarRepo.List()
}

func (s *TimetableService) Add(e *entities.Course) (*entities.Course, error) {
	return s.calendarRepo.Add(e)
}

func (s *TimetableService) Edit(id bson.ObjectId, e *entities.Course) (*entities.Course, error) {
	return s.calendarRepo.Edit(id, e)
}

func (s *TimetableService) Remove(id bson.ObjectId) error {
	return s.calendarRepo.Remove(id)
}
