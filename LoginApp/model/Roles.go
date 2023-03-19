package UserAppModel

import (
	"database/sql"
	"gitee.com/pangxianfei/framework/model"
)

// Roles 角色表
type Roles struct {
	model.BaseModel
	Id        uint64       `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string       `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	TenantsId int64        `gorm:"column:tenants_id;type:int(11) unsigned;NOT NULL" json:"tenants_id"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	CreatedAt sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
}

// TableName 指定表
func (ros *Roles) TableName() string {
	return ros.SetTableName("sys_roles")
}
