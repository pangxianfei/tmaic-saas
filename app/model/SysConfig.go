package model

// SysConfig 系统配置
type SysConfig struct {
	BaseModel
	Key         string `gorm:"not null;size:128;unique" json:"key" form:"key"` // 配置key
	Value       string `gorm:"type:text" json:"value" form:"value"`            // 配置值
	Name        string `gorm:"not null;size:32" json:"name" form:"name"`       // 配置名称
	Description string `gorm:"size:128" json:"description" form:"description"` // 配置描述
	CreateTime  int64  `gorm:"not null" json:"createTime" form:"createTime"`   // 创建时间
	UpdateTime  int64  `gorm:"not null" json:"updateTime" form:"updateTime"`   // 更新时间
}
