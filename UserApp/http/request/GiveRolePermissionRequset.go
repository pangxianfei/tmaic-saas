package request

import "gitee.com/pangxianfei/library/tmaic"

type GiveRolePermissionRequset struct {
	RoleId       int64 `validate:"required,gt=0,lte=130" json:"RoleId"`
	PermissionId int64 `validate:"required,gte=0,lte=130" json:"PermissionId"`
}

func (r *GiveRolePermissionRequset) Messages() map[string][]string {
	messages := tmaic.M{
		"RoleId": []string{
			"required:角色ID为必填项",
			"gt:角色ID必须大于 gt:",
			"lte:角色ID必须小于 lte:",
		},
		"PermissionId": []string{
			"required:权限ID为必填项",
			"gt:权限ID必须大于 gt:",
			"lte:权限ID必须小于 lte:",
		},
	}
	return messages
}
