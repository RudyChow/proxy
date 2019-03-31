package io

import (
	"encoding/json"
	"github.com/RudyChow/proxy/app/models"
	"github.com/RudyChow/proxy/config"
	"github.com/go-redis/redis"
	"log"
)

type redisDriver struct {
	client *redis.Client
}

const (
	REDIS_PROXYPOOL       = "proxypool"
	REDIS_USEFULPROXYPOOL = "usefulproxypool"
)

//保存数据到代理池中
func (this *redisDriver) SaveData2ProxyPool(proxy *models.Proxy) {
	data, _ := json.Marshal(proxy)
	err := this.client.SAdd(REDIS_PROXYPOOL, string(data)).Err()
	if err != nil {
		log.Println(err)
	}
}

//从代理池中获取所有代理数据
func (this *redisDriver) GetDataFromProxyPool() []*models.Proxy {
	data, _ := this.client.SMembers(REDIS_PROXYPOOL).Result()

	var proxypool []*models.Proxy
	for _, v := range data {
		var proxy *models.Proxy
		json.Unmarshal([]byte(v), &proxy)
		proxypool = append(proxypool, proxy)
	}
	return proxypool
}

//代理池的数量
func (this *redisDriver) CountProxyPool() int64 {
	return this.client.SCard(REDIS_PROXYPOOL).Val()
}

//可用代理的数量
func (this *redisDriver) CountUsefulProxy() int64 {
	return this.client.ZCard(REDIS_USEFULPROXYPOOL).Val()
}

//存储数据到可用的代理池中
func (this *redisDriver) SaveData2UsefulProxyPool(proxy *models.Proxy, score int64) {
	this.client.ZAdd(REDIS_USEFULPROXYPOOL, redis.Z{Score: float64(score), Member: proxy.GetLink()})
}

//从可用的代理池中删除数据
func (this *redisDriver) RemoveDataFromProxyPool(proxy *models.Proxy) {
	this.client.ZRem(REDIS_USEFULPROXYPOOL, proxy.GetLink())
}

//获取可用代理
func (this *redisDriver) GetShortcutFromUsefulProxyPool(count int64) []models.ProxyShortcut {
	zArr, _ := this.client.ZRangeByScoreWithScores(REDIS_USEFULPROXYPOOL, redis.ZRangeBy{Min: "0", Max: "5000", Offset: 0, Count: count}).Result()

	var proxyShortcut []models.ProxyShortcut
	for _, z := range zArr {
		addr := z.Member.(string)
		proxyShortcut = append(proxyShortcut, models.ProxyShortcut{Addr: addr, Score: int(z.Score)})
	}
	return proxyShortcut
}

//获取质量最好的可用代理
func (this *redisDriver) GetBestUsefulProxyPool() models.ProxyShortcut {
	shorcutArr := this.GetShortcutFromUsefulProxyPool(1)
	if len(shorcutArr) == 0 {
		return models.ProxyShortcut{}
	}
	return this.GetShortcutFromUsefulProxyPool(1)[0]
}

//新建一个redis客户端
func newRedis(redisConf config.Redis) (driver *redisDriver) {

	driver = &redisDriver{}
	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Auth, // no password set
		DB:       redisConf.Db,   // use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Panic("redis connect failed")
	}
	driver.client = client
	return
}
