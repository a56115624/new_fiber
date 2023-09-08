package main

import (
	"fmt"
	"time"
)

// func sing() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("正在唱,隔壁泰山")
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func dance() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("---正在跳舞")
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func main() {
// 	go sing()
// 	go dance()
// 	for {

//		}
//	}
func main() {
	go func() { // 創建一個子go 程
		for i := 0; i < 10; i++ {
			fmt.Println("I,m goroutine")
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Printf("I'm main\n")
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
