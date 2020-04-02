package helloInterface

import (
	"github.com/gin-gonic/gin"
)

type helloWorldInterface interface {
	GetNoteHandler(c *gin.Context)
}
