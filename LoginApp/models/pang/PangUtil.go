package pang

import (
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/simple/sqlcmd"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func GetDb(ctx iris.Context) *gorm.DB {
	return paas.DB.Initiation(ctx)
}

func Get(ctx iris.Context, id string) (pang Pang) {
	GetDb(ctx).Where("id", id).First(&pang)
	return
}

func GetFirst(ctx iris.Context, field, value string) (pang Pang) {
	GetDb(ctx).Where("? = ?", field, value).First(&pang)
	return
}

func All() (ctx iris.Context, pangs []Pang) {
	GetDb(ctx).Find(&pangs)
	return
}

func IsExist(ctx iris.Context, field, value string) bool {
	var count int64
	GetDb(ctx).Model(Pang{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(ctx iris.Context, perPage int) (pangs []Pang, paging sqlcmd.Paging) {
	return
}
