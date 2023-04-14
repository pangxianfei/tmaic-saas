package request

import "gitee.com/pangxianfei/library/tmaic"

type MenuRequest struct {
	AppId int64 `validate:"required" json:"AppId"` // 应用ID
}

func (r *MenuRequest) Messages() map[string][]string {
	messages := tmaic.M{
		"AppId": []string{
			"required:应用ID必须项",
		},
	}
	return messages
}
