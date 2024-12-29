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

	"PawInHand/generated/dao/model"
)

func newVolunteerwork(db *gorm.DB, opts ...gen.DOOption) volunteerwork {
	_volunteerwork := volunteerwork{}

	_volunteerwork.volunteerworkDo.UseDB(db, opts...)
	_volunteerwork.volunteerworkDo.UseModel(&model.Volunteerwork{})

	tableName := _volunteerwork.volunteerworkDo.TableName()
	_volunteerwork.ALL = field.NewAsterisk(tableName)
	_volunteerwork.ID = field.NewString(tableName, "id")
	_volunteerwork.Title = field.NewString(tableName, "title")
	_volunteerwork.Description = field.NewString(tableName, "description")
	_volunteerwork.Date = field.NewTime(tableName, "date")
	_volunteerwork.Venue = field.NewString(tableName, "venue")
	_volunteerwork.City = field.NewString(tableName, "city")
	_volunteerwork.State = field.NewString(tableName, "state")
	_volunteerwork.CreatedAt = field.NewTime(tableName, "created_at")
	_volunteerwork.UpdatedAt = field.NewTime(tableName, "updated_at")

	_volunteerwork.fillFieldMap()

	return _volunteerwork
}

type volunteerwork struct {
	volunteerworkDo

	ALL         field.Asterisk
	ID          field.String
	Title       field.String
	Description field.String
	Date        field.Time
	Venue       field.String
	City        field.String
	State       field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (v volunteerwork) Table(newTableName string) *volunteerwork {
	v.volunteerworkDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v volunteerwork) As(alias string) *volunteerwork {
	v.volunteerworkDo.DO = *(v.volunteerworkDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *volunteerwork) updateTableName(table string) *volunteerwork {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewString(table, "id")
	v.Title = field.NewString(table, "title")
	v.Description = field.NewString(table, "description")
	v.Date = field.NewTime(table, "date")
	v.Venue = field.NewString(table, "venue")
	v.City = field.NewString(table, "city")
	v.State = field.NewString(table, "state")
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")

	v.fillFieldMap()

	return v
}

func (v *volunteerwork) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *volunteerwork) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 9)
	v.fieldMap["id"] = v.ID
	v.fieldMap["title"] = v.Title
	v.fieldMap["description"] = v.Description
	v.fieldMap["date"] = v.Date
	v.fieldMap["venue"] = v.Venue
	v.fieldMap["city"] = v.City
	v.fieldMap["state"] = v.State
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
}

func (v volunteerwork) clone(db *gorm.DB) volunteerwork {
	v.volunteerworkDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v volunteerwork) replaceDB(db *gorm.DB) volunteerwork {
	v.volunteerworkDo.ReplaceDB(db)
	return v
}

type volunteerworkDo struct{ gen.DO }

type IVolunteerworkDo interface {
	gen.SubQuery
	Debug() IVolunteerworkDo
	WithContext(ctx context.Context) IVolunteerworkDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVolunteerworkDo
	WriteDB() IVolunteerworkDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVolunteerworkDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVolunteerworkDo
	Not(conds ...gen.Condition) IVolunteerworkDo
	Or(conds ...gen.Condition) IVolunteerworkDo
	Select(conds ...field.Expr) IVolunteerworkDo
	Where(conds ...gen.Condition) IVolunteerworkDo
	Order(conds ...field.Expr) IVolunteerworkDo
	Distinct(cols ...field.Expr) IVolunteerworkDo
	Omit(cols ...field.Expr) IVolunteerworkDo
	Join(table schema.Tabler, on ...field.Expr) IVolunteerworkDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVolunteerworkDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVolunteerworkDo
	Group(cols ...field.Expr) IVolunteerworkDo
	Having(conds ...gen.Condition) IVolunteerworkDo
	Limit(limit int) IVolunteerworkDo
	Offset(offset int) IVolunteerworkDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVolunteerworkDo
	Unscoped() IVolunteerworkDo
	Create(values ...*model.Volunteerwork) error
	CreateInBatches(values []*model.Volunteerwork, batchSize int) error
	Save(values ...*model.Volunteerwork) error
	First() (*model.Volunteerwork, error)
	Take() (*model.Volunteerwork, error)
	Last() (*model.Volunteerwork, error)
	Find() ([]*model.Volunteerwork, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Volunteerwork, err error)
	FindInBatches(result *[]*model.Volunteerwork, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Volunteerwork) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVolunteerworkDo
	Assign(attrs ...field.AssignExpr) IVolunteerworkDo
	Joins(fields ...field.RelationField) IVolunteerworkDo
	Preload(fields ...field.RelationField) IVolunteerworkDo
	FirstOrInit() (*model.Volunteerwork, error)
	FirstOrCreate() (*model.Volunteerwork, error)
	FindByPage(offset int, limit int) (result []*model.Volunteerwork, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVolunteerworkDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v volunteerworkDo) Debug() IVolunteerworkDo {
	return v.withDO(v.DO.Debug())
}

func (v volunteerworkDo) WithContext(ctx context.Context) IVolunteerworkDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v volunteerworkDo) ReadDB() IVolunteerworkDo {
	return v.Clauses(dbresolver.Read)
}

