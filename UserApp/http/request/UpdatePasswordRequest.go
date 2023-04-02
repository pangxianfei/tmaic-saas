package request

import (
	"gitee.com/pangxianfei/library/tmaic"
)

type UpdatePasswordRequest struct {
	OldPassword string `validate:"required,min=7" json:"OldPassword"`  // 原密码
	Password    string `validate:"required,min=7" json:"Password"`     // 密码
	RePassword  string `validate:"eqfield=Password" json:"RePassword"` // 重复密码
}

func (r *UpdatePasswordRequest) Messages() map[string][]string {
	messages := tmaic.M{
		"OldPassword": []string{
			"required:原密码为必填项",
			"min:密码长度必须大于或者等于8位以上",
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
