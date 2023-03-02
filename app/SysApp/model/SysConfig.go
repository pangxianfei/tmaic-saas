package model

import "tmaic/app/model"

// SysConfig 系统配置
type SysConfig struct {
	model.BaseModel
	Key         string `gorm:"type:varchar(255);not null;size:128;unique" json:"key" form:"key"` // 配置key
	Value       string `gorm:"type:text" json:"value" form:"value"`                              // 配置值
	Name        string `gorm:"type:varchar(255);not null;" json:"name" form:"name"`              // 配置名称
	Description string `gorm:"type:varchar(255)" json:"description" form:"description"`          // 配置描述
}
