package middleware

import (
	"sync"
	OrderAppMiddleware "tmaic/OrderApp/http/middleware"
	SysAppiddleware "tmaic/SysApp/http/middleware"
	UserAppMiddleware "tmaic/UserApp/http/middleware"

	"gitee.com/pangxianfei/saas/middleware"
)

var once sync.Once

// 所有中间件都要添加在这里注册，注册后才能使用
func init() {
	once.Do(func() {
		//系统中间件
		initCORS()
		middleware.GetPermissions()
		middleware.GetPermissions()
		//用户定义
		SysAppiddleware.GetAdminPermissions()
		UserAppMiddleware.UserAppInfo()
		SysAppiddleware.SysAppInfo()
		OrderAppMiddleware.OrderAppInfo()
	})
}
