package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/uberballo/warehouse/server/routes"
	"github.com/uberballo/warehouse/server/services/product"
)

func Start() {
	router := gin.Default()
	routes.Routes(router)
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(5).Minutes().Do(product.InitializeProductData)
	s1.StartAsync()
	// Start and run the server
	router.Run(":8080")
}
