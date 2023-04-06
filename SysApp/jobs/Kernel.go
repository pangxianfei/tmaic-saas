package jobs

import "gitee.com/pangxianfei/framework/queue/work"

func Initialize() {
	// 初始化主题和频道
	work.RegisterQueue()
}
