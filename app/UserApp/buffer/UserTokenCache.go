package buffer

import (
	"encoding/json"
	"gitee.com/pangxianfei/framework/helpers/cache"
	"gitee.com/pangxianfei/framework/helpers/tmaic"
	UserAppModel "tmaic/app/UserApp/model"
)

var UserTokenCache = newUserTokenCache()

type userTokenCache struct {
}

func newUserTokenCache() *userTokenCache {
	return &userTokenCache{}
}

// Get 获取缓存数据
func (c *userTokenCache) Get(token string) *UserAppModel.UserToken {

	if len(token) == 0 {
		return nil
	}
	tokenKey := tmaic.MD5(token)

	//读取缓存
	cacheData := cache.GetString(tokenKey)

	if len(cacheData) <= 0 {
		return nil
	}

	var UserToken *UserAppModel.UserToken

	newData := tmaic.InterfaceToString(cacheData)
	if len(newData) <= 0 {
		return nil
	}

	if err := json.Unmarshal([]byte(newData), &UserToken); err == nil {
		return UserToken
	}

	return nil
}

func (c *userTokenCache) SetCacheUserToken(token string, userToken *UserAppModel.UserToken) {
	userTokenData, _ := json.Marshal(userToken)
	tokenKey := tmaic.MD5(token)
	cache.AddTokenCache(tokenKey, userTokenData)
}

func (c *userTokenCache) Invalidate(token string) bool {
	if len(token) <= 0 {
		return false
	}
	tokenKey := tmaic.MD5(token)
	return cache.Forget(tokenKey)
}
