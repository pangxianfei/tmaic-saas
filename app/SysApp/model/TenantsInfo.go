package SysAppModel

import (
	"gitee.com/pangxianfei/framework/helpers/zone"
)

type TenantsInfo struct {
	ID         int64     `gorm:"column:id;primary_key;auto_increment"`
	Mobile     string    `gorm:"size:11;unique;" json:"mobile" form:"mobile"` // 手机
	TenantName string    `gorm:"column:tenant_name;type:varchar(255);not null"`
	Status     int       `gorm:"column:status;type:int(1);not null;default:0;"`
	CreatedAt  zone.Time `gorm:"column:created_at"`
	UpdatedAt  zone.Time `gorm:"column:updated_at"`
}
