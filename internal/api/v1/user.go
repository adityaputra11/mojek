package v1

import (
	"net/http"
	"os"
	"reflect"

	"github.com/adityaputra11/mojek/internal/models"
	"github.com/adityaputra11/mojek/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type UserAPI interface {
	GetAllUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type userAPI struct {
	userService services.UserService
}

func init() {
	log.Out = os.Stdout
	log.Debug("Init user")
}

func NewUserAPI(db *gorm.DB) UserAPI {
	return &userAPI{userService: services.NewUserService(db)}
}

func (p userAPI) GetAllUser(c *gin.Context) {
	users := p.userService.GetAllUser(0, -1, "")
	c.JSON(http.StatusOK, users)
}

// GetUser return only one User
func (p userAPI) GetUser(c *gin.Context) {
	username := c.Param("name")
	user := p.userService.GetUser(username)
	if reflect.DeepEqual(*user, models.User{}) {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}
