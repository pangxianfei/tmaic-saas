// Package pang  LoginApp 模型
package pang

import (
	"github.com/kataras/iris/v12"
)

type Pang struct {
	Id uint64 `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
}

// TableName 指定表
func (pang *Pang) TableName() string {
	return "pangs"
}

func (pang *Pang) Create(ctx iris.Context) {}

func (pang *Pang) Save(ctx iris.Context) (rowsAffected int64) {
	return 0
}

func (pang *Pang) Delete(ctx iris.Context) (rowsAffected int64) {
	return 0
}
