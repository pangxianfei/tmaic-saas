package services

import (
	"errors"
	"gitee.com/pangxianfei/saas"
	"gitee.com/pangxianfei/saas/sysmodel"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	"tmaic/app/common"
	"tmaic/app/common/constants"
	"tmaic/app/common/validate"

	"tmaic/app/LoginApp/http/requests"
	"tmaic/app/LoginApp/repositories"

	UserCache "tmaic/app/LoginApp/buffer"
	LoginAppModel "tmaic/app/LoginApp/model"
)

// 邮箱验证邮件有效期（小时）
//const emailVerifyExpireHour = 24

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (s *userService) Get(ctx iris.Context, id int64) *LoginAppModel.Admin {
	return repositories.UserRepository.Get(saas.DB.Initiation(ctx), id)
}

func (s *userService) Take(where ...interface{}) *LoginAppModel.Admin {
	return repositories.UserRepository.Take(simple.DB(), where...)
}

func (s *userService) Find(cnd *simple.SqlCnd) []LoginAppModel.Admin {
	return repositories.UserRepository.Find(simple.DB(), cnd)
}

func (s *userService) Create(t *LoginAppModel.Admin) error {
	err := repositories.UserRepository.Create(simple.DB(), t)
	if err == nil {
		UserCache.UserCache.Invalidate(t.Id)
	}
	return nil
}

// GetByEmail 根据邮箱查找
func (s *userService) GetByEmail(email string) *LoginAppModel.Admin {
	return repositories.UserRepository.GetByEmail(simple.DB(), email)
}

// GetByUsername 根据用户名查找
func (s *userService) GetByUsername(username string) *LoginAppModel.Admin {
	return repositories.UserRepository.GetByUsername(simple.DB(), username)
}

// GetByMobile 根据用户名查找
func (s *userService) GetByMobile(mobile string) *LoginAppModel.Admin {
	return repositories.UserRepository.GetByMobile(simple.DB(), mobile)
}

// isEmailExists 邮箱是否存在
func (s *userService) isEmailExists(email string) bool {
	if len(email) == 0 { // 如果邮箱为空，那么就认为是不存在
		return false
	}
	return s.GetByEmail(email) != nil
}

// isUsernameExists 用户名是否存在
func (s *userService) isUsernameExists(username string) bool {
	return s.GetByUsername(username) != nil
}

// VerifyEmail 验证邮箱
func (s *userService) VerifyEmail(userId int64, token string) error {
	return nil
}

// SignIn 登录
func (s *userService) SignIn(ctx iris.Context, UserLogin requests.UserLogin) (*LoginAppModel.Admin, *LoginAppModel.UserToken, error) {

	if len(UserLogin.Mobile) == 0 {
		return nil, nil, errors.New("手机号/用户名/邮箱不能为空")
	}
	if len(UserLogin.Password) == 0 {
		return nil, nil, errors.New("密码不能为空")
	}
	//出始化租户 DB 连接对象
	tenantDbWhere := &sysmdel.RetrieveDB{
		TenantsId: UserLogin.TenantId,
		Status:    1,
	}
	var InstanceDB []sysmdel.InstanceDb
	var UserInstanceDB sysmdel.InstanceDb
	Result := simple.DB().Debug().Model(&sysmdel.InstanceDb{}).Where(tenantDbWhere).Find(&InstanceDB)

	if Result.RowsAffected <= 0 {
		return nil, nil, errors.New("租户应用不存在-1")
	}

	if Result.RowsAffected > 0 {
		for _, appDb := range InstanceDB {
			db := saas.DB.SetTenantsDb(appDb.TenantsId, appDb.AppId)
			if db != nil && appDb.Code == constants.UserDb {
				UserInstanceDB = appDb
				ctx.Values().Set("TenantId", appDb.TenantsId)
				ctx.Values().Set("AppId", appDb.AppId)
			}
		}
	}
	//检查上下文
	if ctx.Values().Get("TenantId").(int64) <= 0 {
		return nil, nil, errors.New("租户用户不存在-2")
	}
	if ctx.Values().Get("AppId").(int64) <= 0 {
		return nil, nil, errors.New("租户应用不存在-3")
	}

	Admin := new(LoginAppModel.Admin)

	if validate.IsNumber(UserLogin.Mobile) {
		Admin = s.GetByMobile(UserLogin.Mobile)
	}

	if Admin == nil || Admin.Status != constants.StatusOk {
		return nil, nil, errors.New("用户不存在或被禁用")
	}

	//检验密码比较耗时 大约90毫秒
	if !simple.ValidatePassword(Admin.Password, UserLogin.Password) {
		return nil, nil, errors.New("密码错误")
	}

	//生成token
	TokenModel := sysmdel.Token{
		TenantId: Admin.TenantId,
		AppId:    UserInstanceDB.AppId,
		Mobile:   Admin.Mobile,
		UserId:   Admin.Id,
	}

	tokenSTR, err := common.InitMiddleware(TokenModel)

	if err != nil {
		return nil, nil, err
	}

	var UserToken *LoginAppModel.UserToken
	UserToken, err = UserTokenService.Create(Admin, tokenSTR)

	if err != nil {
		return nil, nil, err
	}
	//缓存token信息
	UserCache.UserTokenCache.SetCacheUserToken(tokenSTR, UserToken)
	return Admin, UserToken, nil

}
