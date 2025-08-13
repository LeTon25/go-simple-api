package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id;type:int;not null;primaryKey" json:"id"`
	RoleName string `gorm:"type:varchar(255);not null;unique" json:"role_name"`
	RoleNote string `gorm:"type:varchar(255);" json:"role_note"`
}

func (r *Role) TableName() string {
	return "roles"
}
