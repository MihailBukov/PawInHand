package repositories

import (
	"PawInHand/generated/dao"
	"PawInHand/generated/dao/model"
)

type EventRepo interface {
	FindByID(id string) (*model.Event, error)
	Create(event *model.Event) (*model.Event, error)
	GetAll() ([]*model.Event, error)
}

type eventRepo struct {
	q *dao.Query
}

func NewEventRepo(query *dao.Query) EventRepo {
	return &eventRepo{q: query}
}

func (self *eventRepo) FindByID(id string) (*model.Event, error) {
	return self.q.Event.Where(
		self.q.Event.ID.Eq(id),
	).First()
}

func (self *eventRepo) GetAll() ([]*model.Event, error) {
	return self.q.Event.Find()
}

func (self *eventRepo) Create(event *model.Event) (*model.Event, error) {
	return event, self.q.Event.Create(event)
}
