package engine

import (
	"github.com/LiaoPuJian/crawler/fetcher"
	"fmt"
)

func Work(r Request) (ParserResult, error) {
	//通过fetch获取到了utf8格式的html源代码
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		fmt.Printf("Got url body error : %s", err)
		return ParserResult{}, err
	}
	fmt.Println(r)
	//再通过解析器，来解析源代码，得到下一级的[]requests
	return r.Parser.Parser(body, r.Url), nil
}
