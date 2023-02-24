package migrations

import (
	"gorm.io/gorm"

	"tmaic/vendors/framework/database/migration"
	"tmaic/vendors/framework/helpers/zone"
	"tmaic/vendors/framework/model"
)

func init() {
	migration.AddMigrator(&CreateUserTable1548750742{})
}

type User struct {
	ID        int64      `gorm:"column:user_id;primary_key;auto_increment"`
	TenantsId *int       `gorm:"column:tenants_id;type:int(11)"`
	Name      *string    `gorm:"column:user_name;type:varchar(255)"`
	Email     *string    `gorm:"column:user_email;type:varchar(255);unique_index;not null"`
	Password  *string    `gorm:"column:user_password;type:varchar(255);not null"`
	CreatedAt zone.Time  `gorm:"column:user_created_at"`
	UpdatedAt zone.Time  `gorm:"column:user_updated_at"`
	DeletedAt *zone.Time `gorm:"column:user_deleted_at"`
	model.BaseModel
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
