package main

import (
	"fmt"
	"runtime"
)

// func test() {
// 	defer fmt.Println("cccc")
// 	// return
// 	runtime.Goexit() // 退出當前go 程
// 	fmt.Println("ddddddd")
// }

func main() {
	n := runtime.GOMAXPROCS(2) // 將cpu設置承單核
	fmt.Println("n\n", n)
	for {
		go fmt.Println("0") // 子go 程
		fmt.Print("1")      // 主go 程
	}
	// go func() {
	// 	for {
	// 		fmt.Println("this is goroutine test")

	// 	}
	// }()
	// for {
	// 	runtime.Gosched()
	// 	fmt.Println("main")

	// }

	// go func() {
	// 	fmt.Println("aaaaaaa")
	// 	test()
	// 	fmt.Println("eeeeeee")
	// }()
	// for {

	// }
}