func (v volunteerworkDo) WriteDB() IVolunteerworkDo {
	return v.Clauses(dbresolver.Write)
}

func (v volunteerworkDo) Session(config *gorm.Session) IVolunteerworkDo {
	return v.withDO(v.DO.Session(config))
}

func (v volunteerworkDo) Clauses(conds ...clause.Expression) IVolunteerworkDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v volunteerworkDo) Returning(value interface{}, columns ...string) IVolunteerworkDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v volunteerworkDo) Not(conds ...gen.Condition) IVolunteerworkDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v volunteerworkDo) Or(conds ...gen.Condition) IVolunteerworkDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v volunteerworkDo) Select(conds ...field.Expr) IVolunteerworkDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v volunteerworkDo) Where(conds ...gen.Condition) IVolunteerworkDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v volunteerworkDo) Order(conds ...field.Expr) IVolunteerworkDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v volunteerworkDo) Distinct(cols ...field.Expr) IVolunteerworkDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v volunteerworkDo) Omit(cols ...field.Expr) IVolunteerworkDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v volunteerworkDo) Join(table schema.Tabler, on ...field.Expr) IVolunteerworkDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v volunteerworkDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVolunteerworkDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v volunteerworkDo) RightJoin(table schema.Tabler, on ...field.Expr) IVolunteerworkDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v volunteerworkDo) Group(cols ...field.Expr) IVolunteerworkDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v volunteerworkDo) Having(conds ...gen.Condition) IVolunteerworkDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v volunteerworkDo) Limit(limit int) IVolunteerworkDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v volunteerworkDo) Offset(offset int) IVolunteerworkDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v volunteerworkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVolunteerworkDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v volunteerworkDo) Unscoped() IVolunteerworkDo {
	return v.withDO(v.DO.Unscoped())
}

func (v volunteerworkDo) Create(values ...*model.Volunteerwork) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v volunteerworkDo) CreateInBatches(values []*model.Volunteerwork, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v volunteerworkDo) Save(values ...*model.Volunteerwork) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v volunteerworkDo) First() (*model.Volunteerwork, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Volunteerwork), nil
	}
}

func (v volunteerworkDo) Take() (*model.Volunteerwork, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Volunteerwork), nil
	}
}

func (v volunteerworkDo) Last() (*model.Volunteerwork, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Volunteerwork), nil
	}
}

func (v volunteerworkDo) Find() ([]*model.Volunteerwork, error) {
	result, err := v.DO.Find()
	return result.([]*model.Volunteerwork), err
}

func (v volunteerworkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Volunteerwork, err error) {
	buf := make([]*model.Volunteerwork, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v volunteerworkDo) FindInBatches(result *[]*model.Volunteerwork, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v volunteerworkDo) Attrs(attrs ...field.AssignExpr) IVolunteerworkDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v volunteerworkDo) Assign(attrs ...field.AssignExpr) IVolunteerworkDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v volunteerworkDo) Joins(fields ...field.RelationField) IVolunteerworkDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v volunteerworkDo) Preload(fields ...field.RelationField) IVolunteerworkDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v volunteerworkDo) FirstOrInit() (*model.Volunteerwork, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Volunteerwork), nil
	}
}

func (v volunteerworkDo) FirstOrCreate() (*model.Volunteerwork, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Volunteerwork), nil
	}
}

func (v volunteerworkDo) FindByPage(offset int, limit int) (result []*model.Volunteerwork, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v volunteerworkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v volunteerworkDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v volunteerworkDo) Delete(models ...*model.Volunteerwork) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *volunteerworkDo) withDO(do gen.Dao) *volunteerworkDo {
	v.DO = *do.(*gen.DO)
	return v
}
