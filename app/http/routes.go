package http

import (
	"github.com/gin-gonic/gin"
	"github.com/RudyChow/proxy/app/http/controllers"
)

func registerRouters(r *gin.Engine)  {
	r.GET("/proxy", controllers.GetBest)
	r.GET("/proxy/:count", controllers.GetUserfulProxyList)
}
