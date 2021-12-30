package main

import (
	"fmt"
	"sync"
	"time"
)

// cond 处理边界问题

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func main() {

	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	go func() {
		for {
			q.Enqueue("s")
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			fmt.Println("string: ", q.Dequeue())
		}
	}()

	fmt.Println("main ...")
	time.Sleep(time.Second * 3)
}

func (q *Queue) Enqueue(item string) {

	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Println("puting#{item} to queue, notify all ...")
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	// 判断当前的队列是否为空。解决当数据为空时，需要等待的问题。
	if len(q.queue) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	reslut := q.queue[0]
	q.queue = q.queue[1:]

	return reslut
}
