package main

import (
	"crawler_distributed/config"
	"crawler_distributed/persist"
	"crawler_distributed/rpc_support"
	"flag"
	"fmt"
	"log"
)

var port = flag.String("port", "", "item saver port")

func main() {
	flag.Parse()
	if *port == "" {
		panic("port can not be empty string")
	}
	//启动rpc服务
	err := ServeRpc(*port, config.ES_INDEX)
	if err != nil {
		panic(fmt.Sprintf("rpc service start error! %s", err))
	}
	log.Printf("rpc service start!")
}

func ServeRpc(host, index string) error {
	return rpc_support.ServeRpc(host, &persist.ItemSaverService{
		Index: index,
	})
}
