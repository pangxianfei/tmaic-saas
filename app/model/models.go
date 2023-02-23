package model

var Models = []interface{}{
	&User{}, &UserToken{}, &Article{}, &SysConfig{},
}

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}
