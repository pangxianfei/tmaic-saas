package SysAppModel

import "gitee.com/pangxianfei/framework/helpers/zone"

// AppInfo 应用信息表
type AppInfo struct {
	Id          int64     `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Key         string    `gorm:"type:varchar(20);not null;size:20;unique" json:"key" form:"key"` // 应用key
	Name        string    `gorm:"type:varchar(255);not null;" json:"name" form:"name"`            // 应用名称
	Status      int64     `gorm:"type:int(1);not null;default:0;" json:"status" form:"status"`    // 应用名称
	Description string    `gorm:"type:varchar(255)" json:"description" form:"description"`        // 应用描述
	CreatedAt   zone.Time `gorm:"column:created_at"`
	UpdatedAt   zone.Time `gorm:"column:updated_at"`
}
