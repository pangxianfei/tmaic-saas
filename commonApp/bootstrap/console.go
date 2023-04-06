package bootstrap

import (
	"fmt"
	"strings"

	"gitee.com/pangxianfei/framework/console"
	"gitee.com/pangxianfei/saas/sysmodel"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"

	SysAppModel "tmaic/SysApp/models"
)

func (s *Saas) UserRouteNameList(app *iris.Application, AppName string, Port string, AppId int64) {
	routeList := app.GetRoutes()
	var index int = 1
	//db := simple.DB()
	for _, value := range routeList {
		if strings.Contains(value.MainHandlerName, "tmaic") || strings.Contains(value.MainHandlerName, "iris") {
			continue
		}
		if value.Method == "POST" || value.Method == "GET" {
			/*
				var RouteNameMd5Value string
				RouteNameMd5ValueFile := fmt.Sprintf("%s", value.Name)
				RouteNameMd5Value = simple.MD5(RouteNameMd5ValueFile)
				if value.Name == "" {
					RouteNameMd5Value = ""
				}
				AuthorityMain := sysmdel.Authority{
					Description:       value.Description,
					Pid:               0,
					AppId:             AppId,
					SourceFileName:    value.SourceFileName,
					RouteName:         value.Name,
					RouteNameMd5Value: RouteNameMd5Value,
				}
				db.Model(&sysmdel.Authority{}).FirstOrCreate(&AuthorityMain, sysmdel.Authority{Description: value.Description, Pid: 0, SourceFileName: value.SourceFileName})
				var Authority sysmdel.Authority
				var isAuthority sysmdel.Authority
				db.Where(&sysmdel.Authority{Description: value.Description}).First(&Authority)
				if Authority.Id > 0 {
					Md5ValueFile := fmt.Sprintf("%s_%s", value.MainHandlerName, value.RegisterFileName)
					Md5Value := simple.MD5(Md5ValueFile)
					db.Where("pid > ? and md5_value = ?", 0, simple.MD5(Md5Value)).Find(&isAuthority)

					if isAuthority.Id <= 0 {
						createAuthority := &sysmdel.Authority{
							Pid:               Authority.Id,
							AppId:             AppId,
							Name:              value.MainHandlerName, //simple.Split(value.MainHandlerName, "."),
							Description:       value.Description,
							RegisterFileName:  value.RegisterFileName,
							MainHandlerName:   value.MainHandlerName,
							Method:            value.Method,
							FormattedPath:     value.FormattedPath,
							StaticPath:        value.StaticPath(),
							Path:              value.Path,
							SourceFileName:    value.SourceFileName,
							RouteName:         value.Name,
							RouteNameMd5Value: RouteNameMd5Value,
							Status:            1,
							Md5Value:          simple.MD5(Md5Value),
						}
						db.Create(createAuthority)
					}
				}
			*/

			consoleSTR := ""
			consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-4d", index)
			consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-70s", value.MainHandlerName)
			consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-5s", value.Method)
			//consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-45s", value.Path)
			consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-30s", value.StaticPath())
			//consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-20s", value.RegisterFileName)
			console.Println(console.CODE_SUCCESS, " "+" "+consoleSTR)
			index++
		}
	}
	console.Println(console.CODE_WARNING, " "+console.Sprintf(console.CODE_WARNING, "%s listening on: http://localhost%s", AppName, Port))
}

func (s *Saas) SysRouteNameList(app *iris.Application, AppName string, Port string, AppId int64) {
	routeList := app.GetRoutes()
	var index int = 1
	db := simple.DB()

	for _, value := range routeList {
		if strings.Contains(value.MainHandlerName, "tmaic") || strings.Contains(value.MainHandlerName, "iris") {
			continue
		}
		if value.Method == "POST" || value.Method == "GET" {
			var RouteNameMd5Value string
			RouteNameMd5ValueFile := fmt.Sprintf("%s", value.Name)
			RouteNameMd5Value = simple.MD5(RouteNameMd5ValueFile)
			if value.Name == "" {
				RouteNameMd5Value = ""
			}
			AuthorityMain := SysAppModel.Authority{
				Description:       value.Description,
				Pid:               0,
				AppId:             AppId,
				SourceFileName:    value.SourceFileName,
				RouteName:         value.Name,
				RouteNameMd5Value: RouteNameMd5Value,
			}
			db.Model(&SysAppModel.Authority{}).FirstOrCreate(&AuthorityMain, sysmdel.Authority{Description: value.Description, Pid: 0, SourceFileName: value.SourceFileName})
			var Authority SysAppModel.Authority
			var isAuthority SysAppModel.Authority
			db.Where(&SysAppModel.Authority{Description: value.Description}).First(&Authority)
			if Authority.Id > 0 {
				Md5ValueFile := fmt.Sprintf("%s_%s", value.MainHandlerName, value.RegisterFileName)
				Md5Value := simple.MD5(Md5ValueFile)
				db.Where("pid > ? and md5_value = ?", 0, simple.MD5(Md5Value)).Find(&isAuthority)
				//debug.Dd(Authority)
				if isAuthority.Id <= 0 {
					createAuthority := &SysAppModel.Authority{
						Pid:               Authority.Id,
						AppId:             AppId,
						Name:              value.MainHandlerName, //simple.Split(value.MainHandlerName, "."),
						Description:       value.Description,
						RegisterFileName:  value.RegisterFileName,
						MainHandlerName:   value.MainHandlerName,
						Method:            value.Method,
						FormattedPath:     value.FormattedPath,
						StaticPath:        value.StaticPath(),
						Path:              value.Path,
						SourceFileName:    value.SourceFileName,
						RouteName:         value.Name,
						RouteNameMd5Value: RouteNameMd5Value,
						Status:            1,
						Md5Value:          simple.MD5(Md5Value),
					}
					db.Create(createAuthority)
				}
			}
			consoleSTR := ""
			consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-4d", index)
			consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-70s", value.MainHandlerName)
			consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-5s", value.Method)
			//consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-45s", value.Path)
			consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-30s", value.StaticPath())
			//consoleSTR = consoleSTR + console.Sprintf(console.CODE_SUCCESS, "%-20s", value.RegisterFileName)
			console.Println(console.CODE_SUCCESS, " "+" "+consoleSTR)
			index++
		}
	}
	console.Println(console.CODE_WARNING, " "+console.Sprintf(console.CODE_WARNING, "%s listening on: http://localhost%s", AppName, Port))
}
