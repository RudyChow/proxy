package http

import (
	"github.com/RudyChow/proxy/config"
	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	gin.SetMode(config.Conf.Http.Mode)
	r := gin.Default()
	registerRouters(r)
	r.Run(config.Conf.Http.Addr) // listen and serve on 0.0.0.0:8080
}
