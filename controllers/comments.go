package controllers

import (
	"github.com/gin-gonic/gin"
	"lcrm2/models"
	"net/http"
	"strconv"
)

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

type CommentInput struct {
	Body      string ` json:"body" binding:"required"`
	ProjectID uint   `json:"projectID" binding:"required"`
	//CreatorId string
}

// CreateComment Обработчик для создания комментария
func CreateComment(c *gin.Context) {
	var comment models.Comment
	var commentInput CommentInput

	if err := c.ShouldBindJSON(&commentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creatorIdInterface, _ := c.Get("userId")
	creatorId, ok := creatorIdInterface.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token error"})
		return
	}

	comment.Body = commentInput.Body
	comment.ProjectID = commentInput.ProjectID
	comment.CreatorId = creatorId

	if err := models.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetCommentByID Обработчик для получения комментария по ID
func GetCommentByID(c *gin.Context) {
	id := c.Param("id")
	comment, err := models.GetCommentsByProjectID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

type ReplyInput struct {
	Body      string ` json:"body" binding:"required"`
	CommentID uint   `json:"commentID" binding:"required"`
	//CreatorId string
}

// CreateReply Обработчик для создания ответа на комментарий
func CreateReply(c *gin.Context) {
	var reply models.Reply
	var replyInput ReplyInput
	if err := c.ShouldBindJSON(&replyInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creatorIdInterface, _ := c.Get("userId")
	creatorId, ok := creatorIdInterface.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token error"})
		return
	}

	reply.Body = replyInput.Body
	reply.CommentID = replyInput.CommentID
	reply.CreatorId = creatorId
	if err := models.CreateReply(&reply); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reply"})
		return
	}

	c.JSON(http.StatusCreated, reply)
}

// GetRepliesByCommentID Обработчик для получения ответов по ID комментария
func GetRepliesByCommentID(c *gin.Context) {
	commentID := c.Param("id")
	replies, err := models.GetRepliesByCommentID(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get replies"})
		return
	}

	c.JSON(http.StatusOK, replies)
}

// GetReplyByID Обработчик для получения ответа по ID
func GetReplyByID(c *gin.Context) {
	id := c.Param("id")
	reply, err := models.GetRepliesByCommentID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
		return
	}

	c.JSON(http.StatusOK, reply)
}
