package request

import "gitee.com/pangxianfei/library/tmaic"

type UserRequset struct {
	UserName   string `validate:"required" json:"UserName"`           // 用户名
	Mobile     string `validate:"required" json:"Mobile"`             // 手机
	Password   string `validate:"required,min=8" json:"Password"`     // 密码
	RePassword string `validate:"eqfield=Password" json:"RePassword"` // 重复密码
}

func (r *UserRequset) Messages() map[string][]string {
	messages := tmaic.M{
		"UserName": []string{
			"required:用户名必须项",
		},
		"Mobile": []string{
			"required:手机号必须项",
		},
		"Password": []string{
			"required:密码为必填项",
			"min:密码长度必须大于或者等于8位以上",
		},
		"RePassword": []string{
			"required:重复密码必填项",
			"eqfield:重复密码不正确",
		},
	}
	return messages
}
