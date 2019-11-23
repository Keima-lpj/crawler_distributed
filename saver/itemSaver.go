package saver

import (
	"crawler/engine"
	"crawler_distributed/config"
	"crawler_distributed/rpc_support"
	"fmt"
	"log"
)

//这里使用rpc来实现存储数据
func ItemSaver(host string) chan engine.Item {
	itemChan := make(chan engine.Item)

	//这里调用rpc
	client, err := rpc_support.NewClient(host)
	if err != nil {
		panic(fmt.Sprintf("connect rpc error:%v", err))
	}

	go func() {
		for {
			item := <-itemChan
			result := ""
			err = client.Call(config.ITEM_SAVE_SERVICE, item, &result)
			if err != nil {
				log.Printf("save item error.%v  %v", item, err)
			} else {
				log.Printf("save success! item:%v, result:%v", item, result)
			}
		}
	}()
	return itemChan
}
