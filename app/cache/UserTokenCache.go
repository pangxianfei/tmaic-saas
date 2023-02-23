package cache

import (
	"encoding/json"

	helpers "tmaic/vendors/framework/helpers/cache"
	tokenCache "tmaic/vendors/framework/helpers/tmaic"

	///helpers "tmaic/vendors/framework/helpers/cache

	"tmaic/app/model"
)

var UserTokenCache = newUserTokenCache()

type userTokenCache struct {
}

func newUserTokenCache() *userTokenCache {
	return &userTokenCache{}
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

	if len(newData) <= 0 {
		return nil
	}

	err := json.Unmarshal([]byte(newData), &UserToken)
	//debug.Dump(UserToken)
	//debug.Dump(err)

	if err == nil {
		return UserToken
	}

	return nil

}

func (c *userTokenCache) Invalidate(token string) {
	//c.cache.Invalidate(token)
}
