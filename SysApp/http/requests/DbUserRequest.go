package requests

// DbUserRequest 应用验证器
type DbUserRequest struct {
	UserId int64 `json:"user_id" validate:"required,gt=0"`
}
