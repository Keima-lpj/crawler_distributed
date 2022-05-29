package parser

import (
	"github.com/LiaoPuJian/crawler/engine"
	"regexp"
)

//正则表达式
//注：[^>]表示匹配非'>'符号的任意字符串
const cityListRegexpString = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" [^>]+>([^>]+)</a>`

//这个是城市列表的解析器，通过解析城市列表页面的文本，返回下一级页面的request数组和对应的item
func ParseCityList(contents []byte, url string) engine.ParserResult {
	//使用正则表达式来匹配到对应的城市和链接
	re := regexp.MustCompile(cityListRegexpString)
	match := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	//默认先只取前10个城市
	match = match[0:9]

	for _, v := range match {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(v[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParserCity"), //传入城市解析器
		})
	}

	return result
}
