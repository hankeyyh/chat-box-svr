package model

import "time"

type App struct {
	Id uint64 `json:"id" gorm:"column:id;type:bigint(20) unsigned;not null;autoIncrement;primaryKey;comment:主键"`
	ModelId uint64 `json:"model_id" gorm:"column:model_id;type:bigint(20) unsigned;not null;default:0;comment:模型id"`
	Name string `json:"name" gorm:"column:name;type:varchar(100);not null;default:'';size:100;comment:应用名称"`
	Temperature float32 `json:"temperature" gorm:"column:temperature;type:float;not null;default:0;precision:3;scale:2;comment:temperature"` 
	TopP float32 `json:"top_p" gorm:"column:top_p;type:float;not null;default:0;precision:3;scale:2;comment:top_p"`
	MaxOutputTokens int `json:"max_output_tokens" gorm:"column:max_output_tokens;type:int(11);not null;default:0;comment:最大输出token"`
	CreatedBy uint64 `json:"created_by" gorm:"column:created_by;type:bigint(20) unsigned;not null;default:0;comment:创建者uid,0:未定义,1:system"`
	Introduction string `json:"introduction" gorm:"column:introduction;type:varchar(500);not null;default:'';size:500;comment:介绍"`
	Prologue string `json:"prologue" gorm:"column:prologue;type:text;not null;comment:开场白"`
	Prompt string `json:"prompt" gorm:"column:prompt;type:text;not null;comment:prompt"`
	IsPublic int8 `json:"is_public" gorm:"column:is_public;type:tinyint(2);not null;default:0;precision:2;scale:0;comment:是否公开"`
	ShowPrompt int8 `json:"show_prompt" gorm:"column:show_prompt;type:tinyint(2);not null;default:0;comment:是否展示prompt"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime;<-:create;comment:创建时间"`
	UpdatedAt time.Time `json:"-" gorm:"column:updated_at;type:timestamp ON UPDATE CURRENT_TIMESTAMP;not null;default:CURRENT_TIMESTAMP;autoUpdateTime;comment:更新时间"`
}

func (App) TableName() string {
	return "app"
}