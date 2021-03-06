package main

import (
	"fmt"
	"sync"
	"time"
)

/*
在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。
请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。
用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、
验证相关的token、请求的截止时间。 当一个请求被取消或超时时，所有用来处理该请求的 goroutine
都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。
*/

// 基本示例

var wg sync.WaitGroup

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	// 如何接收外部命令实现退出
	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()
	// 如何优雅的实现结束子goroutine
	wg.Wait()
	fmt.Println("over")
}
