package main

import (
	"flag"
	"github.com/p4vlowVl4d/notify-service/web"
	config "github.com/spf13/viper"
	"math/rand"
	"time"
)

var (
	CONFIG_FILE string = "config/config.yml"	//путь к файлу конфигурации
)

func Init() {
	flag.StringVar(&CONFIG_FILE, "c", CONFIG_FILE, "путь к файлу конфигурации")
}

func main() {
	Init()
	flag.Parse()
	config.SetConfigFile(CONFIG_FILE)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	webConf := createWebConfig()
	web.Execute(webConf)
}

func createWebConfig() *web.WebConf {
	readTime, writeTimeout := config.GetInt32("server.readTimeout"), config.GetInt32("server.writeTimeout")
	if readTime <= int32(0) || writeTimeout <= int32(0) {
		panic("Please set timeout option in config file!")
	}
	port, host := config.GetString("server.port"), config.GetString("server.host")
	if port == "" || host == "" {
		panic("Please set port option in config file!")
	}
	readTime = rand.Int31n(readTime)
	writeTimeout = rand.Int31n(writeTimeout)
	return &web.WebConf{
		ReadTimeout:  time.Duration(readTime),
		WriteTimeout: time.Duration(writeTimeout),
		PortListen:   port,
		Host:         host,
	}
}
