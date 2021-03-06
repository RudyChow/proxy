package spiders

import (
	"github.com/RudyChow/proxy/app/io"
	"github.com/RudyChow/proxy/app/models"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Ip66 struct {
	url string
}

func (this *Ip66) Crawl() {
	//获取数据
	res, err := http.Get(this.url)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	bodyStr := string(body)

	//正则匹配ip
	expr := regexp.MustCompile(`(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})(\.(2(5[0-5]{1}|[0-4]\d{1})|[0-1]?\d{1,2})){3}\:([1-9]+)`)
	ips := expr.FindAllString(bodyStr, 100)
	for i := 0; i < len(ips); i++ {
		ipArr := strings.Split(ips[i], ":")
		port, _ := strconv.Atoi(ipArr[1])

		proxy := new(models.Proxy)
		proxy.Ip = ipArr[0]
		proxy.Port = port
		proxy.Protocol = "http"

		go io.Handler.SaveData2ProxyPool(proxy)
	}
}
