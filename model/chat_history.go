package model

import (
	"time"
)

type ChatHistory struct {
	Id uint64 `json:"id" gorm:"column:id;type:bigint(20) unsigned;not null;autoIncrement;primaryKey;comment:主键"`
	ParentId *uint64 `json:"parent_id" gorm:"column:parent_id;type:bigint(20) unsigned;comment:父id"`
	UserId uint64 `json:"user_id" gorm:"column:user_id;type:bigint(20) unsigned;not null;default:0;index:idx_user_id;comment:用户id"`
	SessionId uint64 `json:"session_id" gorm:"column:session_id;type:bigint(20) unsigned;not null;default:0;comment:所属会话id"`
	AppId uint64 `json:"app_id" gorm:"column:app_id;type:bigint(20) unsigned;not null;default:0;index:idx_app_id;comment:应用id"`
	Sender string `json:"sender" gorm:"column:sender;type:varchar(100);not null;default:'';size:100;comment:发送者"`
	ErrNo int32 `json:"err_no" gorm:"column:err_no;type:int(11);not null;default:0;comment:错误码"`
	Content string `json:"content" gorm:"column:content;type:longtext;comment:内容"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime;<-:create;comment:创建时间"`
}

func (c ChatHistory) TableName() string {
	return "chat_history"
}