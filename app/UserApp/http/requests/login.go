package requests

// UserLogin 登陆验证器
type UserLogin struct {
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required,min=8,max=24"`
}
