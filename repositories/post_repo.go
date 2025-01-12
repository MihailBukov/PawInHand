package repositories

import (
	"PawInHand/generated/dao"
	"PawInHand/generated/dao/model"
	"gorm.io/gen"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_post_repo.go . PostRepo
type PostRepo interface {
	FindByID(id string) (*model.Post, error)
	Create(post *model.Post) (*model.Post, error)
	UpdateByID(id string, post *model.Post) (*model.Post, error)
	DeleteByID(id string) (gen.ResultInfo, error)
	GetAll() ([]*model.Post, error)
}

type postRepo struct {
	q *dao.Query
}

func NewPostRepo(query *dao.Query) PostRepo {
	return &postRepo{q: query}
}

func (self *postRepo) FindByID(id string) (*model.Post, error) {
	return self.q.Post.Where(
		self.q.Post.ID.Eq(id),
	).First()
}

func (self *postRepo) DeleteByID(id string) (gen.ResultInfo, error) {
	return self.q.Post.Where(self.q.Post.ID.Eq(id)).Delete()
}

func (self *postRepo) GetAll() ([]*model.Post, error) {
	return self.q.Post.Find()
}

func (self *postRepo) Create(post *model.Post) (*model.Post, error) {
	return post, self.q.Post.Create(post)
}

func (self *postRepo) UpdateByID(id string, post *model.Post) (*model.Post, error) {
	_, err := self.q.Post.Where(
		self.q.Post.ID.Eq(post.ID)).
		Updates(&post)
	return post, err
}
