package repositories

import (
	"PawInHand/generated/dao"
	"PawInHand/generated/dao/model"
	"gorm.io/gen"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_auth_repo.go . AuthRepo
type AuthRepo interface {
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	DeleteByID(id string) (gen.ResultInfo, error)
	Login(email, password string) (*model.User, error)
	Logout(userID string) error
}

type authRepo struct {
	q *dao.Query
}

func NewAuthRepo(query *dao.Query) AuthRepo {
	return &authRepo{q: query}
}

func (self *authRepo) FindByEmail(email string) (*model.User, error) {
	return self.q.User.Where(
		self.q.User.Email.Eq(email),
	).First()
}

func (self *authRepo) Create(user *model.User) (*model.User, error) {
	return user, self.q.User.Create(user)
}

func (self *authRepo) DeleteByID(id string) (gen.ResultInfo, error) {
	return self.q.User.Where(self.q.User.ID.Eq(id)).Delete()
}

func (self *authRepo) Login(email, password string) (*model.User, error) {
	return self.q.User.Where(
		self.q.User.Email.Eq(email),
		self.q.User.Password.Eq(password),
	).First()
}

func (self *authRepo) Logout(userID string) error {
	// Handle logout logic, e.g., removing session/token
	return nil
}
