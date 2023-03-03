package api

import (
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type OrderController struct {
	Ctx iris.Context
}

func (c *OrderController) GetBy(articleId int64) *simple.JsonResult {

	return simple.JsonSuccess()
}

func (c *OrderController) PostCreate() *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) GetEditBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) PostEditBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostDeleteBy 删除文章
func (c *OrderController) PostDeleteBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) GetRedirectBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// 最近文章
func (c *OrderController) GetRecent() *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) GetUserRecent() *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) GetUserOrder() *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) GetList() *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) GetUserNewestBy(userId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) GetNearlyBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) GetRelatedBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// 推荐
func (c *OrderController) GetRecommend() *simple.JsonResult {
	return simple.JsonSuccess()
}

// 最新文章
func (c *OrderController) GetNewest() *simple.JsonResult {
	return simple.JsonSuccess()
}

// 热门文章
func (c *OrderController) GetHot() *simple.JsonResult {
	return simple.JsonSuccess()
}
