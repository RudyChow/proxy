package spiders

import (
	"log"
	"github.com/RudyChow/proxy/config"
	"github.com/RudyChow/proxy/io"
	"os"
	"time"
)

type Spider interface {
	Crawl()
}

var spiders []Spider

func init() {
	//一些判断，把无法爬的或者没有配置的爬虫信息去掉
	targetLen := len(config.Conf.Targets)
	if targetLen == 0 {
		log.Println("no target")
		os.Exit(1)
	}
	for _, target := range config.Conf.Targets {

		url, ok := config.Conf.Spriders[target]
		if !ok {
			log.Printf("target[%s] is not configured\n", target)
			continue
		}

		switch target {
		case "ip66":
			spiders = append(spiders, &Ip66{url})
		default:
			log.Printf("target %s can not be crawled\n", target)
		}
	}

}

func StartCrawl() {
	for {

		//当可用代理超过100个，则暂缓60分钟进行采集
		if io.Handler.CountUsefulProxy() >= 100 {
			time.Sleep(60 * time.Minute)
		}

		//开始采集数据
		for _, spider := range spiders {
			go spider.Crawl()
		}

		//1分钟采集一次
		time.Sleep(time.Duration(1) * time.Minute)
	}

}
