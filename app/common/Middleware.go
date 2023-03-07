package common

import (
	"gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/saas/sysmodel"
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"
)

func InitMiddleware(token sysmdel.Token) (string, error) {
	authSignKey := []byte(config.GetString("auth.sign_key"))

	var tokenTime time.Duration
	tokenTime = 60
	signer := jwt.NewSigner(jwt.HS256, authSignKey, tokenTime*time.Minute)
	claims := sysmdel.Token{
		UserId:   token.UserId,
		TenantId: token.TenantId,
		AppId:    token.AppId,
		Mobile:   token.Mobile,
		Iss:      "tmaic",
	}

	tokenSTR, err := signer.Sign(claims)

	if err != nil {
		return "", err
	}

	return string(tokenSTR), nil

}
