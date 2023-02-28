package buffer

import (
	"tmaic/app/model"
)

type userCache struct {
}

var UserCache = newUserCache()

func newUserCache() *userCache {
	return &userCache{}
}

func (c *userCache) Get(userId int64) *model.User {

	return nil

}

func (c *userCache) Invalidate(userId int64) {

}

func (c *userCache) GetScoreRank() []model.User {

	return nil

}
