package controllers

import (
	"github.com/gin-gonic/gin"
	"lcrm2/helpers"
	"lcrm2/models"
	"net/http"
	"time"
)

// GetAllProject Нужно убрать из выдачи хэши паролей
func GetAllProject(c *gin.Context) {
	var project []models.Project

	err := models.GetAllProject(&project)
	if err != nil {
		helpers.ResponseJSON(c, 404, project)
		return
	}

	helpers.ResponseJSON(c, 200, project)
}

func GetUsers(c *gin.Context) {
	var u []models.User

	err := models.GetAllUsers(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range u {
		u[i].PrepareGive()
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type MemberInput struct {
	ProjectId string `json:"project_id" binding:"required"`
	UserId    string `json:"user_id" binding:"required"`
}

func RemoveMember(c *gin.Context) {
	var memberInput MemberInput
	err := c.ShouldBindJSON(&memberInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Получаем проект по идентификатору
	var project models.Project
	err = models.GetProjectById(&project, memberInput.ProjectId)

	// Получаем пользователя по идентификатору
	var user models.User
	err = models.GetUserByID(&user, memberInput.UserId)

	if user.RoleId != 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The member must be a student"})
		return
	}

	// Удаляем пользователя из проекта
	err = models.RemoveMember(&project, &user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Member removed"})
}

func AddMember(c *gin.Context) {
	var memberInput MemberInput
	err := c.ShouldBindJSON(&memberInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Получаем проект по идентификатору
	var project models.Project
	err = models.GetProjectById(&project, memberInput.ProjectId)

	// Получаем пользователя по идентификатору
	var user models.User
	err = models.GetUserByID(&user, memberInput.UserId)

	if user.RoleId != 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The member must be a student"})
		return
	}

	// Добавляем пользователя к проекту
	err = models.AddNewMember(&project, &user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Member added"})
}

type ProjectInput struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Deadline    time.Time `json:"deadline"`
}

func CreateProject(c *gin.Context) {
	var project models.Project
	var projectInput ProjectInput

	err := c.ShouldBindJSON(&projectInput)

	project.Title = projectInput.Title
	project.Description = projectInput.Description
	project.Deadline = projectInput.Deadline

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creatorIdInterface, _ := c.Get("userId")
	creatorId, ok := creatorIdInterface.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token error"})
		return
	}

	project.CreatorId = creatorId
	err = models.AddNewProject(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project created"})
}

func GetProjectById(c *gin.Context) {
	id := c.Params.ByName("id")
	var project models.Project
	err := models.GetProjectById(&project, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if project.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": project})
}

func UpdateProjectById(c *gin.Context) {
	id := c.Params.ByName("id")
	var project models.Project
	err := models.GetProjectById(&project, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&project)
	err = models.UpdateProjectById(&project)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": project})
}

func DeleteProjectById(c *gin.Context) {
	id := c.Params.ByName("id")
	var project models.Project
	err := models.DeleteProjectById(&project, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
