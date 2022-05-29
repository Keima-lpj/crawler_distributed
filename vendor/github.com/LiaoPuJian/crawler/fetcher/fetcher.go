package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//新建一个10毫秒触发一次的chan
var timeLimiter = time.Tick(10 * time.Millisecond)

//这里通过获取的url来返回对应解析成功的utf8格式的html文本
func Fetch(url string) ([]byte, error) {
	<-timeLimiter
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	//request.Header.Add("Cookie", "xxxxxx")
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	request.Header.Add("upgrade-insecure-requests", "1")
	request.Header.Add("Accept-Encoding", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	request.Header.Add("Accept-Language", "gzip, deflate")
	request.Header.Add("Cache-Control", "zh-CN,zh;q=0.9,en;q=0.8")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Cookie", "sid=efaad604-ed08-4203-9b22-5b9a2bdc05aa; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1571820651,1571822083,1571822087,1571822312; oneclickLoginSwitch=28635; oneClickRegisterSwitch=27615; __channelId=905821%2C0; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1572586373")
	request.Header.Add("Host", "www.zhenai.com")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//判断响应码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Http code error, %d\n", resp.StatusCode)
	}

	bufReader := bufio.NewReader(resp.Body)
	encode := determineEncoding(bufReader)

	//转换编码
	reader := transform.NewReader(bufReader, encode.NewDecoder())
	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("reader Html Error:%s\n", err)
	}

	return contents, nil
}

// 获取当前页面的编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		// 如果没有获取到编码格式，则返回默认UTF-8编码格式
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
