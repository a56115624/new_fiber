package main

import (
	"fmt"
	"time"
)

// 全局定義channel
var ch = make(chan int)

func printer(s string) {
	for _, ch := range s {
		fmt.Println(ch) // 頻幕 stdout
		time.Sleep(100 * time.Millisecond)
	}
}

// 定義兩個人使用打印機
func person1() { // person1 先執行

	printer("hello")
	ch <- 12
}
func person2() { // person2 後執行
	<-ch
	printer("world")
}
func main() {
	go person1()
	go person2()
	for {
	}
}
