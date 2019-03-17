package models

import (
	"strconv"
)

type Proxy struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Ip       string `json:"ip" gorm:"type:varchar(16);unique_index:link_idx"`
	Port     int    `json:"port" gorm:"type:int;unique_index:link_idx"`
	Protocol string `json:"protocol" gorm:"type:varchar(10);unique_index:link_idx"`
	Score    int    `json:"score" gorm:"type:int"`
}

type ProxyShortcut struct {
	Addr  string `json:"addr"`
	Speed int    `json:"speed"`
}

//生成代理连接，如http://127.0.0.1:1080
func (this *Proxy) GetLink() string {
	return this.Protocol + "://" + this.Ip + ":" + strconv.Itoa(this.Port)
}

func (this *Proxy) GetShortcut() ProxyShortcut {
	return ProxyShortcut{this.GetLink(), this.Score}
}
