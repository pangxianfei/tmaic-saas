package cache

import (
	"encoding/json"
	"errors"
	"gitee.com/pangxianfei/simple"
	"github.com/goburrow/cache"
	"time"
	helpers "tmaic/vendors/framework/helpers/cache"
	tokenCache "tmaic/vendors/framework/helpers/tmaic"

	///helpers "tmaic/vendors/framework/helpers/cache

	"tmaic/app/model"
	"tmaic/app/repositories"
)

var UserTokenCache = newUserTokenCache()

type userTokenCache struct {
	cache cache.LoadingCache
}

func newUserTokenCache() *userTokenCache {
	return &userTokenCache{
		cache: cache.NewLoadingCache(
			func(key cache.Key) (value cache.Value, e error) {
				value = repositories.UserTokenRepository.GetByToken(simple.DB(), key.(string))
				if value == nil {
					e = errors.New("数据不存在")
				}
				return
			},
			cache.WithMaximumSize(1000),
			cache.WithExpireAfterAccess(60*time.Minute),
		),
	}
}

func (c *userTokenCache) Get(token string) *model.UserToken {
	if len(token) == 0 {
		return nil
	}
	tokenKey := tokenCache.MD5(token)
	cacheData := helpers.GetString(tokenKey)

	//val, err := c.cache.Get(token)
	if len(cacheData) <= 0 {
		return nil
	}

	var UserToken *model.UserToken

	newData := tokenCache.InterfaceToString(cacheData)

	err := json.Unmarshal([]byte(newData), &UserToken)
	//debug.Dump(UserToken)
	//debug.Dump(err)

	if err == nil {
		return UserToken
	}

	return nil

}

func (c *userTokenCache) Invalidate(token string) {
	c.cache.Invalidate(token)
}
