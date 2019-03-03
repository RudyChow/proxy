package filters

import (
	"github.com/RudyChow/proxy/io"
	"github.com/RudyChow/proxy/models"
	"net/http"
	"net/url"
	"time"
)

func UpdateUsefulProxy() {

	for {
		proxypool := io.Handler.GetDataFromProxyPool()

		for _, proxy := range proxypool {
			go checkSpeed(proxy)
		}
		//每分钟执行一次
		time.Sleep(1 * time.Minute)
	}

}

//测速
func checkSpeed(proxy *models.Proxy) {

	proxyAddr := proxy.GetLink()

	urli := url.URL{}
	urlproxy, _ := urli.Parse(proxyAddr)

	c := http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}

	startTime := time.Now().UnixNano() / 1e6
	res, err := c.Get("https://httpbin.org/get")
	if err != nil {
		//在可用代理池中删除该代理
		go io.Handler.RemoveDataFromProxyPool(proxy)
		return
	}
	endTime := time.Now().UnixNano() / 1e6
	defer res.Body.Close()
	//把可用的代理进行持久化
	go io.Handler.SaveData2UsefulProxyPool(proxy, float64(endTime-startTime))
}
