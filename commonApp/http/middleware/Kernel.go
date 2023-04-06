package middleware

import (
	"sync"

	"gitee.com/pangxianfei/saas/middleware"

	FmsAppMiddleware "tmaic/FmsApp/http/middleware"
	NbomAppMiddleware "tmaic/NbomApp/http/middleware"
	OaAppMiddleware "tmaic/OaApp/http/middleware"
	OrderAppMiddleware "tmaic/OrderApp/http/middleware"
	ProductMiddleware "tmaic/ProductApp/http/middleware"
	SrmAppMiddleware "tmaic/SrmApp/http/middleware"
	SupplierAppMiddleware "tmaic/SupplierApp/http/middleware"
	SysAppiddleware "tmaic/SysApp/http/middleware"
	UserAppMiddleware "tmaic/UserApp/http/middleware"
	WmAppMiddleware "tmaic/WmApp/http/middleware"
)

var once sync.Once

// 所有中间件都要添加在这里注册，注册后才能使用
func init() {
	once.Do(func() {
		//系统中间件
		initCORS()
		middleware.GetPermissions()
		middleware.GetPermissions()
		SysAppiddleware.GetAdminPermissions()
		//用户定义
		UserAppMiddleware.UserAppInfo()
		SysAppiddleware.SysAppInfo()
		OrderAppMiddleware.OrderAppInfo()
		ProductMiddleware.ProductAppInfo()
		SrmAppMiddleware.SrmAppInfo()
		FmsAppMiddleware.FmsAppInfo()
		NbomAppMiddleware.NbomAppInfo()
		WmAppMiddleware.WmAppInfo()
		OaAppMiddleware.OaAppInfo()
		SupplierAppMiddleware.SupplierAppInfo()
	})
}
