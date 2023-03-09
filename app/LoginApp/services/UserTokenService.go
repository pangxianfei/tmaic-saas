package services

import (
	"errors"
	"gitee.com/pangxianfei/framework/helpers/tmaic"
	"gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/saas"
	"gitee.com/pangxianfei/simple"
	"gitee.com/pangxianfei/simple/date"
	"github.com/kataras/iris/v12"
	"time"
	"tmaic/app/UserApp/buffer"
	UserAppModel "tmaic/app/UserApp/model"
	"tmaic/app/UserApp/repositories"
	"tmaic/app/common"
	"tmaic/app/common/constants"
)

var UserTokenService = newUserTokenService()

func newUserTokenService() *userTokenService {
	return &userTokenService{}
}

type userTokenService struct {
}

// GetUserId 获取当前登录用户的id
func (s *userTokenService) GetUserId(ctx iris.Context) int64 {
	user := s.GetUserInfo(ctx)
	if user != nil {
		return user.Id
	}
	return 0
}

// GetUserInfo 获取当前登录用户
func (s *userTokenService) GetUserInfo(ctx iris.Context) (user *UserAppModel.Admin) {
	token := s.GetUserToken(ctx)

	userToken := buffer.UserTokenCache.Get(token)

	if userToken == nil {
		return nil
	}
	user = UserService.Get(ctx, userToken.UserId)

	// 没找到授权
	if userToken == nil || userToken.Status == constants.StatusDeleted {
		return nil
	}

	// 授权过期
	if userToken.ExpiredAt >= date.NowTimestamp() {
		return nil
	}
	//用户被禁用
	if user == nil || user.Status != constants.StatusOk {
		return nil
	}

	return user
}

// CheckLogin 检查登录状态
func (s *userTokenService) CheckLogin(ctx iris.Context) (*UserAppModel.Admin, *simple.CodeError) {
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
	//退出登陆清除token缓存
	buffer.UserTokenCache.Invalidate(token)
	return repositories.UserTokenRepository.UpdateColumn(simple.DB(), userToken.Id, "status", constants.StatusDeleted)
}

// GetUserToken 从请求体中获取UserToken
func (s *userTokenService) GetUserToken(ctx iris.Context) (userToken string) {
	userToken = ctx.GetHeader("Authorization")
	if len(userToken) > 0 {
		userToken := userToken[7:]
		return userToken
	}
	return userToken
}

// Create 存入DB
func (s *userTokenService) Create(ctx iris.Context, Admin *UserAppModel.Admin, token string) (*UserAppModel.UserToken, error) {
	var iat int64 = time.Now().Unix()
	var exp int64 = config.GetInt64("cache.token_time")
	//保存至DB
	userToken := &UserAppModel.UserToken{
		Token:      token,
		UserId:     Admin.Id,
		TenantId:   Admin.TenantId,
		Mobile:     Admin.Mobile,
		ExpiredAt:  iat + exp,
		Status:     0,
		CreateTime: iat,
		Md5Token:   tmaic.MD5(token),
	}
	err := repositories.UserTokenRepository.Create(saas.DB.Initiation(ctx), userToken)
	if err != nil {
		return nil, errors.New("token创建失败")
	}
	buffer.UserTokenCache.Invalidate(token)
	return userToken, nil
}

// CreateToken 生成token 串
func (s *userTokenService) CreateToken(Admin UserAppModel.Admin) (tokenString string, err error) {
	tokenString, err = common.GetJWTInstantiation(Admin)
	return tokenString, err
}

// Disable 禁用
func (s *userTokenService) Disable(token string) error {
	t := repositories.UserTokenRepository.GetByToken(simple.DB(), token)
	if t == nil {
		return nil
	}
	err := repositories.UserTokenRepository.UpdateColumn(simple.DB(), t.Id, "status", constants.StatusDeleted)
	if err != nil {
		buffer.UserTokenCache.Invalidate(token)
	}
	return err
}
