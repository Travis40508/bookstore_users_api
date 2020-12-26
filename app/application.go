package app

import "github.com/gin-gonic/gin"

// only layer that will define and use http server.
// this way we can plug and play our http server if we ever want to swap it out with minimal refactor

var (
	router = gin.Default()
)
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
