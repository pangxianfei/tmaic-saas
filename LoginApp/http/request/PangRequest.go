package request

import (
	"gitee.com/pangxianfei/library/tmaic"
)

type PangRequest struct {
	// Name        string `validate:"required" json:"name"`
	// Description string `validate:"required" json:"description"`
}

func (r *PangRequest) Messages() map[string][]string {
	messages := tmaic.M{
		// "Name": []string{
		//     "required:名称为必填项",
		//     "min:名称长度需至少 2 个字",
		//     "max:名称长度不能超过 8 个字",
		// },
		// "Description": []string{
		//     "min:描述长度需至少 3 个字",
		//     "max:描述长度不能超过 255 个字",
		// },
	}
	return messages
}
