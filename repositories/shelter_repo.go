package repositories

import (
	"PawInHand/generated/dao"
	"PawInHand/generated/dao/model"
	"errors"

	"gorm.io/gorm/clause"

	"gorm.io/gen"
	"gorm.io/gorm"
)

type ShelterRepo interface {
	FindByID(id string) (*model.Shelter, error)
	Register(shelter *model.Shelter) (*model.Shelter, error)
	UpdateByID(id string, shelter *model.Shelter) (*model.Shelter, error)
	DeleteByID(id string) (gen.ResultInfo, error)
	GetAll() ([]*model.Shelter, error)
}

type shelterRepo struct {
	q *dao.Query
}

func NewShelterRepo(query *dao.Query) ShelterRepo {
	return &shelterRepo{q: query}
}

func (self *shelterRepo) FindByID(id string) (*model.Shelter, error) {
	return self.q.Shelter.Where(
		self.q.Shelter.ID.Eq(id),
	).First()
}

func (self *shelterRepo) DeleteByID(id string) (gen.ResultInfo, error) {
	return self.q.Shelter.Where(self.q.Shelter.ID.Eq(id)).Delete()
}

func (self *shelterRepo) GetAll() ([]*model.Shelter, error) {
	return self.q.Shelter.Find()
}

func (self *shelterRepo) Register(shelter *model.Shelter) (*model.Shelter, error) {
	return shelter, self.q.Shelter.Clauses(clause.OnConflict{
		OnConstraint: "unique_Shelter",
		UpdateAll:    true,
	}).Create(shelter)
}

func (self *shelterRepo) UpdateByID(id string, shelter *model.Shelter) (*model.Shelter, error) {
	_, err := self.q.Shelter.Where(
		self.q.Shelter.ID.Eq(shelter.ID)).
		Updates(&shelter)
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return nil, gorm.ErrForeignKeyViolated
	}
	return shelter, err
}
