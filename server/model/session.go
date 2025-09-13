package model

import "time"

type Session struct {
	Id uint64 `json:"id" gorm:"column:id;type:bigint(20) unsigned;not null;autoIncrement;primaryKey;comment:主键"`
	UserId uint64 `json:"user_id" gorm:"column:user_id;type:bigint(20) unsigned;not null;default:0;index:idx_user_id;comment:用户id"`
	Name string `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';size:100;comment:名称"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime;<-:create;comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp ON UPDATE CURRENT_TIMESTAMP;not null;default:CURRENT_TIMESTAMP;autoUpdateTime;comment:更新时间"`
}

func (Session) TableName() string {
	return "session"
}