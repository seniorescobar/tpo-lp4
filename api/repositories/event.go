package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IEventRepo interface {
	List() ([]entities.Event, error)
	Add(e *entities.Event) (*entities.Event, error)
	Edit(id bson.ObjectId, e *entities.Event) (*entities.Event, error)
	Remove(id bson.ObjectId) error
}

type EventRepo struct {
	c *mgo.Collection
}

func NewEventRepo(db *mgo.Database) *EventRepo {
	return &EventRepo{db.C("event")}
}

func (r *EventRepo) List() ([]entities.Event, error) {
	var events []entities.Event
	if err := r.c.Find(nil).All(&events); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventRepo) Add(e *entities.Event) (*entities.Event, error) {
	eNew := &entities.Event{
		Id:          bson.NewObjectId(),
		Name:        e.Name,
		Date:        e.Date,
		Organizer:   e.Organizer,
		Description: e.Description,
	}

	if err := r.c.Insert(eNew); err != nil {
		return nil, err
	}

	return eNew, nil
}

func (r *EventRepo) Edit(id bson.ObjectId, e *entities.Event) (*entities.Event, error) {
	mNew := bson.M{
		"name":        e.Name,
		"date":        e.Date,
		"organizer":   e.Organizer,
		"description": e.Description,
	}

	if err := r.c.Update(bson.M{"_id": id}, bson.M{"$set": mNew}); err != nil {
		return nil, err
	}

	var eNew entities.Event
	if err := r.c.Find(bson.M{"_id": id}).One(&eNew); err != nil {
		return nil, err
	}

	return &eNew, nil
}

func (r *EventRepo) Remove(id bson.ObjectId) error {
	if err := r.c.Remove(bson.M{"_id": id}); err != nil {
		return err
	}

	return nil
}
