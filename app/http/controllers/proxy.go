package controllers

import (
	"github.com/RudyChow/proxy/app/io"
	"github.com/gin-gonic/gin"
	"strconv"
)

//拿最好的代理出来
func GetBest(c *gin.Context) {
	shortcut := io.Handler.GetBestUsefulProxyPool()

	c.JSON(200, gin.H{
		"addr":  shortcut.Addr,
		"score": shortcut.Speed,
	})
}

//获取有用的代理列表
func GetUserfulProxyList(c *gin.Context) {
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
