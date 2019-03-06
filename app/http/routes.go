package http

import (
	"github.com/RudyChow/proxy/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func registerRouters(r *gin.Engine) {
	r.GET("/proxy", controllers.GetBest)
	r.GET("/proxy/:count", controllers.GetUserfulProxyList)
}
