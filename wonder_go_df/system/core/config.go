package core

import (
	"log"

	"gopkg.in/ini.v1"
)

/**
 * 框架的配置文件相关操作
 */

var Database_file = ""

type Config struct {
	McHostname      string `ini:"mongodb_hostname"`
	McPort          string `ini:"mongodb_port"`
	MongodbAddr     string `ini:"mongodb_addr"`
	MongodbName     string `ini:"mongodb_dbname"`
	Port            string `ini:"port"`
	ExcelExportPath string `ini:"excel_export_path"`
}

var Conf Config = Config{}

//初始化
func Init(file string) {
	conf, err := ReadConfig(file)
	if err != nil {
		log.Fatalln("initialization config file fail")
	}

	Conf = conf
}

//读取配置文件并转成结构体
func ReadConfig(path string) (Config, error) {
	var config Config
	conf, err := ini.Load(path) //加载配置文件
	if err != nil {
		log.Println("load config file fail!", err)
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config) //解析成结构体
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}
