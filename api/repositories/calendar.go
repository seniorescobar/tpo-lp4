package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"github.com/stretchr/testify/mock"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ICalendarRepo interface {
	List() ([]entities.Event, error)
	Add(*entities.Event) (*entities.Event, error)
	Edit(bson.ObjectId, *entities.Event) (*entities.Event, error)
	Remove(bson.ObjectId) error
}

type CalendarRepo struct {
	db *mgo.Database
}

func NewCalendarRepo(db *mgo.Database) *CalendarRepo {
	return &CalendarRepo{db}
}

func (r *CalendarRepo) List() ([]entities.Event, error) {
	var events []entities.Event
	if err := r.db.C("calendar").Find(nil).All(&events); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *CalendarRepo) Add(e *entities.Event) (*entities.Event, error) {
	eNew := &entities.Event{
		Id:          bson.NewObjectId(),
		Name:        e.Name,
		StartDt:     e.StartDt,
		EndDt:       e.EndDt,
		Description: e.Description,
	}

	if err := r.db.C("calendar").Insert(eNew); err != nil {
		return nil, err
	}

	return eNew, nil
}

func (r *CalendarRepo) Edit(id bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	eNew := &entities.Event{
		Id:          id,
		Name:        e.Name,
		StartDt:     e.StartDt,
		EndDt:       e.EndDt,
		Description: e.Description,
	}

	if err := r.db.C("calendar").UpdateId(id, eNew); err != nil {
		return nil, err
	}

	return eNew, nil
}

func (r *CalendarRepo) Remove(id bson.ObjectId) error {
	if err := r.db.C("calendar").RemoveId(id); err != nil {
		return err
	}

	return nil
}

// MOCK

type CalendarRepoMock struct {
	mock.Mock
}

func NewCalendarRepoMock() *CalendarRepoMock {
	return new(CalendarRepoMock)
}

func (m *CalendarRepoMock) List() ([]entities.Event, error) {
	args := m.Called()
	return args.Get(0).([]entities.Event), args.Error(1)
}

func (m *CalendarRepoMock) Add(e *entities.Event) (*entities.Event, error) {
	args := m.Called(e)
	return args.Get(0).(*entities.Event), args.Error(1)
}

func (m *CalendarRepoMock) Edit(id bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	args := m.Called(id, e)
	return args.Get(0).(*entities.Event), args.Error(1)
}

func (m *CalendarRepoMock) Remove(id bson.ObjectId) error {
	args := m.Called(id)
	return args.Error(0)
}
