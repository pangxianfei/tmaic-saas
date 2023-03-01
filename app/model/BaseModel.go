package model

type BaseModel struct {
	Id         int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	CreateTime int64 `gorm:"not null" json:"createTime" form:"createTime"` // 创建时间
	UpdateTime int64 `gorm:"not null" json:"updateTime" form:"updateTime"` // 更新时间
}
