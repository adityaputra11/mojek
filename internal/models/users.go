package models

type User struct {
	Model
	UserID     string `json:"userid" gorm:"primary_key"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Password   string `json:"password,oempty"`
	CreateAt   string `json:"create_at,oempty"`
	ModifiedAt string `json:"modified_at,oempty"`
	CreateBy   string `json:"create_by,oempty"`
	ModifiedBy string `json:"modified_by,oempty"`
}
