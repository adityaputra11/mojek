package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterRouterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	userAPI := NewUserAPI(db)
	router.GET("/users", userAPI.GetAllUser)
	router.GET("/users/:name", userAPI.GetUser)
}
