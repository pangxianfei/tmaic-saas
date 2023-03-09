package services

import (
	"errors"
	"gitee.com/pangxianfei/saas"
	"gitee.com/pangxianfei/saas/sysmodel"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	"strings"
	UserCache "tmaic/app/UserApp/buffer"
	"tmaic/app/UserApp/http/requests"
	UserAppModel "tmaic/app/UserApp/model"
	"tmaic/app/UserApp/repositories"
	"tmaic/app/common"
	"tmaic/app/common/constants"
	"tmaic/app/common/validate"
)

// 邮箱验证邮件有效期（小时）
//const emailVerifyExpireHour = 24

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (s *userService) Get(ctx iris.Context, id int64) *UserAppModel.Admin {
	return repositories.UserRepository.Get(saas.DB.Initiation(ctx), id)
}

func (s *userService) Take(where ...interface{}) *UserAppModel.Admin {
	return repositories.UserRepository.Take(simple.DB(), where...)
}

func (s *userService) Find(cnd *simple.SqlCnd) []UserAppModel.Admin {
	return repositories.UserRepository.Find(simple.DB(), cnd)
}

func (s *userService) FindOne(cnd *simple.SqlCnd) *UserAppModel.Admin {
	return repositories.UserRepository.FindOne(simple.DB(), cnd)
}

func (s *userService) FindPageByParams(params *simple.QueryParams) (list []UserAppModel.Admin, paging *simple.Paging) {
	return repositories.UserRepository.FindPageByParams(simple.DB(), params)
}

func (s *userService) FindPageByCnd(cnd *simple.SqlCnd) (list []UserAppModel.Admin, paging *simple.Paging) {
	return repositories.UserRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *userService) Create(t *UserAppModel.Admin) error {
	err := repositories.UserRepository.Create(simple.DB(), t)
	if err == nil {
		UserCache.UserCache.Invalidate(t.Id)
	}
	return nil
}

func (s *userService) Update(t *UserAppModel.Admin) error {
	err := repositories.UserRepository.Update(simple.DB(), t)
	UserCache.UserCache.Invalidate(t.Id)
	return err
}

func (s *userService) Updates(id int64, columns map[string]interface{}) error {
	err := repositories.UserRepository.Updates(simple.DB(), id, columns)
	UserCache.UserCache.Invalidate(id)
	return err
}

func (s *userService) UpdateColumn(id int64, name string, value interface{}) error {
	err := repositories.UserRepository.UpdateColumn(simple.DB(), id, name, value)
	UserCache.UserCache.Invalidate(id)
	return err
}

func (s *userService) Delete(id int64) {
	repositories.UserRepository.Delete(simple.DB(), id)
	UserCache.UserCache.Invalidate(id)
}

// Scan 扫描
func (s *userService) Scan(callback func(users []UserAppModel.Admin)) {
	var cursor int64
	for {
		list := repositories.UserRepository.Find(simple.DB(), simple.NewSqlCnd().Where("id > ?", cursor).Asc("id").Limit(100))
		if list == nil || len(list) == 0 {
			break
		}
		cursor = list[len(list)-1].Id
		callback(list)
	}
}

// GetByEmail 根据邮箱查找
func (s *userService) GetByEmail(email string) *UserAppModel.Admin {
	return repositories.UserRepository.GetByEmail(simple.DB(), email)
}

// GetByUsername 根据用户名查找
func (s *userService) GetByUsername(username string) *UserAppModel.Admin {
	return repositories.UserRepository.GetByUsername(simple.DB(), username)
}

// GetByMobile 根据用户名查找
func (s *userService) GetByMobile(ctx iris.Context, mobile string) *UserAppModel.Admin {
	return repositories.UserRepository.GetByMobile(saas.DB.Initiation(ctx), mobile)
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

// SetUsername 设置用户名
func (s *userService) SetUsername(ctx iris.Context, userId int64, username string) error {
	username = strings.TrimSpace(username)
	if err := validate.IsUsername(username); err != nil {
		return err
	}

	user := s.Get(ctx, userId)
	if len(user.UserName) > 0 {
		return errors.New("你已设置了用户名，无法重复设置。")
	}
	if s.isUsernameExists(username) {
		return errors.New("用户名：" + username + " 已被占用")
	}
	return s.UpdateColumn(userId, "username", username)
}

// SetEmail 设置密码
func (s *userService) SetEmail(userId int64, email string) error {
	email = strings.TrimSpace(email)
	if err := validate.IsEmail(email); err != nil {
		return err
	}
	if s.isEmailExists(email) {
		return errors.New("邮箱：" + email + " 已被占用")
	}
	return s.UpdateColumn(userId, "email", email)
}

// SetPassword 设置密码
func (s *userService) SetPassword(ctx iris.Context, userId int64, password, rePassword string) error {
	if err := validate.IsPassword(password, rePassword); err != nil {
		return err
	}
	user := s.Get(ctx, userId)
	if len(user.Password) > 0 {
		return errors.New("你已设置了密码，如需修改请前往修改页面。")
	}
	password = simple.EncodePassword(password)
	return s.UpdateColumn(userId, "password", password)
}

// UpdatePassword 修改密码
func (s *userService) UpdatePassword(ctx iris.Context, userId int64, oldPassword, password, rePassword string) error {
	if err := validate.IsPassword(password, rePassword); err != nil {
		return err
	}
	user := s.Get(ctx, userId)

	if len(user.Password) == 0 {
		return errors.New("你没设置密码，请先设置密码")
	}

	if !simple.ValidatePassword(user.Password, oldPassword) {
		return errors.New("旧密码验证失败")
	}

	return s.UpdateColumn(userId, "password", simple.EncodePassword(password))
}

// VerifyEmail 验证邮箱
func (s *userService) VerifyEmail(userId int64, token string) error {
	return nil
}

// SignIn 登录
func (s *userService) SignIn(ctx iris.Context, UserLogin requests.UserLogin) (*UserAppModel.Admin, *UserAppModel.UserToken, error) {

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
		return nil, nil, errors.New("租户应用不存在")
	}

	if Result.RowsAffected > 0 {
		for _, appDb := range InstanceDB {
			db := saas.DB.SetTenantsDb(appDb.TenantsId, appDb.AppId)
			if db != nil && appDb.AppName == constants.UserDb {
				UserInstanceDB = appDb
				ctx.Values().Set("TenantId", appDb.TenantsId)
				ctx.Values().Set("AppId", appDb.AppId)
			}
		}
	}
	//检查上下文
	if ctx.Values().Get("TenantId").(int64) <= 0 {
		return nil, nil, errors.New("租户用户不存在")
	}
	if ctx.Values().Get("AppId").(int64) <= 0 {
		return nil, nil, errors.New("租户应用不存在")
	}

	Admin := new(UserAppModel.Admin)

	if validate.IsNumber(UserLogin.Mobile) {
		Admin = s.GetByMobile(ctx, UserLogin.Mobile)
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

	var UserToken *UserAppModel.UserToken
	UserToken, err = UserTokenService.Create(ctx, Admin, tokenSTR)

	if err != nil {
		return nil, nil, err
	}
	//缓存token信息
	UserCache.UserTokenCache.SetCacheUserToken(tokenSTR, UserToken)
	return Admin, UserToken, nil

}
