package model

type UserToken struct {
	Id         int64  `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Mobile     string `gorm:"size:11;" json:"mobile" form:"mobile"` // 手机
	Token      string `gorm:"column:token;type:longtext"`
	UserId     int64  `gorm:"column:user_id;not null;index:idx_user_token_user_id;" json:"userId" form:"userId"`
	ExpiredAt  int64  `gorm:"not null" json:"expiredAt" form:"expiredAt"`
	Status     int    `gorm:"not null;index:idx_user_token_status" json:"status" form:"status"`
	CreateTime int64  `gorm:"not null" json:"createTime" form:"createTime"`
	Md5Token   string `gorm:"size:255;unique;" json:"md5_token" form:"md5_token"` // token MD5

}
