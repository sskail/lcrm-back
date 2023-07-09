package helpers

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

func ResponseJSON(w *gin.Context, status_code int, data interface{}) {
	log.Println("status code :", status_code)
	var message Message
	message.StatusCode = status_code
	message.Data = data
	w.JSON(200, message)
}
