package models

import (
	"strconv"
)

type Proxy struct {
	Ip       string `json:"ip"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type ProxyShortcut struct {
	Addr  string `json:"addr"`
	Speed int    `json:"speed"`
}

//生成代理连接，如http://127.0.0.1:1080
func (this *Proxy) GetLink() string {
	return this.Protocol + "://" + this.Ip + ":" + strconv.Itoa(this.Port)
}
