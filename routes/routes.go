package routes

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	v1 "github.com/uberballo/webstore/routes/v1"
)

func Routes(router *gin.Engine) {
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	router.Use(static.Serve("/facemasks", static.LocalFile("./client/build", true)))
	router.Use(static.Serve("/gloves", static.LocalFile("./client/build", true)))
	router.Use(static.Serve("/beanies", static.LocalFile("./client/build", true)))
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
