package client

import (
	"crawler/engine"
	"crawler_distributed/config"
	"crawler_distributed/rpc_support"
	"crawler_distributed/worker"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpc_support.NewClient(fmt.Sprintf(":%v", config.WORKER_PORT0))
	if err != nil {
		return nil, err
	}

	//返回一个engine.Processor格式的函数
	return func(r engine.Request) (result engine.ParserResult, e error) {
		sReq := worker.SerializeRequest(r)
		sResult := worker.ParserResult{}

		err := client.Call(config.WORKER_SERVICE, sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}

		//将sResult转换为engine.ParserResult
		return worker.DeserializeResult(sResult), nil
	}, nil
}
