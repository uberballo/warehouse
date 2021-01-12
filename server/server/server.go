package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/uberballo/warehouse/server/routes"
	"github.com/uberballo/warehouse/server/services/product"
)

//Start begins fetching and serving bad api data
func Start() {
	router := gin.Default()
	routes.Routes(router)

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(5).Minutes().Do(product.InitializeProductData)
	scheduler.StartAsync()

	// Start and run the server
	router.Run()
}
