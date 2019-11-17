package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
	"crawler_distributed/config"
	"crawler_distributed/saver"
	"crawler_distributed/worker/client"
	"fmt"
)

func main() {

	//新建一个worker的存储方法，单机版爬虫直接为Work函数，分布式这里生成调用RPC的服务
	processor, err := client.CreateProcessor()
	if err != nil {
		fmt.Println("create Processor fail, ", err)
		return
	}

	engine.ConcurrentEngine{
		Schedule:         &engine.QueueScheduler{},
		WorkerCount:      10,
		ItemChan:         saver.ItemSaver(),
		RequestProcessor: processor,
	}.Run(engine.Request{
		Url:    config.URL,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}
