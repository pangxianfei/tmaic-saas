package listeners

import "tmaic/vendors/framework/hub"

func Initialize() {
	// 初始化主题和频道
	hub.RegisterQueue()
}
