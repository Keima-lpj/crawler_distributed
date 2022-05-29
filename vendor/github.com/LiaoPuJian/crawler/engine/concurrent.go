package engine

import (
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Schedule         Schedule
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(r Request) (ParserResult, error)

func (ce ConcurrentEngine) Run(seeds ...Request) {
	//启动schedule，内部开启goroutine进行调度
	ce.Schedule.Run()

	//将seeds提交到scheduler中
	for _, r := range seeds {
		ce.Schedule.Submit(r)
	}

	out := make(chan ParserResult)

	//创建对应数量的worker
	for i := 1; i <= ce.WorkerCount; i++ {
		ce.CreateWorker(out, i)
	}

	for {
		result := <-out
		//这里将结果的request提交的scheduler中
		for _, r := range result.Requests {
			ce.Schedule.Submit(r)
		}
		for _, v := range result.Item {
			//这里将获取到的item放入itemchan中
			go func() {
				ce.ItemChan <- v
			}()
		}
	}

}

/**
创建worker
*/
func (ce ConcurrentEngine) CreateWorker(out chan ParserResult, i int) {
	in := make(chan Request)
	go func(i int) {
		fmt.Printf("Worker #%d Start!\n", i)
		for {
			//告诉scheduler，我这个worker已经准备好了
			ce.Schedule.WorkerReady(in)
			//从in中获取数据，进行消费
			r := <-in
			//调用worker方法
			result, err := ce.RequestProcessor(r)
			if err != nil {
				log.Printf("get work error:%s, url:%s", err, r.Url)
			}
			//将结果写入到out channel中
			out <- result
		}
	}(i)
}
