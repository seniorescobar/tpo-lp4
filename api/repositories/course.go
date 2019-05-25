package repositories

import (
	"context"

	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/middleware"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ICourseRepo interface {
	List(ctx context.Context) ([]entities.Course, error)
	Add(ctx context.Context, c *entities.Course) (*entities.Course, error)
	Edit(ctx context.Context, id bson.ObjectId, c *entities.Course) (*entities.Course, error)
	Remove(ctx context.Context, id bson.ObjectId) error
}

type CourseRepo struct {
	c *mgo.Collection
}

func NewCourseRepo(db *mgo.Database) *CourseRepo {
	return &CourseRepo{db.C("course")}
}

func (r *CourseRepo) List(ctx context.Context) ([]entities.Course, error) {
	uid, err := middleware.GetUID(ctx)
	if err != nil {
		return nil, err
	}

	var courses []entities.Course
	if err := r.c.Find(bson.M{"user_id": uid}).All(&courses); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *CourseRepo) Add(ctx context.Context, e *entities.Course) (*entities.Course, error) {
	uid, err := middleware.GetUID(ctx)
	if err != nil {
		return nil, err
	}

	cNew := &entities.Course{
		Id:       bson.NewObjectId(),
		UserId:   uid,
		Name:     e.Name,
		Duration: e.Duration,
		Color:    e.Color,
	}

	if err := r.c.Insert(cNew); err != nil {
		return nil, err
	}

	return cNew, nil
}

func (r *CourseRepo) Edit(ctx context.Context, id bson.ObjectId, c *entities.Course) (*entities.Course, error) {
	uid, err := middleware.GetUID(ctx)
	if err != nil {
		return nil, err
	}

	mNew := bson.M{
		"name":     c.Name,
		"duration": c.Duration,
		"color":    c.Color,
	}

	if err := r.c.Update(bson.M{"_id": id, "user_id": uid}, bson.M{"$set": mNew}); err != nil {
		return nil, err
	}

	var cNew entities.Course
	if err := r.c.Find(bson.M{"_id": id, "user_id": uid}).One(&cNew); err != nil {
		return nil, err
	}

	return &cNew, nil
}

func (r *CourseRepo) Remove(ctx context.Context, id bson.ObjectId) error {
	uid, err := middleware.GetUID(ctx)
	if err != nil {
		return err
	}

	if err := r.c.Remove(bson.M{"_id": id, "user_id": uid}); err != nil {
		return err
	}

	return nil
}
