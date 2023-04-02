package MQ

import (
	"gitee.com/pangxianfei/framework/queue/work"
	"gitee.com/pangxianfei/saas/sysmodel"

	"tmaic/LoginApp/jobs"
	"tmaic/LoginApp/jobs/proto3/protomodel"
)

var LoginMessage = new(LoginDispatch)

type loginMessage interface {
	Dispatch(adminInfo *sysmdel.PlatformAdmin) bool
}
type LoginDispatch struct{}

func (l *LoginDispatch) Dispatch(adminInfo *sysmdel.PlatformAdmin) bool {

	LoginJob := jobs.LoginJob
	LoginJob.SetParam(&protomodel.LoginJob{
		UserName: adminInfo.UserName,
		Mobile:   adminInfo.Mobile,
		TenantId: adminInfo.TenantId,
		UserType: adminInfo.UserType,
	})
	if jobErr := work.Dispatch(LoginJob); jobErr != nil {
		return false
	}
	return true
}
