package v1

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type UserAPI interface {
	GetAllUser(c *gin.Context)
	GetUser(c *gin.Context)
}

// type userAPI struct {
// 	userService services.UserService
// }

func init() {
	log.Out = os.Stdout
	log.Debug("Init user")
}

// func NewUserAPI(db *gorm.DB) UserAPI {
// 	return &userAPI{userService: services.NewUserService(db)}
// }

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test user",
	})
}
