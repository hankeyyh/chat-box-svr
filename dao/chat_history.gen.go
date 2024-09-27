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

func newChatHistory(db *gorm.DB, opts ...gen.DOOption) chatHistory {
	_chatHistory := chatHistory{}

	_chatHistory.chatHistoryDo.UseDB(db, opts...)
	_chatHistory.chatHistoryDo.UseModel(&model.ChatHistory{})

	tableName := _chatHistory.chatHistoryDo.TableName()
	_chatHistory.ALL = field.NewAsterisk(tableName)
	_chatHistory.Id = field.NewUint64(tableName, "id")
	_chatHistory.ParentId = field.NewUint64(tableName, "parent_id")
	_chatHistory.UserId = field.NewUint64(tableName, "user_id")
	_chatHistory.AppId = field.NewUint64(tableName, "app_id")
	_chatHistory.Sender = field.NewString(tableName, "sender")
	_chatHistory.ErrNo = field.NewInt32(tableName, "err_no")
	_chatHistory.Content = field.NewString(tableName, "content")
	_chatHistory.CreatedAt = field.NewTime(tableName, "created_at")

	_chatHistory.fillFieldMap()

	return _chatHistory
}

type chatHistory struct {
	chatHistoryDo

	ALL       field.Asterisk
	Id        field.Uint64
	ParentId  field.Uint64
	UserId    field.Uint64
	AppId     field.Uint64
	Sender    field.String
	ErrNo     field.Int32
	Content   field.String
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (c chatHistory) Table(newTableName string) *chatHistory {
	c.chatHistoryDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chatHistory) As(alias string) *chatHistory {
	c.chatHistoryDo.DO = *(c.chatHistoryDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chatHistory) updateTableName(table string) *chatHistory {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewUint64(table, "id")
	c.ParentId = field.NewUint64(table, "parent_id")
	c.UserId = field.NewUint64(table, "user_id")
	c.AppId = field.NewUint64(table, "app_id")
	c.Sender = field.NewString(table, "sender")
	c.ErrNo = field.NewInt32(table, "err_no")
	c.Content = field.NewString(table, "content")
	c.CreatedAt = field.NewTime(table, "created_at")

	c.fillFieldMap()

	return c
}

func (c *chatHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chatHistory) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 8)
	c.fieldMap["id"] = c.Id
	c.fieldMap["parent_id"] = c.ParentId
	c.fieldMap["user_id"] = c.UserId
	c.fieldMap["app_id"] = c.AppId
	c.fieldMap["sender"] = c.Sender
	c.fieldMap["err_no"] = c.ErrNo
	c.fieldMap["content"] = c.Content
	c.fieldMap["created_at"] = c.CreatedAt
}

func (c chatHistory) clone(db *gorm.DB) chatHistory {
	c.chatHistoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chatHistory) replaceDB(db *gorm.DB) chatHistory {
	c.chatHistoryDo.ReplaceDB(db)
	return c
}

type chatHistoryDo struct{ gen.DO }

