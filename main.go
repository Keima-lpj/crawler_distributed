package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
	"crawler_distributed/config"
	"crawler_distributed/saver"
)

func main() {
	engine.ConcurrentEngine{
		Schedule:    &engine.QueueScheduler{},
		WorkerCount: 10,
		ItemChan:    saver.ItemSaver(),
	}.Run(engine.Request{
		Url:        config.URL,
		ParserFunc: parser.ParseCityList,
	})

}
