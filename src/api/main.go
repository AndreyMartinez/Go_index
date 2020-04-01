package main

import (
	"goLagash/src/api/internal/implement/helloInterface"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", helloInterface.GetNoteHandler)
	router.Run()
}
