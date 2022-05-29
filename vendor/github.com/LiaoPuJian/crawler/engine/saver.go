package engine

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

var (
	host   = "http://127.0.0.1:9200/"
	client *elastic.Client
)

func init() {
	//建立elasticSearch的链接
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(fmt.Sprintf("connect elasticSearch search error:【%s】", err))
	}
}

func ItemSaver(index string) chan Item {
	itemChan := make(chan Item)
	go func() {
		for {
			item := <-itemChan
			result, err := Save(item, index)
			if err != nil {
				log.Printf("save item : %v error. %v", item, err)
			}
			log.Printf("save success! item:%v, result:%v", item, result)
		}
	}()
	return itemChan
}

//这个方法将item保存到elasticSearch中
func Save(item Item, index string) (string, error) {
	if item.Type == "" {
		return "", errors.New("Type can not be empty!")
	}
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	put, err := indexService.Do(context.Background())
	if err != nil {
		return "", err
	}
	return put.Id, nil
}

//查找
func Gets(id string) (*elastic.GetResult, error) {
	//通过id查找
	return client.Get().Index("immoc").Type("github.com/LiaoPuJian/crawler").Id(id).Do(context.Background())
}
