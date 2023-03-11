package UserAppModel

import (
	"gitee.com/pangxianfei/framework/helpers/zone"
	"gitee.com/pangxianfei/framework/model"
)

type TenantsInfo struct {
	model.BaseModel
	ID         int64     `gorm:"column:id;primary_key;auto_increment"`
	Mobile     string    `gorm:"size:11;unique;" json:"mobile" form:"mobile"` // 手机
	TenantName string    `gorm:"column:tenant_name;type:varchar(255);not null"`
	Status     int       `gorm:"column:status;type:int(1);not null;default:0;"`
	CreatedAt  zone.Time `gorm:"column:created_at"`
	UpdatedAt  zone.Time `gorm:"column:updated_at"`
}

// TableName 指定表
func (tenantsInfo *TenantsInfo) TableName() string {
	return tenantsInfo.SetTableName("sys_tenants_info")
}
