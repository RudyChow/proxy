package http

import (
	"github.com/RudyChow/proxy/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

func StartHttpServer() {
	r := gin.Default()
	registerRouters(r)
	r.Run(":" + strconv.Itoa(int(config.Conf.Http.Port))) // listen and serve on 0.0.0.0:8080
}
