package buffer

import (
	"encoding/json"

	"gitee.com/pangxianfei/framework/kernel/cache"
	"gitee.com/pangxianfei/simple"
	"tmaic/Loginapp/models/pang"
)

var PangCache = new(newPangCache)

type newPangCache struct {
}

// Get 获取缓存数据
func (c *newPangCache) Get(getKey string) (returnData *pang.Pang) {

	if len(getKey) == 0 || getKey == "" {
		return nil
	}
	getNewKey := simple.MD5(getKey)

	//读取缓存
	cacheData := cache.GetString(getNewKey)

	if len(cacheData) <= 0 {
		return nil
	}

	if err := simple.InterfaceToStruct(cacheData, &returnData); err != nil {
		return nil
	}
	return returnData
}

func (c *newPangCache) SetCacheData(setKey string, cacheData *pang.Pang) {
	newData, _ := json.Marshal(cacheData)

	setNewKey := simple.MD5(setKey)

	cache.AddTokenCache(setNewKey, newData)

}
