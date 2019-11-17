package main

import (
	"crawler_distributed/config"
	"crawler_distributed/persist"
	"crawler_distributed/rpc_support"
	"fmt"
	"log"
)

func main() {
	//启动rpc服务
	err := ServeRpc(config.ITEM_SAVE_HOST, config.ES_INDEX)
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
