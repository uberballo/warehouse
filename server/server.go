package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uberballo/webstore/routes"
)

func Start() {
	router := gin.Default()
	routes.Routes(router)
	// Start and run the server
	router.Run(":5000")
}
