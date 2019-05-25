package repositories

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ITodoRepo interface {
	List(uid bson.ObjectId) ([]entities.Todo, error)
	Add(uid bson.ObjectId, t *entities.Todo) (*entities.Todo, error)
	Edit(uid bson.ObjectId, id bson.ObjectId, t *entities.Todo) (*entities.Todo, error)
	Remove(uid bson.ObjectId, id bson.ObjectId) error
}

type TodoRepo struct {
	c *mgo.Collection
}

func NewTodoRepo(db *mgo.Database) *TodoRepo {
	return &TodoRepo{db.C("todo")}
}

func (r *TodoRepo) List(uid bson.ObjectId) ([]entities.Todo, error) {
	var todos []entities.Todo
	if err := r.c.Find(bson.M{"user_id": uid}).All(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepo) Add(uid bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
	tNew := &entities.Todo{
		Id:          bson.NewObjectId(),
		UserId:      uid,
		Description: t.Description,
	}

	if err := r.c.Insert(tNew); err != nil {
		return nil, err
	}

	return tNew, nil
}

func (r *TodoRepo) Edit(uid bson.ObjectId, id bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
	if err := r.c.Update(bson.M{"_id": id, "user_id": uid}, bson.M{"$set": bson.M{"description": t.Description}}); err != nil {
		return nil, err
	}

	var tNew entities.Todo
	if err := r.c.Find(bson.M{"_id": id, "user_id": uid}).One(&tNew); err != nil {
		return nil, err
	}

	return &tNew, nil
}

func (r *TodoRepo) Remove(uid bson.ObjectId, id bson.ObjectId) error {
	if err := r.c.Remove(bson.M{"_id": id, "user_id": uid}); err != nil {
		return err
	}

	return nil
}

// MOCK

// type TodoRepoMock struct {
// 	mock.Mock
// }

// func NewTodoRepoMock() *TodoRepoMock {
// 	return new(TodoRepoMock)
// }

// func (m *TodoRepoMock) List() ([]entities.Todo, error) {
// 	args := m.Called()
// 	return args.Get(0).([]entities.Todo), args.Error(1)
// }

// func (m *TodoRepoMock) Add(t *entities.Todo) (*entities.Todo, error) {
// 	args := m.Called(t)
// 	return args.Get(0).(*entities.Todo), args.Error(1)
// }

// func (m *TodoRepoMock) Edit(id bson.ObjectId, t *entities.Todo) (*entities.Todo, error) {
// 	args := m.Called(id, t)
// 	return args.Get(0).(*entities.Todo), args.Error(1)
// }

// func (m *TodoRepoMock) Remove(id bson.ObjectId) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }
