package io

import (
	"encoding/json"
	"fmt"
	"github.com/RudyChow/proxy/config"
	"github.com/RudyChow/proxy/models"
	"github.com/go-redis/redis"
	"os"
)

type RedisDriver struct {
	client *redis.Client
}

const (
	REDIS_PROXYPOOL       = "proxypool"
	REDIS_USEFULPROXYPOOL = "usefulproxypool"
)

//保存数据到代理池中
func (this *RedisDriver) SaveData2ProxyPool(proxy *models.Proxy) {
	data, _ := json.Marshal(proxy)
	err := this.client.SAdd(REDIS_PROXYPOOL, string(data)).Err()
	if err != nil {
		fmt.Println(err)
	}
}

//从代理池中获取所有代理数据
func (this *RedisDriver) GetDataFromProxyPool() []*models.Proxy {
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
func (this *RedisDriver) CountProxyPool() int64 {
	return this.client.SCard(REDIS_PROXYPOOL).Val()
}

//可用代理的数量
func (this *RedisDriver) CountUsefulProxy() int64 {
	return this.client.ZCard(REDIS_USEFULPROXYPOOL).Val()
}

//存储数据到可用的代理池中
func (this *RedisDriver) SaveData2UsefulProxyPool(proxy *models.Proxy, score float64) {
	this.client.ZAdd(REDIS_USEFULPROXYPOOL, redis.Z{Score: score, Member: proxy.GetLink()})
}

//从可用的代理池中删除数据
func (this *RedisDriver) RemoveDataFromProxyPool(proxy *models.Proxy) {
	this.client.ZRem(REDIS_USEFULPROXYPOOL, proxy.GetLink())
}

//获取可用代理
func (this *RedisDriver) GetShortcutFromUsefulProxyPool(count int64) []models.ProxyShortcut {
	zArr, _ := this.client.ZRangeByScoreWithScores(REDIS_USEFULPROXYPOOL, redis.ZRangeBy{Min: "0", Max: "5000", Offset: 0, Count: count}).Result()

	var proxyShortcut []models.ProxyShortcut
	for _, z := range zArr {
		addr := z.Member.(string)
		proxyShortcut = append(proxyShortcut, models.ProxyShortcut{Addr: addr, Speed: int(z.Score)})
	}
	return proxyShortcut
}

//获取质量最好的可用代理
func (this *RedisDriver) GetBestUsefulProxyPool() models.ProxyShortcut {
	shorcutArr := this.GetShortcutFromUsefulProxyPool(1)
	if len(shorcutArr) == 0 {
		return models.ProxyShortcut{}
	}
	return this.GetShortcutFromUsefulProxyPool(1)[0]
}

//新建一个redis客户端
func newRedis(redisConf config.Redis) (redisDriver *RedisDriver) {

	redisDriver = &RedisDriver{}
	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Auth, // no password set
		DB:       redisConf.Db,   // use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		fmt.Println("redis connect failed")
		os.Exit(1)
	}
	redisDriver.client = client
	return
}
