package services

import (
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/simple"
	"gitee.com/pangxianfei/simple/sqlcmd"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"

	"tmaic/LoginApp/models/pang"
	"tmaic/LoginApp/repositories"
)

var PangService = new(pangService)

type pangService struct{}

func GetPaasDb(ctx iris.Context) *gorm.DB {
	return paas.DB.Initiation(ctx)
}

func (pang *pangService) Get(ctx iris.Context, id int64) *pang.Pang {
	return repositories.PangRepository.Get(GetPaasDb(ctx),id)
}

func (pang *pangService) Take(ctx iris.Context, where ...interface{}) *pang.Pang {
	return nil
}

func (pang *pangService) Find(ctx iris.Context, cnd *sqlcmd.Cnd) []pang.Pang {
	return nil
}

func (pang *pangService) FindOne(ctx iris.Context, cnd *sqlcmd.Cnd) *pang.Pang {
	return nil
}

func (pang *pangService) FindPageByParams(ctx iris.Context, params *simple.QueryParams) (list []pang.Pang, paging *sqlcmd.Paging) {
	return
}

func (pang *pangService) FindPageByCnd(ctx iris.Context, cnd *sqlcmd.Cnd) (list []pang.Pang, paging *sqlcmd.Paging) {
	return
}

func (pang *pangService) Update(ctx iris.Context, t *pang.Pang) error {
	return nil
}

func (pang *pangService) Updates(ctx iris.Context, id int64, columns map[string]interface{}) error {
	return nil
}

func (pang *pangService) UpdateColumn(ctx iris.Context, id int64, name string, value interface{}) error {
	return nil
}

func (pang *pangService) Delete(ctx iris.Context, id int64) error {
	return nil
}
