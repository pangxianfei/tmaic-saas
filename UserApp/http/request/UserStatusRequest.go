package request

import "gitee.com/pangxianfei/library/tmaic"

type UserStatusRequset struct {
	Status int64 `validate:"required,gt=0" json:"Status"`
}

func (r *UserStatusRequset) Messages() map[string][]string {
	messages := tmaic.M{
		"Status": []string{
			"required:状态为必填项",
			"gt:状态不正确",
		},
	}
	return messages
}
