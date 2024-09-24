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
	Q           = new(Query)
	AiModel     *aiModel
	App         *app
	ChatHistory *chatHistory
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AiModel = &Q.AiModel
	App = &Q.App
	ChatHistory = &Q.ChatHistory
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:          db,
		AiModel:     newAiModel(db, opts...),
		App:         newApp(db, opts...),
		ChatHistory: newChatHistory(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AiModel     aiModel
	App         app
	ChatHistory chatHistory
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		AiModel:     q.AiModel.clone(db),
		App:         q.App.clone(db),
		ChatHistory: q.ChatHistory.clone(db),
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
		db:          db,
		AiModel:     q.AiModel.replaceDB(db),
		App:         q.App.replaceDB(db),
		ChatHistory: q.ChatHistory.replaceDB(db),
	}
}

type queryCtx struct {
	AiModel     IAiModelDo
	App         IAppDo
	ChatHistory IChatHistoryDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AiModel:     q.AiModel.WithContext(ctx),
		App:         q.App.WithContext(ctx),
		ChatHistory: q.ChatHistory.WithContext(ctx),
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
