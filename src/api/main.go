package main

import (
	"github.com/gin-gonic/gin"
	helloInterface "github.com/goLagash/src/api/internal/controller/helloInterface"
)

/**
main principal de Go
*/
func main() {
	router := gin.Default()
	content := helloInterface.Content{}
	router.GET("/ping", content.GetNoteHandler)
	router.Run()
}
