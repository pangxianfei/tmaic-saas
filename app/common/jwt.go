package common

import (
	"github.com/iris-contrib/middleware/jwt"
	"time"
	"tmaic/app/model"
	"tmaic/vendors/framework/config"
)

func GetJWTInstantiation(user model.User) (tokenString string, err error) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 根据需求，可以存一些必要的数据
		"UserName": user.UserName,
		"UserInfo": user,
		"UserId":   user.Id,
		// 签发人
		"iss": "tmaic",
		// 签发时间
		"iat": time.Now().Unix(),
		// 设定过期时间
		"exp": config.GetInt64("cache.token_time"),
	})
	// 使用设置的秘钥，签名生成jwt字符串
	tokenString, err = token.SignedString([]byte(config.GetString("auth.sign_key")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
