package services

import (
	"context"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
	"gopkg.in/mgo.v2/bson"
)

type TimetableService struct {
	calendarRepo repositories.ICourseRepo
}

func NewTimetableService(calendarRepo repositories.ICourseRepo) *TimetableService {
	return &TimetableService{calendarRepo}
}

func (s *TimetableService) List(ctx context.Context) ([]entities.Course, error) {
	return s.calendarRepo.List(ctx)
}

func (s *TimetableService) Add(ctx context.Context, e *entities.Course) (*entities.Course, error) {
	return s.calendarRepo.Add(ctx, e)
}

func (s *TimetableService) Edit(ctx context.Context, id bson.ObjectId, e *entities.Course) (*entities.Course, error) {
	return s.calendarRepo.Edit(ctx, id, e)
}

func (s *TimetableService) Remove(ctx context.Context, id bson.ObjectId) error {
	return s.calendarRepo.Remove(ctx, id)
}
