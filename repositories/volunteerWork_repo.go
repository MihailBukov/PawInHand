package repositories

import (
	"PawInHand/generated/dao"
	"PawInHand/generated/dao/model"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_volunteer_repo.go . VolunteerworkRepo
type VolunteerworkRepo interface {
	FindByID(id string) (*model.Volunteerwork, error)
	Register(work *model.Volunteerwork) (*model.Volunteerwork, error)
	GetAll() ([]*model.Volunteerwork, error)
}

type volunteerWorkRepo struct {
	q *dao.Query
}

func NewVolunteerworkRepo(query *dao.Query) VolunteerworkRepo {
	return &volunteerWorkRepo{q: query}
}

func (self *volunteerWorkRepo) FindByID(id string) (*model.Volunteerwork, error) {
	return self.q.Volunteerwork.Where(
		self.q.Volunteerwork.ID.Eq(id),
	).First()
}

func (self *volunteerWorkRepo) GetAll() ([]*model.Volunteerwork, error) {
	return self.q.Volunteerwork.Find()
}

func (self *volunteerWorkRepo) Register(work *model.Volunteerwork) (*model.Volunteerwork, error) {
	return work, self.q.Volunteerwork.Create(work)
}
