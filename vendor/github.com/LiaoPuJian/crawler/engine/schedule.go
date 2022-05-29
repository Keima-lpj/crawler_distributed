package engine

type Schedule interface {
	Submit(request Request)
	WorkerReady(w chan Request)
	Run()
}

//队列调度schedule
type QueueScheduler struct {
	RequestChan chan Request
	WorkerChan  chan chan Request
}

func (q *QueueScheduler) Submit(r Request) {
	q.RequestChan <- r
}

//这个方法用于告诉schedule，有一个worker就绪了
func (q *QueueScheduler) WorkerReady(w chan Request) {
	q.WorkerChan <- w
}

//启动整个scheduler
func (q *QueueScheduler) Run() {
	q.RequestChan = make(chan Request)
	q.WorkerChan = make(chan chan Request)

	go func() {
		//声明两个队列
		var requestQ []Request
		var workerQ []chan Request

		//开启循环调度器。这个循环调度器的作用是将request和worker分别从队列中取出
		//将取出的request放入取出的worker中
		for {
			//声明本次循环中活动的request和worker
			var activeRequest Request
			var activeWorker chan Request
			//如果两个队列中都有值，则从两个队列中获取第一个值
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-q.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-q.WorkerChan:
				workerQ = append(workerQ, w)
			//这里实现调度，将获取到的request传给类型威chanRequest的worker，并将对应的worker和request中队列中踢出
			//这里同样运用了nil channel的特性，如果activeWorker并未从worker中获取到值，则其为nil，则select不会调度到nil channel上
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}

		}
	}()
}
