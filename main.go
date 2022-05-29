package main

import (
	"flag"
	"fmt"
	"github.com/LiaoPuJian/crawler/engine"
	"github.com/LiaoPuJian/crawler/zhenai/parser"
	"github.com/LiaoPuJian/crawler_distributed/config"
	"github.com/LiaoPuJian/crawler_distributed/rpc_support"
	"github.com/LiaoPuJian/crawler_distributed/saver"
	"github.com/LiaoPuJian/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemsaver_host = flag.String("itemsavor_host", "", "用于itemsavor服务的host")

	worker_hosts = flag.String("worker_host", "", "用于连接多个worker，中间用逗号分隔")
)

func main() {

	flag.Parse()

	pool := createClientPool(strings.Split(*worker_hosts, ","))
	//新建一个worker的存储方法，单机版爬虫直接为Work函数，分布式这里生成调用RPC的服务
	processor, err := client.CreateProcessor(pool)
	if err != nil {
		fmt.Println("create Processor fail, ", err)
		return
	}

	engine.ConcurrentEngine{
		Schedule:         &engine.QueueScheduler{},
		WorkerCount:      10,
		ItemChan:         saver.ItemSaver(*itemsaver_host),
		RequestProcessor: processor,
	}.Run(engine.Request{
		Url:    config.URL,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}

func createClientPool(host []string) chan *rpc.Client {

	var clients []*rpc.Client
	for _, h := range host {
		client, err := rpc_support.NewClient(h)
		if err != nil {
			log.Printf("process connect %s error", h)
		} else {
			log.Printf("process connect %s success", h)
			clients = append(clients, client)
		}
	}

	clientChan := make(chan *rpc.Client)

	//这里起一个goroutine不停的向clientChan中发送client
	go func() {
		for {
			for _, c := range clients {
				clientChan <- c
			}
		}
	}()
	return clientChan

}
