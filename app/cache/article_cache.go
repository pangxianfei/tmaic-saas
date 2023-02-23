package cache

import (
	"github.com/goburrow/cache"

	"tmaic/app/model"
)

var (
	articleRecommendCacheKey = "recommend_articles_cache"
	articleHotCacheKey       = "hot_articles_cache"
)

var ArticleCache = newArticleCache()

type articleCache struct {
	recommendCache cache.LoadingCache
	hotCache       cache.LoadingCache
}

func newArticleCache() *articleCache {
	return nil
}

func (c *articleCache) GetRecommendArticles() []model.Article {
	val, err := c.recommendCache.Get(articleRecommendCacheKey)
	if err != nil {
		return nil
	}
	if val != nil {
		return val.([]model.Article)
	}
	return nil
}

func (c *articleCache) InvalidateRecommend() {
	c.recommendCache.Invalidate(articleRecommendCacheKey)
}

func (c *articleCache) GetHotArticles() []model.Article {
	val, err := c.hotCache.Get(articleHotCacheKey)
	if err != nil {
		return nil
	}
	if val != nil {
		return val.([]model.Article)
	}
	return nil
}
