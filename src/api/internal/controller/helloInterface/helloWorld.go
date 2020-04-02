package helloInterface

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Content struct {
	helloWorldInterface
}

func (content Content) GetNoteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "pong"})
}
