package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ModelUUID struct {
	UUID uuid.UUID
}

type User struct {
	Model
	UserID     uuid.UUID `json:"userid" gorm:"primaryKey;type:uuid;column:userid;" binding:"required,uuid"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Password   string    `json:"password,oempty"`
	CreateAt   string    `json:"create_at,oempty"`
	ModifiedAt string    `json:"modified_at,oempty"`
	CreateBy   string    `json:"create_by,oempty"`
	ModifiedBy string    `json:"modified_by,oempty"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("userid", uuid.NewV4().String())
	return nil
}
