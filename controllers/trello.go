package controllers

import (
	"github.com/gin-gonic/gin"
	"lcrm2/models"
	"net/http"
	"strconv"
)

type TaskInput struct {
	Title          string `json:"title"`
	Name           string `json:"name"`
	Status         string `json:"status"`
	BoardID        uint   `json:"boardID"`
	AssignedUserID uint   `json:"assignedUserID"`
}

func CreateTask(c *gin.Context) {
	var task models.Task
	var taskInput TaskInput

	err := c.ShouldBindJSON(&taskInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.Title = taskInput.Title
	task.Name = taskInput.Name
	task.Status = taskInput.Status
	task.BoardID = taskInput.BoardID
	task.AssignedUserID = taskInput.AssignedUserID

	err = models.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task created"})
}

func UpdateTaskById(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task
	err := models.GetTaskById(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = models.UpdateTaskById(&task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": task})
}

func DeleteTaskById(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task
	err := models.DeleteTaskById(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

//func GetTask(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//	task, err := models.GetTask(uint(id))
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, task)
//}

type BoardInput struct {
	Name        string ` json:"name"`
	Description string ` json:"description"`
	ProjectID   uint   `json:"projectID"`
}

func CreateBoard(c *gin.Context) {
	var board models.Board
	var boardInput BoardInput

	err := c.ShouldBindJSON(&boardInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	board.Name = boardInput.Name
	board.Description = boardInput.Description
	board.ProjectID = boardInput.ProjectID

	err = models.CreateBoard(&board)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Board created"})
}

func GetBoardById(c *gin.Context) {
	id := c.Params.ByName("id")
	var board models.Board

	if err := models.GetBoardById(&board, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": board})
}

// Implement other Task related methods like CreateTask, UpdateTask, DeleteTask, etc. here

func CreateRating(c *gin.Context) {
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateRating(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rating)
}

func GetRatings(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ratings, err := models.GetRatings(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ratings)
}
