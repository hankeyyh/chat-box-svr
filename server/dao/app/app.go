package app


import (
	"context"

	"github.com/hankeyyh/chat-box-svr/model/table"
	"github.com/hankeyyh/chat-box-svr/model"
)

func Create(ctx context.Context, app *table.App) error {
	return model.App.WithContext(ctx).Create(app)
}

func GetByID(ctx context.Context, id uint64) (*table.App, error) {
	return model.App.WithContext(ctx).Where(model.App.ID.Eq(id)).First()
}