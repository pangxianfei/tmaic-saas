package common

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"
	UserAppModel "tmaic/app/UserApp/model"
)

func InitMiddleware(token UserAppModel.Token) (string, error) {
	authSignKey := []byte(config.GetString("auth.sign_key"))

	var tokenTime time.Duration
	tokenTime = 60
	signer := jwt.NewSigner(jwt.HS256, authSignKey, tokenTime*time.Minute)
	claims := UserAppModel.Token{
		UserId:   token.UserId,
		TenantId: token.TenantId,
		Mobile:   token.Mobile,
		Iss:      "tmaic",
	}

	tokenSTR, err := signer.Sign(claims)

	if err != nil {
		return "", err
	}

	return string(tokenSTR), nil

}
