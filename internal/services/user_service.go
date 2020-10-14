package services

import (
	"github.com/adityaputra11/mojek/internal/models"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	GetAllUser(offset int, limit int, search string) *[]models.User
	GetUser(name string) *models.User

	InitUser() *models.User
	CreateUser(user *models.User) *models.User
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db}
}

// GetAllUser return all User
func (p userService) GetAllUser(offset int, limit int, search string) *[]models.User {
	var users []models.User
	p.db.Find(&users)
	return &users
}

// GetUser return only one User
func (p userService) GetUser(name string) *models.User {
	var user models.User
	p.db.First(&user, &models.User{Username: name})
	return &user
}

// InitUser return only one User
func (p userService) InitUser() *models.User {
	var user models.User
	return &user
}

// CreateUser return only one User
func (p userService) CreateUser(user *models.User) *models.User {
	p.db.Create(user)
	return user
}
