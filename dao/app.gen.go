// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/hankeyyh/chat-box-svr/model"
)

func newApp(db *gorm.DB, opts ...gen.DOOption) app {
	_app := app{}

	_app.appDo.UseDB(db, opts...)
	_app.appDo.UseModel(&model.App{})

	tableName := _app.appDo.TableName()
	_app.ALL = field.NewAsterisk(tableName)
	_app.Id = field.NewUint64(tableName, "id")
	_app.ModelId = field.NewUint64(tableName, "model_id")
	_app.Name = field.NewString(tableName, "name")
	_app.Temperature = field.NewFloat32(tableName, "temperature")
	_app.TopP = field.NewFloat32(tableName, "top_p")
	_app.MaxOutputTokens = field.NewInt(tableName, "max_output_tokens")
	_app.Context = field.NewInt(tableName, "context")
	_app.CreatedBy = field.NewUint64(tableName, "created_by")
	_app.Introduction = field.NewString(tableName, "introduction")
	_app.Prologue = field.NewString(tableName, "prologue")
	_app.Prompt = field.NewString(tableName, "prompt")
	_app.IsPublic = field.NewInt8(tableName, "is_public")
	_app.CreatedAt = field.NewTime(tableName, "created_at")
	_app.UpdatedAt = field.NewTime(tableName, "updated_at")

	_app.fillFieldMap()

	return _app
}

type app struct {
	appDo

	ALL             field.Asterisk
	Id              field.Uint64
	ModelId         field.Uint64
	Name            field.String
	Temperature     field.Float32
	TopP            field.Float32
	MaxOutputTokens field.Int
	Context         field.Int
	CreatedBy       field.Uint64
	Introduction    field.String
	Prologue        field.String
	Prompt          field.String
	IsPublic        field.Int8
	CreatedAt       field.Time
	UpdatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (a app) Table(newTableName string) *app {
	a.appDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a app) As(alias string) *app {
	a.appDo.DO = *(a.appDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *app) updateTableName(table string) *app {
	a.ALL = field.NewAsterisk(table)
	a.Id = field.NewUint64(table, "id")
	a.ModelId = field.NewUint64(table, "model_id")
	a.Name = field.NewString(table, "name")
	a.Temperature = field.NewFloat32(table, "temperature")
	a.TopP = field.NewFloat32(table, "top_p")
	a.MaxOutputTokens = field.NewInt(table, "max_output_tokens")
	a.Context = field.NewInt(table, "context")
	a.CreatedBy = field.NewUint64(table, "created_by")
	a.Introduction = field.NewString(table, "introduction")
	a.Prologue = field.NewString(table, "prologue")
	a.Prompt = field.NewString(table, "prompt")
	a.IsPublic = field.NewInt8(table, "is_public")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *app) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *app) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 14)
	a.fieldMap["id"] = a.Id
	a.fieldMap["model_id"] = a.ModelId
	a.fieldMap["name"] = a.Name
	a.fieldMap["temperature"] = a.Temperature
	a.fieldMap["top_p"] = a.TopP
	a.fieldMap["max_output_tokens"] = a.MaxOutputTokens
	a.fieldMap["context"] = a.Context
	a.fieldMap["created_by"] = a.CreatedBy
	a.fieldMap["introduction"] = a.Introduction
	a.fieldMap["prologue"] = a.Prologue
	a.fieldMap["prompt"] = a.Prompt
	a.fieldMap["is_public"] = a.IsPublic
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a app) clone(db *gorm.DB) app {
	a.appDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a app) replaceDB(db *gorm.DB) app {
	a.appDo.ReplaceDB(db)
	return a
}

type appDo struct{ gen.DO }

