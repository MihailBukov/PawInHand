package repositories

import (
	"PawInHand/generated/dao"
	"PawInHand/generated/dao/model"
	"errors"

	"gorm.io/gorm/clause"

	"gorm.io/gen"
	"gorm.io/gorm"
)

type UserRepo interface {
	FindByID(id string) (*model.User, error)
	Register(user *model.User) (*model.User, error)
	UpdateByID(id string, user *model.User) (*model.User, error)
	DeleteByID(id string) (gen.ResultInfo, error)
	GetAll() ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type userRepo struct {
	q *dao.Query
}

func NewUserRepo(query *dao.Query) UserRepo {
	return &userRepo{q: query}
}

func (self *userRepo) FindByID(id string) (*model.User, error) {
	return self.q.User.Where(
		self.q.User.ID.Eq(id),
	).First()
}

func (self *userRepo) DeleteByID(id string) (gen.ResultInfo, error) {
	return self.q.User.Where(self.q.User.ID.Eq(id)).Delete()
}

func (self *userRepo) GetAll() ([]*model.User, error) {
	return self.q.User.Find()
}

func (self *userRepo) Register(user *model.User) (*model.User, error) {
	return user, self.q.User.Clauses(clause.OnConflict{
		OnConstraint: "unique_User",
		UpdateAll:    true,
	}).Create(user)
}

func (self *userRepo) FindByEmail(email string) (*model.User, error) {
	return self.q.User.Where(
		self.q.User.Email.Eq(email),
	).First()
}

func (self *userRepo) UpdateByID(id string, user *model.User) (*model.User, error) {
	_, err := self.q.User.Where(
		self.q.User.ID.Eq(user.ID)).
		Updates(&user)
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return nil, gorm.ErrForeignKeyViolated
	}
	return user, err
}
