package middlewares

import (
	"github.com/gin-gonic/gin"
	"lcrm2/models"
	"lcrm2/utils/token"
	"net/http"
	"strconv"
)

func roleInSlice(a uint, list []uint) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.ValidToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func JwtRoleMiddleware(roleList []uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := token.ExtractTokenID(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Not Permission 1")
			c.Abort()
			return
		}

		var u models.User

		err = models.GetUserByID(&u, strconv.Itoa(int(userId)))
		if err != nil {
			c.String(http.StatusUnauthorized, "Not Permission 2")
			c.Abort()
			return
		}

		if !roleInSlice(u.RoleId, roleList) {
			c.String(http.StatusUnauthorized, "Not Permission 3")
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Set("RoleId", u.RoleId)
		c.Next()
	}
}
