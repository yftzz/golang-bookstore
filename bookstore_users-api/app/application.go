package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication App entry point
func StartApplication() {
	mapUrls()
	router.Run(":8000")
}
