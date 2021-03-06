package main

import (
	"fmt"
	"time"
)

/*
· worker pool（goroutine池）
	在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，
控制goroutine的数量，防止goroutine泄漏和暴涨。
*/

/* 一个简易的work pool示例代码如下： */

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		i := <-results
		fmt.Println(i)
	}
}
