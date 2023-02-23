package middleware

import (
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/mitchellh/mapstructure"
	"tmaic/app/model"
)

// Login 获取JWT中的用户数据，封装成实体存放在ctx中供请求调用
var Login context.Handler

func initUserInfo() {
	Login = func(ctx iris.Context) {
		jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
		UserInfo := jwtInfo["UserInfo"]
		var singleUser model.User
		//将 map 转换为指定的结构体
		if err := mapstructure.Decode(UserInfo, &singleUser); err != nil {
			fmt.Println(err)
		}
		ctx.Values().Set("UserInfo", singleUser)
		ctx.Next()
	}
}
