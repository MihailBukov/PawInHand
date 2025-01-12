package repositories

import (
	"PawInHand/generated/dao"
	"PawInHand/generated/dao/model"
	"gorm.io/gen"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_ad_repo.go . AdRepo
type AdRepo interface {
	FindByID(id string) (*model.Ad, error)
	Create(ad *model.Ad) (*model.Ad, error)
	DeleteByID(id string) (gen.ResultInfo, error)
	GetAll() ([]*model.Ad, error)
	UpdateById(ad *model.Ad) (*model.Ad, error)
}

type adRepo struct {
	q *dao.Query
}

func NewAdRepo(query *dao.Query) AdRepo {
	return &adRepo{q: query}
}

func (self *adRepo) FindByID(id string) (*model.Ad, error) {
	return self.q.Ad.Where(
		self.q.Ad.ID.Eq(id),
	).First()
}

func (self *adRepo) DeleteByID(id string) (gen.ResultInfo, error) {
	return self.q.Ad.Where(self.q.Ad.ID.Eq(id)).Delete()
}

func (self *adRepo) GetAll() ([]*model.Ad, error) {
	return self.q.Ad.Find()
}

func (self *adRepo) Create(ad *model.Ad) (*model.Ad, error) {
	return ad, self.q.Ad.Create(ad)
}

func (self *adRepo) UpdateById(ad *model.Ad) (*model.Ad, error) {
	_, err := self.q.Post.Where(
		self.q.Post.ID.Eq(ad.ID)).
		Updates(&ad)
	return ad, err
}
