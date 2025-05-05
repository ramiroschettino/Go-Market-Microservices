package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(orderClient any) *gin.Engine {
	r := gin.Default()

	r.POST("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Aquí se creará una orden",
		})
	})

	return r
}
