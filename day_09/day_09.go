package main

import (
	"fmt"
	"sync"
	"time"
)

// 死鎖範例
// func main() {
// 	ch := make(chan int)
// 	ch <- 798
// 	num := <-ch
// 	fmt.Println("num", num)
// }

// 互斥鎖
// 使用channel完成同步
// var ch = make(chan int)

// func printer(str string) {
// 	for _, ch := range str {
// 		fmt.Printf("%c", ch)
// 		time.Sleep(time.Millisecond * 1000)

// 	}
// }
// func person1() {
// 	printer("hello")
// 	ch <- 98
// }

// func person2() {
// 	<-ch
// 	printer("hello")

// }

// func main() {
// 	go person1()
// 	go person2()

// }
var mutex sync.Mutex // 新建的互斥锁状态为0

func printer(str string) {
	mutex.Lock()
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	fmt.Println()
	mutex.Unlock()
}

func person1() {
	printer("hello")
}

func person2() {
	printer("world")
}

func main() {
	var wg sync.WaitGroup // 使用WaitGroup来等待所有goroutine完成

	wg.Add(2) // 设置需要等待的goroutine数量为2

	go func() {
		person1()
		wg.Done() // 当person1完成后，减少一个等待的goroutine
	}()

	go func() {
		person2()
		wg.Done() // 当person2完成后，减少一个等待的goroutine
	}()

	wg.Wait() // 等待所有goroutine完成
}
