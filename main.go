package main

import (
	"net/http"

	"github.com/desarso/habit_tracker_server/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r := router.Group("/api/v1")

	//redirect might not work cause chrome caches the redirect
	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/v1/docs/index.html")
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
	))

	routes.Habit_Routes(r.Group("/habits"))

	router.Run(":8080")
}
