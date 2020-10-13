package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func testHelloWorld() {
}

func RegisterRouterAPIV1(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/users", GetUser)
}
