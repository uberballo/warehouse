package routes

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	v1 "github.com/uberballo/warehouse/server/routes/v1"
)

//Routes adds routes to basic funtionality
func Routes(router *gin.Engine) {
	router.Use(static.Serve("/", static.LocalFile("./web", true)))
	router.Use(static.Serve("/facemasks", static.LocalFile("./web", true)))
	router.Use(static.Serve("/gloves", static.LocalFile("./web", true)))
	router.Use(static.Serve("/beanies", static.LocalFile("./web", true)))
	router.NoRoute(notFound)
	api := router.Group("/api")
	{
		api.GET("/", helloWorld)
		api.GET("/products/:category", v1.GetProducts)
	}
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Hello world",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
