package http

import (
	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	r := gin.Default()
	registerRouters(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

