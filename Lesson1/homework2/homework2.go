// 基于Channel编写一个简单的单线程生产者消费者模型
// 队列: 队列长度10，队列元素类型为int
// 生产者: 每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
// 消费者: 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
		fmt.Printf("Produce int value: %d\n", i)
	}
}

func consumer(ch <-chan int) {
	for i := range ch {
		time.Sleep(time.Second * 2)
		fmt.Printf("Consume int value: %d\n", i)
	}
}

func main() {
	ch := make(chan int, 2)
	defer close(ch)
	go producer(ch)
	consumer(ch)
}
