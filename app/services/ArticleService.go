package services

import (
	"tmaic/app/model"
	"tmaic/vendors/framework/simple"
)

var ArticleService = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct{}

func (s *articleService) Get(id int64) *model.Article {
	return nil
}

func (s *articleService) Take(where ...interface{}) *model.Article {
	return nil
}

func (s *articleService) Find(cnd *simple.SqlCnd) []model.Article {
	return nil
}

func (s *articleService) FindOne(cnd *simple.SqlCnd) *model.Article {
	return nil
}

func (s *articleService) FindPageByParams(params *simple.QueryParams) (list []model.Article, paging *simple.Paging) {
	return
}

func (s *articleService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Article, paging *simple.Paging) {
	return
}

func (s *articleService) Update(t *model.Article) error {
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
