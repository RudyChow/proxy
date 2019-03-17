package config

import (
	"github.com/spf13/viper"
	"log"
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
	//web
	Http Http
}

type IO struct {
	Driver string
	Redis  Redis
	Mysql  Mysql
}

type Redis struct {
	Addr string
	Auth string
	Db   int
}

type Mysql struct {
	Addr     string
	User     string
	Password string
	Db       string
}

type Http struct {
	Port uint16
}

func init() {
	Conf = &Config{}
	Conf.initConfig()
}

//初始化数据
func (this *Config) initConfig() {

	log.Println("[init config params]......")
	//设置配置文件类型
	viper.SetConfigType("yaml")
	//设置配置文件名称（除去后缀）
	viper.SetConfigName(".env")
	//配置文件目录
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
		return
	}

	err = viper.Unmarshal(&this)
	if err != nil {
		log.Panicf("Can not Unmarshal config file: %s \n", err)
		return
	}

	log.Println("[finish config params]......")
}
