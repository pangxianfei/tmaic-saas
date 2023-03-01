package common

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/iris-contrib/middleware/jwt"
	"time"
	"tmaic/app/model"
)

func GetJWTInstantiation(user model.User) (tokenString string, err error) {
	var exp int64 = config.GetInt64("cache.token_time")
	var iat int64 = time.Now().Unix()
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 根据需求，可以存一些必要的数据
		"UserName": user.UserName,
		"Mobile":   user.Mobile,
		"UserInfo": user,
		"UserId":   user.Id,
		// 签发人
		"iss": "tmaic",
		// 签发时间
		"iat": iat,
		// 设定过期时间
		"exp": iat + exp,
	})

	// 使用设置的秘钥，签名生成jwt字符串
	tokenString, err = token.SignedString([]byte(config.GetString("auth.sign_key")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
