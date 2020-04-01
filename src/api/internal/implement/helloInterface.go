package implement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type helloWorld interface {
	GetNoteHandler(c *gin.Context)
}

func GetNoteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "pong"})
}
