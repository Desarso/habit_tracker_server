package main

import (
	"net/http"

	"github.com/desarso/habit_tracker_server/docs"
	"github.com/desarso/habit_tracker_server/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	r := router.Group("/")

	//redirect might not work cause chrome caches the redirect
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/v1/docs/index.html")
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
	))

	routes.Habit_Routes(r.Group("/habits"))

	router.Run(":8080")
}
