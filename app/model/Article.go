package model

// Article 文章
type Article struct {
	Id          int64  `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	UserId      int64  `gorm:"index:idx_article_user_id" json:"userId" form:"userId"`             // 所属用户编号
	Title       string `gorm:"size:255;not null;" json:"title" form:"title"`                      // 标题
	Summary     string `gorm:"type:text" json:"summary" form:"summary"`                           // 摘要
	Content     string `gorm:"type:longtext;not null;" json:"content" form:"content"`             // 内容
	ContentType string `gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`   // 内容类型：markdown、html
	Status      int    `gorm:"type:int(11);index:idx_article_status" json:"status" form:"status"` // 状态
	Share       bool   `gorm:"not null" json:"share" form:"share"`                                // 是否是分享的文章，如果是这里只会显示文章摘要，原文需要跳往原链接查看
	SourceUrl   string `gorm:"type:text" json:"sourceUrl" form:"sourceUrl"`                       // 原文链接
	ViewCount   int64  `gorm:"not null;index:idx_view_count;" json:"viewCount" form:"viewCount"`  // 查看数量
	CreateTime  int64  `gorm:"index:idx_article_create_time" json:"createTime" form:"createTime"` // 创建时间
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`                                      // 更新时间
}