type IAppDo interface {
	gen.SubQuery
	Debug() IAppDo
	WithContext(ctx context.Context) IAppDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAppDo
	WriteDB() IAppDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAppDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAppDo
	Not(conds ...gen.Condition) IAppDo
	Or(conds ...gen.Condition) IAppDo
	Select(conds ...field.Expr) IAppDo
	Where(conds ...gen.Condition) IAppDo
	Order(conds ...field.Expr) IAppDo
	Distinct(cols ...field.Expr) IAppDo
	Omit(cols ...field.Expr) IAppDo
	Join(table schema.Tabler, on ...field.Expr) IAppDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAppDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAppDo
	Group(cols ...field.Expr) IAppDo
	Having(conds ...gen.Condition) IAppDo
	Limit(limit int) IAppDo
	Offset(offset int) IAppDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAppDo
	Unscoped() IAppDo
	Create(values ...*model.App) error
	CreateInBatches(values []*model.App, batchSize int) error
	Save(values ...*model.App) error
	First() (*model.App, error)
	Take() (*model.App, error)
	Last() (*model.App, error)
	Find() ([]*model.App, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.App, err error)
	FindInBatches(result *[]*model.App, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.App) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAppDo
	Assign(attrs ...field.AssignExpr) IAppDo
	Joins(fields ...field.RelationField) IAppDo
	Preload(fields ...field.RelationField) IAppDo
	FirstOrInit() (*model.App, error)
	FirstOrCreate() (*model.App, error)
	FindByPage(offset int, limit int) (result []*model.App, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAppDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByName(name string) (result []model.App, err error)
	GetByModelID(modelId uint64) (result []model.App, err error)
	GetByID(id uint64) (result []model.App, err error)
	GetByAuthorAndId(createdBy uint64, id uint64) (result model.App, err error)
	AllPublic() (result []model.App, err error)
	AllPrivateByAuthor(createdBy uint64) (result []model.App, err error)
	UpdateIsPublic(id uint64, isPublic bool) (err error)
}

// SELECT * FROM @@table WHERE name = @name
func (a appDo) GetByName(name string) (result []model.App, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM app WHERE name = ? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE model_id = @modelId
func (a appDo) GetByModelID(modelId uint64) (result []model.App, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, modelId)
	generateSQL.WriteString("SELECT * FROM app WHERE model_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE id = @id LIMIT 1
func (a appDo) GetByID(id uint64) (result []model.App, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM app WHERE id = ? LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE created_by = @createdBy and id = @id LIMIT 1
func (a appDo) GetByAuthorAndId(createdBy uint64, id uint64) (result model.App, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, createdBy)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM app WHERE created_by = ? and id = ? LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE is_public = 1
func (a appDo) AllPublic() (result []model.App, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM app WHERE is_public = 1 ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE is_public = 0 AND created_by = @createdBy
func (a appDo) AllPrivateByAuthor(createdBy uint64) (result []model.App, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, createdBy)
	generateSQL.WriteString("SELECT * FROM app WHERE is_public = 0 AND created_by = ? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET is_public = @isPublic WHERE id = @id
func (a appDo) UpdateIsPublic(id uint64, isPublic bool) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, isPublic)
	params = append(params, id)
	generateSQL.WriteString("UPDATE app SET is_public = ? WHERE id = ? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (a appDo) Debug() IAppDo {
	return a.withDO(a.DO.Debug())
}

func (a appDo) WithContext(ctx context.Context) IAppDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appDo) ReadDB() IAppDo {
	return a.Clauses(dbresolver.Read)
}

func (a appDo) WriteDB() IAppDo {
	return a.Clauses(dbresolver.Write)
}

func (a appDo) Session(config *gorm.Session) IAppDo {
	return a.withDO(a.DO.Session(config))
}

func (a appDo) Clauses(conds ...clause.Expression) IAppDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appDo) Returning(value interface{}, columns ...string) IAppDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appDo) Not(conds ...gen.Condition) IAppDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appDo) Or(conds ...gen.Condition) IAppDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appDo) Select(conds ...field.Expr) IAppDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appDo) Where(conds ...gen.Condition) IAppDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appDo) Order(conds ...field.Expr) IAppDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appDo) Distinct(cols ...field.Expr) IAppDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appDo) Omit(cols ...field.Expr) IAppDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appDo) Join(table schema.Tabler, on ...field.Expr) IAppDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAppDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appDo) RightJoin(table schema.Tabler, on ...field.Expr) IAppDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appDo) Group(cols ...field.Expr) IAppDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appDo) Having(conds ...gen.Condition) IAppDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appDo) Limit(limit int) IAppDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appDo) Offset(offset int) IAppDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAppDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appDo) Unscoped() IAppDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appDo) Create(values ...*model.App) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appDo) CreateInBatches(values []*model.App, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appDo) Save(values ...*model.App) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appDo) First() (*model.App, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) Take() (*model.App, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) Last() (*model.App, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) Find() ([]*model.App, error) {
	result, err := a.DO.Find()
	return result.([]*model.App), err
}

func (a appDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.App, err error) {
	buf := make([]*model.App, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appDo) FindInBatches(result *[]*model.App, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appDo) Attrs(attrs ...field.AssignExpr) IAppDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appDo) Assign(attrs ...field.AssignExpr) IAppDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appDo) Joins(fields ...field.RelationField) IAppDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appDo) Preload(fields ...field.RelationField) IAppDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appDo) FirstOrInit() (*model.App, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) FirstOrCreate() (*model.App, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) FindByPage(offset int, limit int) (result []*model.App, count int64, err error) {
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

func (a appDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appDo) Delete(models ...*model.App) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appDo) withDO(do gen.Dao) *appDo {
	a.DO = *do.(*gen.DO)
	return a
}
