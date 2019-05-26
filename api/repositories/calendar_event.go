package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ICalendarEventRepo interface {
	List(uid bson.ObjectId) ([]entities.CalendarEvent, error)
	Add(uid bson.ObjectId, e *entities.CalendarEvent) (*entities.CalendarEvent, error)
	Edit(uid bson.ObjectId, id bson.ObjectId, e *entities.CalendarEvent) (*entities.CalendarEvent, error)
	Remove(uid bson.ObjectId, id bson.ObjectId) error
}

type CalendarEventRepo struct {
	c *mgo.Collection
}

func NewCalendarEventRepo(db *mgo.Database) *CalendarEventRepo {
	return &CalendarEventRepo{db.C("calendar_event")}
}

func (r *CalendarEventRepo) List(uid bson.ObjectId) ([]entities.CalendarEvent, error) {
	var events []entities.CalendarEvent
	if err := r.c.Find(bson.M{"user_id": uid}).All(&events); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *CalendarEventRepo) Add(uid bson.ObjectId, e *entities.CalendarEvent) (*entities.CalendarEvent, error) {
	eNew := &entities.CalendarEvent{
		Id:          bson.NewObjectId(),
		UserId:      uid,
		Name:        e.Name,
		StartDt:     e.StartDt,
		EndDt:       e.EndDt,
		Description: e.Description,
	}

	if err := r.c.Insert(eNew); err != nil {
		return nil, err
	}

	return eNew, nil
}

func (r *CalendarEventRepo) Edit(uid bson.ObjectId, id bson.ObjectId, e *entities.CalendarEvent) (*entities.CalendarEvent, error) {
	mNew := bson.M{
		"name":        e.Name,
		"start_dt":    e.StartDt,
		"end_dt":      e.EndDt,
		"description": e.Description,
	}

	if err := r.c.Update(bson.M{"_id": id, "user_id": uid}, bson.M{"$set": mNew}); err != nil {
		return nil, err
	}

	var eNew entities.CalendarEvent
	if err := r.c.Find(bson.M{"_id": id, "user_id": uid}).One(&eNew); err != nil {
		return nil, err
	}

	return &eNew, nil
}

func (r *CalendarEventRepo) Remove(uid bson.ObjectId, id bson.ObjectId) error {
	if err := r.c.Remove(bson.M{"_id": id, "user_id": uid}); err != nil {
		return err
	}

	return nil
}
