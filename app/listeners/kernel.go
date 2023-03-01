package listeners

import "gitee.com/pangxianfei/framework/hub"

func Initialize() {
	// 初始化主题和频道
	hub.RegisterQueue()
}
