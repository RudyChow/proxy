package http

import (
	"github.com/RudyChow/proxy/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func registerRouters(r *gin.Engine) {
	registerApiRoutes(r)
}

func registerApiRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/proxy", controllers.GetBest)
		v1.GET("/proxy/:count", controllers.GetUserfulProxyList)
	}
}
