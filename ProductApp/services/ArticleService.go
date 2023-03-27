package services

import (
	"gitee.com/pangxianfei/simple"
	"gitee.com/pangxianfei/simple/sqlcmd"

	"tmaic/OrderApp/models"
)

var ArticleService = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct{}

func (s *articleService) Get(id int64) *models.Article {
	return nil
}

func (s *articleService) Take(where ...interface{}) *models.Article {
	return nil
}

func (s *articleService) Find(cnd *sqlcmd.Cnd) []models.Article {
	return nil
}

func (s *articleService) FindOne(cnd *sqlcmd.Cnd) *models.Article {
	return nil
}

func (s *articleService) FindPageByParams(params *simple.QueryParams) (list []models.Article, paging *sqlcmd.Paging) {
	return
}

func (s *articleService) FindPageByCnd(cnd *sqlcmd.Cnd) (list []models.Article, paging *sqlcmd.Paging) {
	return
}

func (s *articleService) Update(t *models.Article) error {
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
