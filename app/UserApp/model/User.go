package UserAppModel

type User struct {
	Id               int64  `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	TenantId         int64  `gorm:"size:11;not null" json:"TenantId" form:"TenantId"`
	Mobile           string `gorm:"size:11;unique;" json:"mobile" form:"mobile"`                        // 手机
	UserName         string `gorm:"size:100;" json:"username" form:"username"`                          // 用户名
	Email            string `gorm:"size:255;" json:"email" form:"email"`                                // 邮箱
	EmailVerified    bool   `gorm:"not null;default:false" json:"emailVerified" form:"emailVerified"`   // 邮箱是否验证
	Nickname         string `gorm:"size:255;" json:"nickname" form:"nickname"`                          // 昵称
	Avatar           string `gorm:"type:text" json:"avatar" form:"avatar"`                              // 头像
	BackgroundImage  string `gorm:"type:text" json:"backgroundImage" form:"backgroundImage"`            // 个人中心背景图片
	Password         string `gorm:"size:512" json:"-" form:"password"`                                  // 密码
	HomePage         string `gorm:"size:1024" json:"homePage" form:"homePage"`                          // 个人主页
	Description      string `gorm:"type:text" json:"description" form:"description"`                    // 个人描述
	Score            int    `gorm:"not null;" json:"score" form:"score"`                                // 积分
	Status           int    `gorm:"not null" json:"status" form:"status"`                               // 状态
	TopicCount       int    `gorm:"not null" json:"topicCount" form:"topicCount"`                       // 帖子数量
	CommentCount     int    `gorm:"not null" json:"commentCount" form:"commentCount"`                   // 跟帖数量
	Roles            string `gorm:"type:text" json:"roles" form:"roles"`                                // 角色
	Type             int    `gorm:"not null" json:"type" form:"type"`                                   // 用户类型
	ForbiddenEndTime int64  `gorm:"not null;default:0" json:"forbiddenEndTime" form:"forbiddenEndTime"` // 禁言结束时间
	CreateTime       int64  `json:"createTime" form:"createTime"`                                       // 创建时间
	UpdateTime       int64  `json:"updateTime" form:"updateTime"`                                       // 更新时间
}
