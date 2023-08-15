package router

import (
	docs "github.com/FelipeRomao/go-opportunities/docs"
	"github.com/FelipeRomao/go-opportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"

	docs.SwaggerInfo.Title = "Go Opportunities API"
	docs.SwaggerInfo.Description = "This is a sample server for Go Opportunities API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/openings/:id", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
