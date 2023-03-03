package common

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"
	UserAppModel "tmaic/app/UserApp/model"
)

func InitMiddleware(token UserAppModel.Token) (string, error) {
	authSignKey := []byte(config.GetString("auth.sign_key"))
	signer := jwt.NewSigner(jwt.HS256, authSignKey, 1*time.Minute)

	var exp int64 = config.GetInt64("cache.token_time")
	var iat int64 = time.Now().Unix()

	claims := UserAppModel.Token{
		UserId:   token.UserId,
		TenantId: token.TenantId,
		Mobile:   token.Mobile,
		Exp:      iat + exp,
		Iat:      iat,
		Iss:      "tmaic",
	}

	tokenSTR, err := signer.Sign(claims)

	if err != nil {
		return "", err
	}

	return string(tokenSTR), nil

}
