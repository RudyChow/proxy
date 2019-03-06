package spiders

import (
	"github.com/RudyChow/proxy/config"
	"log"
	"os"
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
	log.Println("[start crawling]")
	//开始采集数据
	for _, spider := range spiders {
		go spider.Crawl()
	}
}
