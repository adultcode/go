package api

import (
	"github.com/gin-gonic/gin"
	"golan-clean-web-api/api/middlewares"
	"golan-clean-web-api/config"
	"golan-clean-web-api/routers"
)

func InitServer() {

	cfg := config.GetConfig()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.TestMiddleWare())

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
		//v1.GET("/health", func(c *gin.Context) {
		//})
	}

	//r.Run(fmt.Sprint(":%s", cfg.Server.Port))
	r.Run(":" + cfg.Server.Port) // For string port
}
