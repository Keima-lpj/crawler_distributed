package main

import (
	"flag"
	"fmt"
	"github.com/LiaoPuJian/crawler_distributed/rpc_support"
	"github.com/LiaoPuJian/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "worker server port")

func main() {
	flag.Parse()
	if *port == 0 {
		panic("worker server port can not be 0")
	}
	log.Fatal(rpc_support.ServeRpc(fmt.Sprintf(":%v", *port), worker.CrawlService{}))
}
