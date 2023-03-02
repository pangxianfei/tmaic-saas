package SysAppModel

// SysConfig 系统配置
type SysConfig struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	CreateTime  int64  `gorm:"not null" json:"createTime" form:"createTime"`                     // 创建时间
	UpdateTime  int64  `gorm:"not null" json:"updateTime" form:"updateTime"`                     // 更新时间
	Key         string `gorm:"type:varchar(255);not null;size:128;unique" json:"key" form:"key"` // 配置key
	Value       string `gorm:"type:text" json:"value" form:"value"`                              // 配置值
	Name        string `gorm:"type:varchar(255);not null;" json:"name" form:"name"`              // 配置名称
	Description string `gorm:"type:varchar(255)" json:"description" form:"description"`          // 配置描述
}
