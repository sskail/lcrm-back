package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lcrm2/models"
	"lcrm2/utils/token"
	"net/http"
	"strconv"
)

type RegisterInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	RoleId    string `json:"role"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	GitApi    string `json:"gitApi"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	u.FirstName = input.FirstName
	u.LastName = input.LastName
	u.Email = input.Email
	u.GitApi = input.GitApi

	var role models.Role
	if input.RoleId != "" {
		// Parse the role and look it up in the database
		err := models.GetRoleById(&role, input.RoleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
			return
		}
		u.RoleId = role.ID
	} else {
		// Use a default role if none was specified
		u.RoleId = 3
	}
	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var u models.User

	err = models.GetUserByID(&u, strconv.Itoa(int(userId)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func Logout(c *gin.Context) {
	_, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

func GetUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err = models.GetUserByID(&user, strconv.Itoa(int(userId)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
