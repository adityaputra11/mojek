package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ModelUUID struct {
	UUID uuid.UUID
}

type User struct {
	// Model
	UserID    uuid.UUID `json:"userid" gorm:"primaryKey;type:uuid;column:userid;" binding:"required,uuid"`
	Username  string    `json:"username" form:"username"`
	Email     string    `json:"email" form:"email"`
	Phone     string    `json:"phone" form:"phone"`
	FirstName string    `json:"first_name" form:"first_name"`
	LastName  string    `json:"last_name" form:"last_name"`
	Password  string    `json:"password,oempty"form:"password"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("userid", uuid.NewV4().String())
	return nil
}
