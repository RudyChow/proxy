package io

import (
	"fmt"
	"github.com/RudyChow/proxy/app/models"
	"github.com/RudyChow/proxy/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type mysqlDriver struct {
	db *gorm.DB
}

func (this *mysqlDriver) GetDataFromProxyPool() []*models.Proxy {
	var arr = []*models.Proxy{}
	this.db.Find(&arr)
	return arr
}
func (this *mysqlDriver) GetShortcutFromUsefulProxyPool(count int64) []models.ProxyShortcut {
	var shortcut = []models.ProxyShortcut{}
	var arr = []*models.Proxy{}

	this.db.Where("score > ?", 0).Order("score asc").Limit(count).Find(&arr)
	for _, proxy := range arr {
		shortcut = append(shortcut, proxy.GetShortcut())
	}

	return shortcut
}
func (this *mysqlDriver) GetBestUsefulProxyPool() models.ProxyShortcut {
	proxy := &models.Proxy{}
	this.db.Where("score > ?", 0).Order("score asc").First(&proxy)
	log.Println(proxy)
	return proxy.GetShortcut()
}
func (this *mysqlDriver) SaveData2ProxyPool(proxy *models.Proxy) {
	this.db.Create(&proxy)
}
func (this *mysqlDriver) SaveData2UsefulProxyPool(proxy *models.Proxy, score int64) {
	proxy.Score = int(score)
	this.db.Save(&proxy)
}
func (this *mysqlDriver) RemoveDataFromProxyPool(proxy *models.Proxy) {
	proxy.Score = 0
	this.db.Save(&proxy)
}
func (this *mysqlDriver) CountProxyPool() int64 {
	var count int64
	this.db.Model(&models.Proxy{}).Count(&count)
	return count
}
func (this *mysqlDriver) CountUsefulProxy() int64 {
	var count int64
	this.db.Model(&models.Proxy{}).Where("score > 0").Count(&count)
	return count
}

func newDb(mysqlConf config.Mysql) (driver *mysqlDriver) {
	arg := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Addr, mysqlConf.Db)
	fmt.Println(arg)
	db, err := gorm.Open("mysql", arg)
	//db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/proxy")
	if err != nil {
		log.Println(err)
		log.Panic("mysql connect failed")
	}

	driver = &mysqlDriver{}
	driver.db = db
	//测试连接
	err = db.DB().Ping()
	if err != nil {
		log.Panic("mysql connect failed")
	}
	//是否存在表
	temp := &models.Proxy{}
	if !db.HasTable(temp) {
		db.CreateTable(temp)
	}
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(150)
	return
}
