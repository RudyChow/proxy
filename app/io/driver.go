package io

import (
	"github.com/RudyChow/proxy/app/models"
	"github.com/RudyChow/proxy/config"
	"log"
	"os"
)

var Handler Driver

func init() {
	switch config.Conf.IO.Driver {
	case "redis":
		Handler = newRedis(config.Conf.IO.Redis)
	case "mysql":
		Handler = newDb(config.Conf.IO.Mysql)
	default:
		log.Println("unsupported io driver ")
		os.Exit(1)
	}

}

type Driver interface {
	GetDataFromProxyPool() []*models.Proxy
	GetShortcutFromUsefulProxyPool(count int64) []models.ProxyShortcut
	GetBestUsefulProxyPool() models.ProxyShortcut
	SaveData2ProxyPool(proxy *models.Proxy)
	SaveData2UsefulProxyPool(proxy *models.Proxy, score int64)
	RemoveDataFromProxyPool(proxy *models.Proxy)
	CountProxyPool() int64
	CountUsefulProxy() int64
}
