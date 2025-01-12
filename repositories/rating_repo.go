package repositories

import (
	"PawInHand/generated/dao"
	"PawInHand/generated/dao/model"
	"gorm.io/gen"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_rating_repo.go . RatingRepo
type RatingRepo interface {
	FindByID(id string) (*model.Rating, error)
	Create(rating *model.Rating) (*model.Rating, error)
	GetAll() ([]*model.Rating, error)
	FindByUserID(userID string) ([]*model.Rating, error)
	DeleteByID(id string) (gen.ResultInfo, error)
}

type ratingRepo struct {
	q *dao.Query
}

func NewRatingRepo(query *dao.Query) RatingRepo {
	return &ratingRepo{q: query}
}

func (self *ratingRepo) FindByID(id string) (*model.Rating, error) {
	return self.q.Rating.Where(
		self.q.Rating.ID.Eq(id),
	).First()
}

func (self *ratingRepo) Create(rating *model.Rating) (*model.Rating, error) {
	return rating, self.q.Rating.Create(rating)
}

func (self *ratingRepo) GetAll() ([]*model.Rating, error) {
	return self.q.Rating.Find()
}

func (self *ratingRepo) FindByUserID(userID string) ([]*model.Rating, error) {
	return self.q.Rating.Where(
		self.q.Rating.UserID.Eq(userID),
	).Find()
}

func (self *ratingRepo) DeleteByID(id string) (gen.ResultInfo, error) {
	return self.q.Rating.Where(
		self.q.Rating.ID.Eq(id),
	).Delete()
}
