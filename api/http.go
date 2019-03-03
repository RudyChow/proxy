package api

import (
	"github.com/RudyChow/proxy/io"
	"github.com/gin-gonic/gin"
	"strconv"
)

func StartHttpServer() {
	r := gin.Default()
	r.GET("/proxy", getBest)
	r.GET("/proxy/:count", getUserfulProxyList)
	r.Run() // listen and serve on 0.0.0.0:8080
}

//拿最好的代理出来
func getBest(c *gin.Context) {
	shortcut := io.Handler.GetBestUsefulProxyPool()

	c.JSON(200, gin.H{
		"addr":  shortcut.Addr,
		"score": shortcut.Speed,
	})
}

//获取有用的代理列表
func getUserfulProxyList(c *gin.Context) {
	param := c.Param("count")
	count, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "unsupported param",
		})
	} else {
		shortcutArr := io.Handler.GetShortcutFromUsefulProxyPool(int64(count))

		c.JSON(200, gin.H{
			"data": shortcutArr,
		})
	}
}
