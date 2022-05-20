package Router

import (
	controllers "TSLC/Controllers"

	"github.com/gin-gonic/gin"
)

// Router ...
func Router() {
	Router := gin.Default()
	r := Router.Group("/api")
	r.POST("divide", controllers.Divide)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	Router.Run(":8080")
}
