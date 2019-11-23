package client

import (
	"crawler/engine"
	"crawler_distributed/config"
	"crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor, error) {
	//返回一个engine.Processor格式的函数
	return func(r engine.Request) (result engine.ParserResult, e error) {
		sReq := worker.SerializeRequest(r)
		sResult := worker.ParserResult{}

		c := <-clientChan
		err := c.Call(config.WORKER_SERVICE, sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}
		//将sResult转换为engine.ParserResult
		return worker.DeserializeResult(sResult), nil
	}, nil
}
