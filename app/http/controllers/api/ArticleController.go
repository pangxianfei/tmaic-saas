package api

import (
	"github.com/kataras/iris/v12"
	"tmaic/vendors/framework/simple"
)

type ArticleController struct {
	Ctx iris.Context
}

// GetBy 文章详情
func (c *ArticleController) GetBy(articleId int64) *simple.JsonResult {

	return simple.JsonSuccess()
}

// PostCreate 发表文章
func (c *ArticleController) PostCreate() *simple.JsonResult {
	return simple.JsonSuccess()
}

// GetEditBy 编辑时获取详情
func (c *ArticleController) GetEditBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostEditBy 编辑文章
func (c *ArticleController) PostEditBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostDeleteBy 删除文章
func (c *ArticleController) PostDeleteBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// 收藏文章
func (c *ArticleController) PostFavoriteBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// 文章跳转链接
func (c *ArticleController) GetRedirectBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// 最近文章
func (c *ArticleController) GetRecent() *simple.JsonResult {
	return simple.JsonSuccess()
}

// 用户最近的文章
func (c *ArticleController) GetUserRecent() *simple.JsonResult {
	return simple.JsonSuccess()
}

// 用户文章列表
func (c *ArticleController) GetUserArticles() *simple.JsonResult {
	return simple.JsonSuccess()
}

// 文章列表
func (c *ArticleController) GetArticles() *simple.JsonResult {
	return simple.JsonSuccess()
}

// 标签文章列表
func (c *ArticleController) GetTagArticles() *simple.JsonResult {
	return simple.JsonSuccess()
}

// 用户最新的文章
func (c *ArticleController) GetUserNewestBy(userId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// 近期文章
func (c *ArticleController) GetNearlyBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// 相关文章
func (c *ArticleController) GetRelatedBy(articleId int64) *simple.JsonResult {
	return simple.JsonSuccess()
}

// 推荐
func (c *ArticleController) GetRecommend() *simple.JsonResult {
	return simple.JsonSuccess()
}

// 最新文章
func (c *ArticleController) GetNewest() *simple.JsonResult {
	return simple.JsonSuccess()
}

// 热门文章
func (c *ArticleController) GetHot() *simple.JsonResult {
	return simple.JsonSuccess()
}
