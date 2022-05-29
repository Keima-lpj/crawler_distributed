package worker

import (
	"errors"
	"fmt"
	"github.com/LiaoPuJian/crawler/engine"
	"github.com/LiaoPuJian/crawler/zhenai/parser"
	"github.com/LiaoPuJian/crawler_distributed/config"
	"log"
)

//将远端的Parser服务定义到本地，并且写成可以在网络上传输的格式
type SerializedParser struct {
	Name string
	Args interface{}
}

//为了在网上可以通过RPC进行传递，需要我们自己定义自己的Request和ParserResult
type Request struct {
	Url    string
	Parser SerializedParser
}

type ParserResult struct {
	Items   []engine.Item
	Request []Request
}

//同时，这里要将自己的Request和engine的Request进行转换
func SerializeRequest(r engine.Request) Request {
	//这里调用engine.Request的序列化方法，获取到engine.Request中的函数名字和函数参数
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

//ParserResult同理
func SerializeParserResult(r engine.ParserResult) ParserResult {
	result := ParserResult{}
	//由于Item可以直接在网络上传输，所以直接拿过来
	result.Items = r.Item
	//由于engine.Request不能直接拿过来，所以需要调用上面的方法转一次
	for _, v := range r.Requests {
		result.Request = append(result.Request, SerializeRequest(v))
	}
	return result
}

//将自己的SerializedParser转换为engine的Parser
//这里这样做不一定好
func deserializeParser(sp SerializedParser) (engine.Parser, error) {
	switch sp.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		if args, ok := sp.Args.(string); ok {
			return parser.NewProfileParser(args), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", sp.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}

//将自己的Request转换为engine.Request
func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}

//将自己的ParserResult转换为engine.ParserResult
func DeserializeResult(p ParserResult) engine.ParserResult {
	result := engine.ParserResult{}
	result.Item = p.Items
	for _, v := range p.Request {
		engineReq, err := DeserializeRequest(v)
		if err != nil {
			log.Printf("error deserializing request %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}
