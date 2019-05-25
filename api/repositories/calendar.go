package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ICalendarRepo interface {
	List(uid bson.ObjectId) ([]entities.Event, error)
	Add(uid bson.ObjectId, e *entities.Event) (*entities.Event, error)
	Edit(uid bson.ObjectId, id bson.ObjectId, e *entities.Event) (*entities.Event, error)
	Remove(uid bson.ObjectId, id bson.ObjectId) error
}

type CalendarRepo struct {
	db *mgo.Database
}

func NewCalendarRepo(db *mgo.Database) *CalendarRepo {
	return &CalendarRepo{db}
}

func (r *CalendarRepo) List(uid bson.ObjectId) ([]entities.Event, error) {
	var events []entities.Event
	if err := r.db.C("calendar").Find(bson.M{"user_id": uid}).All(&events); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *CalendarRepo) Add(uid bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	eNew := &entities.Event{
		Id:          bson.NewObjectId(),
		UserId:      uid,
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

func (r *CalendarRepo) Edit(uid bson.ObjectId, id bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	mNew := bson.M{
		"name":        e.Name,
		"start_dt":    e.StartDt,
		"end_dt":      e.EndDt,
		"description": e.Description,
	}

	if err := r.db.C("calendar").Update(bson.M{"_id": id, "user_id": uid}, bson.M{"$set": mNew}); err != nil {
		return nil, err
	}

	var eNew entities.Event
	if err := r.db.C("calendar").Find(bson.M{"_id": id, "user_id": uid}).One(&eNew); err != nil {
		return nil, err
	}

	return &eNew, nil
}

func (r *CalendarRepo) Remove(uid bson.ObjectId, id bson.ObjectId) error {
	if err := r.db.C("calendar").Remove(bson.M{"_id": id, "user_id": uid}); err != nil {
		return err
	}

	return nil
}
