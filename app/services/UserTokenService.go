package services

import (
	"errors"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"time"
	"tmaic/app/cache"
	"tmaic/app/common/constants"
	"tmaic/app/model"
	"tmaic/app/repositories"
	"tmaic/vendors/framework/config"
	"tmaic/vendors/framework/simple"
	"tmaic/vendors/framework/simple/date"
)

var UserTokenService = newUserTokenService()

func newUserTokenService() *userTokenService {
	return &userTokenService{}
}

type userTokenService struct {
}

// GetCurrentUserId 获取当前登录用户的id
func (s *userTokenService) GetCurrentUserId(ctx iris.Context) int64 {
	user := s.GetUserInfo(ctx)
	if user != nil {
		return user.Id
	}
	return 0
}

// GetUserInfo 获取当前登录用户
func (s *userTokenService) GetUserInfo(ctx iris.Context) (user *model.User) {
	token := s.GetUserToken(ctx)

	userToken := cache.UserTokenCache.Get(token)

	if userToken == nil {
		return nil
	}
	user = UserService.Get(userToken.UserId)

	// 没找到授权
	if userToken == nil || userToken.Status == constants.StatusDeleted {
		return nil
	}

	// 授权过期
	ExpiredAt := userToken.ExpiredAt + userToken.CreateTime
	if ExpiredAt >= date.NowTimestamp() {
		return nil
	}
	//用户被禁用
	if user == nil || user.Status != constants.StatusOk {
		return nil
	}

	return user
}

// CheckLogin 检查登录状态
func (s *userTokenService) CheckLogin(ctx iris.Context) (*model.User, *simple.CodeError) {
	user := s.GetUserInfo(ctx)
	if user == nil {
		return nil, simple.ErrorNotLogin
	}
	return user, nil
}

// Signout 退出登录
func (s *userTokenService) Signout(ctx iris.Context) error {
	token := s.GetUserToken(ctx)
	userToken := repositories.UserTokenRepository.GetByToken(simple.DB(), token)
	if userToken == nil {
		return nil
	}
	return repositories.UserTokenRepository.UpdateColumn(simple.DB(), userToken.Id, "status", constants.StatusDeleted)
}

// GetUserToken 从请求体中获取UserToken
func (s *userTokenService) GetUserToken(ctx iris.Context) string {

	userToken := ctx.GetHeader("Authorization")

	if len(userToken) > 0 {
		userToken := userToken[7:]
		return userToken
	}
	userToken = ctx.GetHeader("X-User-Token")
	userToken = userToken[7:]
	return userToken

}
func (s *userTokenService) Create(t *model.UserToken) error {
	err := repositories.UserTokenRepository.Create(simple.DB(), t)
	if err != nil {
		return errors.New("token创建失败")
	}
	cache.UserTokenCache.Invalidate(t.Token)
	return nil
}

// CreateToken get jwt string with expiration time 20 minutes
func (s *userTokenService) CreateToken(user model.User) (tokenString string, err error) {
	//tokenString, err = common.GetJWTInstantiation(user)

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

// Disable 禁用
func (s *userTokenService) Disable(token string) error {
	t := repositories.UserTokenRepository.GetByToken(simple.DB(), token)
	if t == nil {
		return nil
	}
	err := repositories.UserTokenRepository.UpdateColumn(simple.DB(), t.Id, "status", constants.StatusDeleted)
	if err != nil {
		cache.UserTokenCache.Invalidate(token)
	}
	return err
}
