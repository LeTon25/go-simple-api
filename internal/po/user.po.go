package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid;type:varchar(255);not null;index:idx_uuid;unique" json:"uuid"`
	UserName string    `gorm:"type:varchar(255);not null;unique" json:"user_name"`
	IsActive bool      `gorm:"type:boolean;default:true" json:"is_active"`
	Roles    []Role    `gorm:"many2many:user_roles" json:"roles"`
}

func (u *User) TableName() string {
	return "users"
}
