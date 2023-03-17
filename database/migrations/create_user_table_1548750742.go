package migrations

import (
	"gorm.io/gorm"

	"gitee.com/pangxianfei/framework/database/migration"
	"gitee.com/pangxianfei/framework/kernel/zone"
	"gitee.com/pangxianfei/framework/model"
)

func init() {
	migration.AddMigrator(&CreateUserTable1548750742{})
}

type User struct {
	model.BaseModel
	ID        int64      `gorm:"column:id;primary_key;auto_increment"`
	TenantsId *int       `gorm:"column:tenants_id;type:int(11)"`
	Name      *string    `gorm:"column:name;type:varchar(255)"`
	Email     *string    `gorm:"column:email;type:varchar(255);unique_index;not null"`
	Password  *string    `gorm:"column:password;type:varchar(255);not null"`
	CreatedAt zone.Time  `gorm:"column:created_at"`
	UpdatedAt zone.Time  `gorm:"column:updated_at"`
	DeletedAt *zone.Time `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return u.SetTableName("user")
}

type CreateUserTable1548750742 struct {
	migration.MigratorIdentify
	migration.MigrationUtils
}

func (*CreateUserTable1548750742) Up(db *gorm.DB) *gorm.DB {
	_ = db.Migrator().CreateTable(&User{})
	return db
}

func (*CreateUserTable1548750742) Down(db *gorm.DB) *gorm.DB {

	_ = db.Migrator().DropTable(&User{})

	return db
}
