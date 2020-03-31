package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNoteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "pong"})
}

func main() {
	router := gin.Default()
	router.GET("/ping", GetNoteHandler)
	router.Run()

}
