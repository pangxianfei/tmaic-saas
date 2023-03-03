package services

import (
	"errors"
	"gitee.com/pangxianfei/simple"
	"gitee.com/pangxianfei/simple/date"
	"gorm.io/gorm"
	"strings"
	"tmaic/app/OrderApp/repositories"
	buffer2 "tmaic/app/UserApp/buffer"
	UserAppModel "tmaic/app/UserApp/model"
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

func (s *userService) Get(id int64) *UserAppModel.User {
	return repositories.UserRepository.Get(simple.DB(), id)
}

func (s *userService) Take(where ...interface{}) *UserAppModel.User {
	return repositories.UserRepository.Take(simple.DB(), where...)
}

func (s *userService) Find(cnd *simple.SqlCnd) []UserAppModel.User {
	return repositories.UserRepository.Find(simple.DB(), cnd)
}

func (s *userService) FindOne(cnd *simple.SqlCnd) *UserAppModel.User {
	return repositories.UserRepository.FindOne(simple.DB(), cnd)
}

func (s *userService) FindPageByParams(params *simple.QueryParams) (list []UserAppModel.User, paging *simple.Paging) {
	return repositories.UserRepository.FindPageByParams(simple.DB(), params)
}

func (s *userService) FindPageByCnd(cnd *simple.SqlCnd) (list []UserAppModel.User, paging *simple.Paging) {
	return repositories.UserRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *userService) Create(t *UserAppModel.User) error {
	err := repositories.UserRepository.Create(simple.DB(), t)
	if err == nil {
		buffer2.UserCache.Invalidate(t.Id)
	}
	return nil
}

func (s *userService) Update(t *UserAppModel.User) error {
	err := repositories.UserRepository.Update(simple.DB(), t)
	buffer2.UserCache.Invalidate(t.Id)
	return err
}

func (s *userService) Updates(id int64, columns map[string]interface{}) error {
	err := repositories.UserRepository.Updates(simple.DB(), id, columns)
	buffer2.UserCache.Invalidate(id)
	return err
}

func (s *userService) UpdateColumn(id int64, name string, value interface{}) error {
	err := repositories.UserRepository.UpdateColumn(simple.DB(), id, name, value)
	buffer2.UserCache.Invalidate(id)
	return err
}

func (s *userService) Delete(id int64) {
	repositories.UserRepository.Delete(simple.DB(), id)
	buffer2.UserCache.Invalidate(id)
}

// Scan 扫描
func (s *userService) Scan(callback func(users []UserAppModel.User)) {
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
func (s *userService) GetByEmail(email string) *UserAppModel.User {
	return repositories.UserRepository.GetByEmail(simple.DB(), email)
}

// GetByUsername 根据用户名查找
func (s *userService) GetByUsername(username string) *UserAppModel.User {
	return repositories.UserRepository.GetByUsername(simple.DB(), username)
}

// GetByMobile 根据用户名查找
func (s *userService) GetByMobile(mobile string) *UserAppModel.User {
	return repositories.UserRepository.GetByMobile(simple.DB(), mobile)
}

// SignUp 注册
func (s *userService) SignUp(mobile, password, rePassword string) (*UserAppModel.User, error) {
	mobile = strings.TrimSpace(mobile)
	if len(mobile) == 0 {
		return nil, errors.New("手机号不能为空")
	}
	err := validate.IsPassword(password, rePassword)
	if err != nil {
		return nil, err
	}
	user := &UserAppModel.User{}
	user.Nickname = mobile
	// 验证手机号
	if validate.IsMobile(mobile) {
		if s.GetByMobile(mobile) != nil {
			return nil, errors.New("手机号：" + mobile + " 已被占用")
		}
	}
	user.Mobile = mobile
	user.UserName = mobile
	user.Password = simple.EncodePassword(password)
	user.Status = constants.StatusOk
	user.CreateTime = date.NowTimestamp()
	user.UpdateTime = date.NowTimestamp()

	err = simple.DB().Transaction(func(tx *gorm.DB) error {
		if err := repositories.UserRepository.Create(tx, user); err != nil {
			return err
		}

		avatarUrl := ""
		if err != nil {
			return err
		}

		if err := repositories.UserRepository.UpdateColumn(tx, user.Id, "avatar", avatarUrl); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
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
func (s *userService) SetUsername(userId int64, username string) error {
	username = strings.TrimSpace(username)
	if err := validate.IsUsername(username); err != nil {
		return err
	}

	user := s.Get(userId)
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
func (s *userService) SetPassword(userId int64, password, rePassword string) error {
	if err := validate.IsPassword(password, rePassword); err != nil {
		return err
	}
	user := s.Get(userId)
	if len(user.Password) > 0 {
		return errors.New("你已设置了密码，如需修改请前往修改页面。")
	}
	password = simple.EncodePassword(password)
	return s.UpdateColumn(userId, "password", password)
}

// UpdatePassword 修改密码
func (s *userService) UpdatePassword(userId int64, oldPassword, password, rePassword string) error {
	if err := validate.IsPassword(password, rePassword); err != nil {
		return err
	}
	user := s.Get(userId)

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
func (s *userService) SignIn(username string, password string) (*UserAppModel.User, *UserAppModel.UserToken, error) {

	if len(username) == 0 {
		return nil, nil, errors.New("手机号/用户名/邮箱不能为空")
	}
	if len(password) == 0 {
		return nil, nil, errors.New("密码不能为空")
	}
	user := new(UserAppModel.User)

	if validate.IsNumber(username) {
		user = s.GetByMobile(username)
	} else if err := validate.IsEmail(username); err == nil { // 如果用户输入的是邮箱

		user = s.GetByEmail(username)

	} else if err := validate.IsEmail(username); err == nil { //手机号登陆

		user = s.GetByUsername(username)

	} else { //用户名登陆
		user = s.GetByUsername(username)
	}

	if user == nil || user.Status != constants.StatusOk {
		return nil, nil, errors.New("用户不存在或被禁用")
	}

	//检验密码比较耗时 大约90毫秒
	if !simple.ValidatePassword(user.Password, password) {
		return nil, nil, errors.New("密码错误")
	}

	//生成token
	//token, err := UserTokenService.CreateToken(*user)

	TokenModel := UserAppModel.Token{
		TenantId: user.TenantId,
		Mobile:   user.Mobile,
		UserId:   user.Id,
	}

	tokenSTR, err := common.InitMiddleware(TokenModel)

	//debug.Dd(tokenSTR)

	if err != nil {
		return nil, nil, err
	}

	var UserToken *UserAppModel.UserToken
	UserToken, err = UserTokenService.Create(user, tokenSTR)

	if err != nil {
		return nil, nil, err
	}
	//缓存token信息
	buffer2.UserTokenCache.SetCacheUserToken(tokenSTR, UserToken)
	return user, UserToken, nil

}
