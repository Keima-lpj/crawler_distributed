package main

import (
	"crawler_distributed/config"
	"crawler_distributed/rpc_support"
	"crawler_distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpc_support.ServeRpc(fmt.Sprintf(":%v", config.WORKER_PORT0), worker.CrawlService{}))
}
