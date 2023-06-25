package users

import (
	"github.com/gin-gonic/gin"
	"github.com/MadhavVankina/go-react-auth-project/backend/domain/user"
)

func Register(c *gin.Context) {
	var user user.User

	c.ShouldBindJSON(&user)
}
