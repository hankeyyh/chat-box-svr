package aimodel

import (
	"context"

	"github.com/hankeyyh/chat-box-svr/model/table"
	"github.com/hankeyyh/chat-box-svr/model"
)

func Create(ctx context.Context, aiModel *table.AiModel) error {
	return model.AiModel.WithContext(ctx).Create(aiModel)
}

func GetByID(ctx context.Context, id uint64) (*table.AiModel, error) {
	return model.AiModel.WithContext(ctx).Where(model.AiModel.ID.Eq(id)).First()
}