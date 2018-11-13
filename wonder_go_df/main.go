package main

/**
 * 执行命令
 * go run main.go -config application/config/config.conf
 */

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"wonder_go_df/application/config"
	"wonder_go_df/system/core"
)

var configPath = flag.String("config", "", "config path")

func main() {
	fmt.Println("wonderGoDF start")

	flag.Parse()
	if *configPath == "" {
		log.Fatalln("config path not allow empty")
	}
	core.Init(*configPath)

	// 连接 Mongodb
	// database.ConnectMongoDB(core.Conf.MongodbAddr)

	router := core.NewRouter(config.RoutesVal)
	host := ":" + core.Conf.Port
	log.Println("start server at", host)
	log.Fatal(http.ListenAndServe(host, router))
}
