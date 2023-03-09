package requests

// UserLogin 登陆验证器
type UserLogin struct {
	TenantId int64  `json:"tenantId" binding:"required"`
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required,min=8,max=24"`
}
