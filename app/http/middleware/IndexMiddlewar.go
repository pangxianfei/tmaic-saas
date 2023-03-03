package middleware

import (
	"sync"
	OrderAppMiddleware "tmaic/app/OrderApp/http/middleware"
	SysAppiddleware "tmaic/app/SysApp/http/middleware"
	UserAppMiddleware "tmaic/app/UserApp/http/middleware"
)

var once sync.Once

// 所有中间件都要添加在这里注册，注册后才能使用
func init() {
	once.Do(func() {
		initCORS()
		TenantInfo()
		UserAppMiddleware.UserAppInfo()
		SysAppiddleware.SysAppInfo()
		OrderAppMiddleware.OrderAppInfo()
	})
}
