// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"go.uber.org/fx/generated/dao/model"
)

func newAd(db *gorm.DB, opts ...gen.DOOption) ad {
	_ad := ad{}

	_ad.adDo.UseDB(db, opts...)
	_ad.adDo.UseModel(&model.Ad{})

	tableName := _ad.adDo.TableName()
	_ad.ALL = field.NewAsterisk(tableName)
	_ad.ID = field.NewString(tableName, "id")
	_ad.Title = field.NewString(tableName, "title")
	_ad.Description = field.NewString(tableName, "description")
	_ad.ShelterID = field.NewString(tableName, "shelter_id")
	_ad.DateCreated = field.NewTime(tableName, "date_created")
	_ad.Images = field.NewField(tableName, "images")
	_ad.CreatedAt = field.NewTime(tableName, "created_at")
	_ad.UpdatedAt = field.NewTime(tableName, "updated_at")

	_ad.fillFieldMap()

	return _ad
}

type ad struct {
	adDo

	ALL         field.Asterisk
	ID          field.String
	Title       field.String
	Description field.String
	ShelterID   field.String
	DateCreated field.Time
	Images      field.Field
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (a ad) Table(newTableName string) *ad {
	a.adDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a ad) As(alias string) *ad {
	a.adDo.DO = *(a.adDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *ad) updateTableName(table string) *ad {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.Title = field.NewString(table, "title")
	a.Description = field.NewString(table, "description")
	a.ShelterID = field.NewString(table, "shelter_id")
	a.DateCreated = field.NewTime(table, "date_created")
	a.Images = field.NewField(table, "images")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *ad) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *ad) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["title"] = a.Title
	a.fieldMap["description"] = a.Description
	a.fieldMap["shelter_id"] = a.ShelterID
	a.fieldMap["date_created"] = a.DateCreated
	a.fieldMap["images"] = a.Images
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a ad) clone(db *gorm.DB) ad {
	a.adDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a ad) replaceDB(db *gorm.DB) ad {
	a.adDo.ReplaceDB(db)
	return a
}

type adDo struct{ gen.DO }

type IAdDo interface {
	gen.SubQuery
	Debug() IAdDo
	WithContext(ctx context.Context) IAdDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdDo
	WriteDB() IAdDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdDo
	Not(conds ...gen.Condition) IAdDo
	Or(conds ...gen.Condition) IAdDo
	Select(conds ...field.Expr) IAdDo
	Where(conds ...gen.Condition) IAdDo
	Order(conds ...field.Expr) IAdDo
	Distinct(cols ...field.Expr) IAdDo
	Omit(cols ...field.Expr) IAdDo
	Join(table schema.Tabler, on ...field.Expr) IAdDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdDo
	Group(cols ...field.Expr) IAdDo
	Having(conds ...gen.Condition) IAdDo
	Limit(limit int) IAdDo
	Offset(offset int) IAdDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdDo
	Unscoped() IAdDo
	Create(values ...*model.Ad) error
	CreateInBatches(values []*model.Ad, batchSize int) error
	Save(values ...*model.Ad) error
	First() (*model.Ad, error)
	Take() (*model.Ad, error)
	Last() (*model.Ad, error)
	Find() ([]*model.Ad, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Ad, err error)
	FindInBatches(result *[]*model.Ad, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Ad) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdDo
	Assign(attrs ...field.AssignExpr) IAdDo
	Joins(fields ...field.RelationField) IAdDo
	Preload(fields ...field.RelationField) IAdDo
	FirstOrInit() (*model.Ad, error)
	FirstOrCreate() (*model.Ad, error)
	FindByPage(offset int, limit int) (result []*model.Ad, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adDo) Debug() IAdDo {
	return a.withDO(a.DO.Debug())
}

func (a adDo) WithContext(ctx context.Context) IAdDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adDo) ReadDB() IAdDo {
	return a.Clauses(dbresolver.Read)
}

func (a adDo) WriteDB() IAdDo {
	return a.Clauses(dbresolver.Write)
}

func (a adDo) Session(config *gorm.Session) IAdDo {
	return a.withDO(a.DO.Session(config))
}

func (a adDo) Clauses(conds ...clause.Expression) IAdDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adDo) Returning(value interface{}, columns ...string) IAdDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adDo) Not(conds ...gen.Condition) IAdDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adDo) Or(conds ...gen.Condition) IAdDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adDo) Select(conds ...field.Expr) IAdDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adDo) Where(conds ...gen.Condition) IAdDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adDo) Order(conds ...field.Expr) IAdDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adDo) Distinct(cols ...field.Expr) IAdDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adDo) Omit(cols ...field.Expr) IAdDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adDo) Join(table schema.Tabler, on ...field.Expr) IAdDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adDo) Group(cols ...field.Expr) IAdDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adDo) Having(conds ...gen.Condition) IAdDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adDo) Limit(limit int) IAdDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adDo) Offset(offset int) IAdDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adDo) Unscoped() IAdDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adDo) Create(values ...*model.Ad) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adDo) CreateInBatches(values []*model.Ad, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adDo) Save(values ...*model.Ad) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adDo) First() (*model.Ad, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ad), nil
	}
}

func (a adDo) Take() (*model.Ad, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ad), nil
	}
}

func (a adDo) Last() (*model.Ad, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ad), nil
	}
}

func (a adDo) Find() ([]*model.Ad, error) {
	result, err := a.DO.Find()
	return result.([]*model.Ad), err
}

func (a adDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Ad, err error) {
	buf := make([]*model.Ad, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adDo) FindInBatches(result *[]*model.Ad, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adDo) Attrs(attrs ...field.AssignExpr) IAdDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adDo) Assign(attrs ...field.AssignExpr) IAdDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adDo) Joins(fields ...field.RelationField) IAdDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adDo) Preload(fields ...field.RelationField) IAdDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adDo) FirstOrInit() (*model.Ad, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ad), nil
	}
}

func (a adDo) FirstOrCreate() (*model.Ad, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ad), nil
	}
}

func (a adDo) FindByPage(offset int, limit int) (result []*model.Ad, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adDo) Delete(models ...*model.Ad) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adDo) withDO(do gen.Dao) *adDo {
	a.DO = *do.(*gen.DO)
	return a
}
