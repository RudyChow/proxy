package main

import (
	"github.com/RudyChow/proxy/api"
	_ "github.com/RudyChow/proxy/config"
	"github.com/RudyChow/proxy/spiders"
	"github.com/RudyChow/proxy/utils/filters"
)

func main() {

	//更新可用的代理ip池
	go filters.UpdateUsefulProxy()

	//开启http服务
	go api.StartHttpServer()

	//开始采集数据
	spiders.StartCrawl()
}
