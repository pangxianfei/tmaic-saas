package model

type BaseModel struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}
