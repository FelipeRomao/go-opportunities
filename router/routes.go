package router

import (
	"github.com/FelipeRomao/go-opportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/openings/:id", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
	}

}
