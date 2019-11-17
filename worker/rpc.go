package worker

import (
	"crawler/engine"
)

//这里是worker的RPC服务
type CrawlService struct{}

func (CrawlService) Process(req Request, result *ParserResult) error {
	//这里将调用服务传递的网络格式的Request转换为本地的engine.Request
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	//通过engine.Work方法解析这个engine.Request，得到对应的engine.ParserResult
	engineResult, err := engine.Work(engineReq)
	if err != nil {
		return err
	}
	//再将这个engine.ParserResult转换为网络格式的ParserResult，写入到result中
	*result = SerializeParserResult(engineResult)
	return nil
}
