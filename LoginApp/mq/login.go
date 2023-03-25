package MQ

import (
	"gitee.com/pangxianfei/framework/work"
	"gitee.com/pangxianfei/saas/sysmodel"

	"tmaic/LoginApp/jobs"
	"tmaic/LoginApp/jobs/proto3/protomodel"
)

var LoginMessage = new(LoginDispatch)

type loginMessage interface {
	Dispatch(adminInfo *sysmdel.PlatformAdmin)
}
type LoginDispatch struct {
}

// Dispatch 推送登陆消息队列
func (l *LoginDispatch) Dispatch(adminInfo *sysmdel.PlatformAdmin) {

	LoginJob := jobs.LoginJob
	LoginJob.SetParam(&protomodel.LoginJob{
		UserName: adminInfo.UserName,
		Mobile:   adminInfo.Mobile,
		TenantId: adminInfo.TenantId,
		UserType: adminInfo.UserType,
	})
	if jobErr := work.Dispatch(LoginJob); jobErr != nil {

	}
}
