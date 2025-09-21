package sessionv2

import (
	"context"

	"github.com/hankeyyh/chat-box-svr/model"
	"github.com/hankeyyh/chat-box-svr/model/table"
)


func Create(ctx context.Context, session *table.SessionV2) error {
	return model.SessionV2.WithContext(ctx).Create(session)
}

func GetByID(ctx context.Context, id uint64) (*table.SessionV2, error) {
	return model.SessionV2.WithContext(ctx).Where(model.SessionV2.ID.Eq(id)).First()
}