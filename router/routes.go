package router

import (
	"github.com/RafaelRMJesus/op-api/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handlers.GetOpeningHandler)
		v1.POST("/opening", handlers.CreateOpeningHandler)
		v1.PUT("/opening", handlers.UpdateOpeningHandler)
		v1.DELETE("/opening", handlers.DeleteOpeningHandler)
		v1.GET("/openings", handlers.ListOpeningsHandler)
	}
}
