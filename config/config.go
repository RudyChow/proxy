package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	Conf *Config
)

type Config struct {
	//爬虫网站和url
	Spriders map[string]string
	//要爬哪一些网站
	Targets []string
	//持久层
	IO IO
}

type IO struct {
	Driver string
	Redis  *Redis
	Mysql  *Mysql
}

type Redis struct {
	Addr string `ini:"addr"`
	Auth string `ini:"auth"`
	Db   int    `ini:"db"`
}

type Mysql struct {
	Addr     string `ini:"addr"`
	User     string `ini:"auth"`
	Password int    `ini:"password"`
	Db       string `ini:"db"`
}

func init() {
	Conf = &Config{}

	cfg, err := ini.ShadowLoad("config/config.ini")
	if err != nil {
		fmt.Print(err)
	}

	//初始化爬虫网站相关的配置
	initSpridersConf(cfg)
	//初始化持久层相关配置
	initIOConf(cfg)
}

func initSpridersConf(cfg *ini.File) {
	Conf.Spriders = cfg.Section("spriders").KeysHash()
	Conf.Targets = cfg.Section("targets").Key("target").ValueWithShadows()
}

func initIOConf(cfg *ini.File) {
	Conf.IO.Driver = cfg.Section("io").Key("driver").String()
	//redis
	redisConf := &Redis{}
	if err := cfg.Section("redis").MapTo(redisConf); err != nil {
		fmt.Println(err)
	}
	Conf.IO.Redis = redisConf
	//mysql
	mysqlConf := &Mysql{}
	if err := cfg.Section("mysql").MapTo(mysqlConf); err != nil {
		fmt.Println(err)
	}
	Conf.IO.Mysql = mysqlConf
}
