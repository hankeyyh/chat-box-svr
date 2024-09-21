package model

import "time"

type App struct {
	Id uint64 `gorm:"column:id;type:bigint(20) unsigned;not null;autoIncrement;primaryKey;comment:主键"`
	ModelId uint64 `gorm:"column:model_id;type:bigint(20) unsigned;not null;default:0;comment:模型id"`
	Name string `gorm:"column:name;type:varchar(100);not null;default:'';size:100;comment:应用名称"`
	CreatedBy string `gorm:"column:created_by;type:varchar(200);not null;default:'';size:200;comment:创建者"`
	Introduction string `gorm:"column:introduction;type:varchar(500);not null;default:'';size:500;comment:介绍"`
	IsPublic int8 `gorm:"column:is_public;type:tinyint(2);not null;default:0;precision:2;scale:0;comment:是否公开"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime;<-:create;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp ON UPDATE CURRENT_TIMESTAMP;not null;default:CURRENT_TIMESTAMP;autoUpdateTime;comment:更新时间"`
}

func (App) TableName() string {
	return "app"
}