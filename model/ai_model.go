package model

import "time"

type AiModel struct {
	Id uint64 `gorm:"column:id;type:bigint(20) unsigned;not null;autoIncrement;precision:20;scale:0;primaryKey;comment:主键"`
	Name string `gorm:"column:name;type:varchar(100);not null;default:'';size:100;comment:名称"`
	Enabled int8 `gorm:"column:enabled;type:tinyint(2);not null;default:0;precision:2;scale:0;comment:是否启用"`
	MaxOutputToken uint64 `gorm:"column:max_output_token;type:bigint(20) unsigned;not null;default:0;precision:20;scale:0;comment:最大输出token"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime;<-:create;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp ON UPDATE CURRENT_TIMESTAMP;not null;default:CURRENT_TIMESTAMP;autoUpdateTime;comment:更新时间"`
}

func (AiModel) TableName() string {
	return "ai_model"
}