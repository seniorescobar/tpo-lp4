package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IEventRepo interface {
	List(uid bson.ObjectId) ([]entities.Event, error)
	Add(uid bson.ObjectId, e *entities.Event) (*entities.Event, error)
	Edit(uid bson.ObjectId, id bson.ObjectId, e *entities.Event) (*entities.Event, error)
	Remove(uid bson.ObjectId, id bson.ObjectId) error
}

type EventRepo struct {
	c *mgo.Collection
}

func NewEventRepo(db *mgo.Database) *EventRepo {
	return &EventRepo{db.C("event")}
}

func (r *EventRepo) List(uid bson.ObjectId) ([]entities.Event, error) {
	var events []entities.Event
	if err := r.c.Find(bson.M{"user_id": uid}).All(&events); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventRepo) Add(uid bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	eNew := &entities.Event{
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

func (r *EventRepo) Edit(uid bson.ObjectId, id bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	mNew := bson.M{
		"name":        e.Name,
		"start_dt":    e.StartDt,
		"end_dt":      e.EndDt,
		"description": e.Description,
	}

	if err := r.c.Update(bson.M{"_id": id, "user_id": uid}, bson.M{"$set": mNew}); err != nil {
		return nil, err
	}

	var eNew entities.Event
	if err := r.c.Find(bson.M{"_id": id, "user_id": uid}).One(&eNew); err != nil {
		return nil, err
	}

	return &eNew, nil
}

func (r *EventRepo) Remove(uid bson.ObjectId, id bson.ObjectId) error {
	if err := r.c.Remove(bson.M{"_id": id, "user_id": uid}); err != nil {
		return err
	}

	return nil
}
