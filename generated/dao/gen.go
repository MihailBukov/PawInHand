// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q             = new(Query)
	Ad            *ad
	Post          *post
	User          *user
	Volunteerwork *volunteerwork
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Ad = &Q.Ad
	Post = &Q.Post
	User = &Q.User
	Volunteerwork = &Q.Volunteerwork
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:            db,
		Ad:            newAd(db, opts...),
		Post:          newPost(db, opts...),
		User:          newUser(db, opts...),
		Volunteerwork: newVolunteerwork(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Ad            ad
	Post          post
	User          user
	Volunteerwork volunteerwork
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		Ad:            q.Ad.clone(db),
		Post:          q.Post.clone(db),
		User:          q.User.clone(db),
		Volunteerwork: q.Volunteerwork.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		Ad:            q.Ad.replaceDB(db),
		Post:          q.Post.replaceDB(db),
		User:          q.User.replaceDB(db),
		Volunteerwork: q.Volunteerwork.replaceDB(db),
	}
}

type queryCtx struct {
	Ad            IAdDo
	Post          IPostDo
	User          IUserDo
	Volunteerwork IVolunteerworkDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Ad:            q.Ad.WithContext(ctx),
		Post:          q.Post.WithContext(ctx),
		User:          q.User.WithContext(ctx),
		Volunteerwork: q.Volunteerwork.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
