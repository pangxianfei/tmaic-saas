package buffer

import (
	UserAppModel "tmaic/app/UserApp/model"
)

type userCache struct {
}

var UserCache = newUserCache()

func newUserCache() *userCache {
	return &userCache{}
}

func (c *userCache) Get(userId int64) *UserAppModel.Admin {

	return nil

}

func (c *userCache) Invalidate(userId int64) {

}

func (c *userCache) GetScoreRank() []UserAppModel.Admin {

	return nil

}