type IChatHistoryDo interface {
	gen.SubQuery
	Debug() IChatHistoryDo
	WithContext(ctx context.Context) IChatHistoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IChatHistoryDo
	WriteDB() IChatHistoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IChatHistoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IChatHistoryDo
	Not(conds ...gen.Condition) IChatHistoryDo
	Or(conds ...gen.Condition) IChatHistoryDo
	Select(conds ...field.Expr) IChatHistoryDo
	Where(conds ...gen.Condition) IChatHistoryDo
	Order(conds ...field.Expr) IChatHistoryDo
	Distinct(cols ...field.Expr) IChatHistoryDo
	Omit(cols ...field.Expr) IChatHistoryDo
	Join(table schema.Tabler, on ...field.Expr) IChatHistoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IChatHistoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IChatHistoryDo
	Group(cols ...field.Expr) IChatHistoryDo
	Having(conds ...gen.Condition) IChatHistoryDo
	Limit(limit int) IChatHistoryDo
	Offset(offset int) IChatHistoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IChatHistoryDo
	Unscoped() IChatHistoryDo
	Create(values ...*model.ChatHistory) error
	CreateInBatches(values []*model.ChatHistory, batchSize int) error
	Save(values ...*model.ChatHistory) error
	First() (*model.ChatHistory, error)
	Take() (*model.ChatHistory, error)
	Last() (*model.ChatHistory, error)
	Find() ([]*model.ChatHistory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChatHistory, err error)
	FindInBatches(result *[]*model.ChatHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ChatHistory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IChatHistoryDo
	Assign(attrs ...field.AssignExpr) IChatHistoryDo
	Joins(fields ...field.RelationField) IChatHistoryDo
	Preload(fields ...field.RelationField) IChatHistoryDo
	FirstOrInit() (*model.ChatHistory, error)
	FirstOrCreate() (*model.ChatHistory, error)
	FindByPage(offset int, limit int) (result []*model.ChatHistory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IChatHistoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id uint64) (result model.ChatHistory, err error)
	GetByParentID(parentId uint64) (result []model.ChatHistory, err error)
	BatchGetRecentByUserID(appId uint64, userId uint64, lastId uint64, offset int, limit int) (result []model.ChatHistory, err error)
}

// SELECT * FROM @@table WHERE id = @id LIMIT 1
func (c chatHistoryDo) GetByID(id uint64) (result model.ChatHistory, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM chat_history WHERE id = ? LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE parent_id = @parentId
func (c chatHistoryDo) GetByParentID(parentId uint64) (result []model.ChatHistory, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, parentId)
	generateSQL.WriteString("SELECT * FROM chat_history WHERE parent_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE
// {{if lastId != 0}}
//
//	id < @lastId AND
//
// {{end}}
// app_id = @appId AND user_id = @userId ORDER BY id DESC LIMIT @offset, @limit
func (c chatHistoryDo) BatchGetRecentByUserID(appId uint64, userId uint64, lastId uint64, offset int, limit int) (result []model.ChatHistory, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM chat_history WHERE ")
	if lastId != 0 {
		params = append(params, lastId)
		generateSQL.WriteString("id < ? AND ")
	}
	params = append(params, appId)
	params = append(params, userId)
	params = append(params, offset)
	params = append(params, limit)
	generateSQL.WriteString("app_id = ? AND user_id = ? ORDER BY id DESC LIMIT ?, ? ")

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c chatHistoryDo) Debug() IChatHistoryDo {
	return c.withDO(c.DO.Debug())
}

func (c chatHistoryDo) WithContext(ctx context.Context) IChatHistoryDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chatHistoryDo) ReadDB() IChatHistoryDo {
	return c.Clauses(dbresolver.Read)
}

func (c chatHistoryDo) WriteDB() IChatHistoryDo {
	return c.Clauses(dbresolver.Write)
}

func (c chatHistoryDo) Session(config *gorm.Session) IChatHistoryDo {
	return c.withDO(c.DO.Session(config))
}

func (c chatHistoryDo) Clauses(conds ...clause.Expression) IChatHistoryDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chatHistoryDo) Returning(value interface{}, columns ...string) IChatHistoryDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chatHistoryDo) Not(conds ...gen.Condition) IChatHistoryDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chatHistoryDo) Or(conds ...gen.Condition) IChatHistoryDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chatHistoryDo) Select(conds ...field.Expr) IChatHistoryDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chatHistoryDo) Where(conds ...gen.Condition) IChatHistoryDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chatHistoryDo) Order(conds ...field.Expr) IChatHistoryDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chatHistoryDo) Distinct(cols ...field.Expr) IChatHistoryDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chatHistoryDo) Omit(cols ...field.Expr) IChatHistoryDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chatHistoryDo) Join(table schema.Tabler, on ...field.Expr) IChatHistoryDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chatHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IChatHistoryDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chatHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IChatHistoryDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chatHistoryDo) Group(cols ...field.Expr) IChatHistoryDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chatHistoryDo) Having(conds ...gen.Condition) IChatHistoryDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chatHistoryDo) Limit(limit int) IChatHistoryDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chatHistoryDo) Offset(offset int) IChatHistoryDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chatHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IChatHistoryDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chatHistoryDo) Unscoped() IChatHistoryDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chatHistoryDo) Create(values ...*model.ChatHistory) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chatHistoryDo) CreateInBatches(values []*model.ChatHistory, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chatHistoryDo) Save(values ...*model.ChatHistory) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chatHistoryDo) First() (*model.ChatHistory, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatHistory), nil
	}
}

func (c chatHistoryDo) Take() (*model.ChatHistory, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatHistory), nil
	}
}

func (c chatHistoryDo) Last() (*model.ChatHistory, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatHistory), nil
	}
}

func (c chatHistoryDo) Find() ([]*model.ChatHistory, error) {
	result, err := c.DO.Find()
	return result.([]*model.ChatHistory), err
}

func (c chatHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChatHistory, err error) {
	buf := make([]*model.ChatHistory, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chatHistoryDo) FindInBatches(result *[]*model.ChatHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chatHistoryDo) Attrs(attrs ...field.AssignExpr) IChatHistoryDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chatHistoryDo) Assign(attrs ...field.AssignExpr) IChatHistoryDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chatHistoryDo) Joins(fields ...field.RelationField) IChatHistoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chatHistoryDo) Preload(fields ...field.RelationField) IChatHistoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chatHistoryDo) FirstOrInit() (*model.ChatHistory, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatHistory), nil
	}
}

func (c chatHistoryDo) FirstOrCreate() (*model.ChatHistory, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatHistory), nil
	}
}

func (c chatHistoryDo) FindByPage(offset int, limit int) (result []*model.ChatHistory, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c chatHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chatHistoryDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chatHistoryDo) Delete(models ...*model.ChatHistory) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chatHistoryDo) withDO(do gen.Dao) *chatHistoryDo {
	c.DO = *do.(*gen.DO)
	return c
}
