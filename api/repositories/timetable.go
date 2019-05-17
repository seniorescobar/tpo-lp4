package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"github.com/stretchr/testify/mock"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ITimetableRepo interface {
	List() ([]entities.Course, error)
	Add(*entities.Course) (*entities.Course, error)
	Edit(bson.ObjectId, *entities.Course) (*entities.Course, error)
	Remove(bson.ObjectId) error
}

type TimetableRepo struct {
	db *mgo.Database
}

func NewTimetableRepo(db *mgo.Database) *TimetableRepo {
	return &TimetableRepo{db}
}

func (r *TimetableRepo) List() ([]entities.Course, error) {
	var courses []entities.Course
	if err := r.db.C("timetable").Find(nil).All(&courses); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *TimetableRepo) Add(e *entities.Course) (*entities.Course, error) {
	cNew := &entities.Course{
		Id:       bson.NewObjectId(),
		Name:     e.Name,
		Duration: e.Duration,
		Color:    e.Color,
	}

	if err := r.db.C("timetable").Insert(cNew); err != nil {
		return nil, err
	}

	return cNew, nil
}

func (r *TimetableRepo) Edit(id bson.ObjectId, e *entities.Course) (*entities.Course, error) {
	cNew := &entities.Course{
		Id:       id,
		Name:     e.Name,
		Duration: e.Duration,
		Color:    e.Color,
	}

	if err := r.db.C("timetable").UpdateId(id, cNew); err != nil {
		return nil, err
	}

	return cNew, nil
}

func (r *TimetableRepo) Remove(id bson.ObjectId) error {
	if err := r.db.C("timetable").RemoveId(id); err != nil {
		return err
	}

	return nil
}

// MOCK

type TimetableRepoMock struct {
	mock.Mock
}

func NewTimetableRepoMock() *TimetableRepoMock {
	return new(TimetableRepoMock)
}

func (m *TimetableRepoMock) List() ([]entities.Course, error) {
	args := m.Called()
	return args.Get(0).([]entities.Course), args.Error(1)
}

func (m *TimetableRepoMock) Add(e *entities.Course) (*entities.Course, error) {
	args := m.Called(e)
	return args.Get(0).(*entities.Course), args.Error(1)
}

func (m *TimetableRepoMock) Edit(id bson.ObjectId, e *entities.Course) (*entities.Course, error) {
	args := m.Called(id, e)
	return args.Get(0).(*entities.Course), args.Error(1)
}

func (m *TimetableRepoMock) Remove(id bson.ObjectId) error {
	args := m.Called(id)
	return args.Error(0)
}
