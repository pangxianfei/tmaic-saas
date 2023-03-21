package services

import (
	"gitee.com/pangxianfei/simple"
	"gitee.com/pangxianfei/simple/sqlcmd"
	"tmaic/OrderApp/model"
)

var ArticleService = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct{}

func (s *articleService) Get(id int64) *OrderAppModel.Article {
	return nil
}

func (s *articleService) Take(where ...interface{}) *OrderAppModel.Article {
	return nil
}

func (s *articleService) Find(cnd *sqlcmd.Cnd) []OrderAppModel.Article {
	return nil
}

func (s *articleService) FindOne(cnd *sqlcmd.Cnd) *OrderAppModel.Article {
	return nil
}

func (s *articleService) FindPageByParams(params *simple.QueryParams) (list []OrderAppModel.Article, paging *sqlcmd.Paging) {
	return
}

func (s *articleService) FindPageByCnd(cnd *sqlcmd.Cnd) (list []OrderAppModel.Article, paging *sqlcmd.Paging) {
	return
}

func (s *articleService) Update(t *OrderAppModel.Article) error {
	return nil
}

func (s *articleService) Updates(id int64, columns map[string]interface{}) error {
	return nil
}

func (s *articleService) UpdateColumn(id int64, name string, value interface{}) error {
	return nil
}

func (s *articleService) Delete(id int64) error {
	return nil
}
