package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		// gin.H is a type alias for map[string] interface{}{"":""}
		context.JSON(http.StatusOK, gin.H{"data": "hello world!"})
	})
	_ = router.Run()
}
