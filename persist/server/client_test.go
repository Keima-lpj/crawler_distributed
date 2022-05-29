package main

import (
	"github.com/LiaoPuJian/crawler/engine"
	"github.com/LiaoPuJian/crawler_distributed/rpc_support"
	"testing"
	"time"
)

func TestItemSaverService(t *testing.T) {
	//启动rpc server
	go ServeRpc(":1234", "test1")
	time.Sleep(time.Second)
	//启动client调用rpc里面的服务
	client, err := rpc_support.NewClient(":1234")
	if err != nil {
		t.Errorf("rpc客户端链接失败:%s", err)
	}

	//调用save方法保存对象
	Item := engine.Item{
		Id:      "1",
		Url:     "www.baidu.com",
		Type:    "zhenai",
		Payload: "rua!",
	}
	var result string
	err = client.Call("ItemSaverService.Save", Item, &result)
	if err != nil {
		t.Errorf("调用ItemSaverService.Save失败:%s", err)
	}

}
