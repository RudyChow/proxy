package config

import (
	"fmt"
	"github.com/spf13/viper"
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
	Password int
	Db       string
}

func init() {
	Conf = &Config{}
	Conf.initConfig()
}
//初始化数据
func (this *Config) initConfig() {

	fmt.Println("[init config params]......")
	//设置配置文件类型
	viper.SetConfigType("yaml")
	//设置配置文件名称（除去后缀）
	viper.SetConfigName(".env")
	//配置文件目录
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}

	err = viper.Unmarshal(&this)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}


	fmt.Println("[finish config params]......")
}
