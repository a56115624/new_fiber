package main

import (
	"fmt"
)

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

// func main() {
// 	c := make(chan int, 10)
// 	go fibonacci(cap(c), c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// switch case
// func main() {
// 	var weekday uint8
// 	fmt.Println("請輸入星期（）")
// 	fmt.Scanln(&weekday)
// 	switch weekday {
// 	case 1:
// 		fmt.Println("醬油炒飯")
// 		fallthrough
// 	case 2:
// 		fmt.Println("醬油炒麵")
// 	default:
// 		fmt.Println("輸入有誤")
// 	}
// }

func main() {
	fmt.Println("無限循環")
	i := 1
	for {
		fmt.Print(i, "\t")
		i++
		if i == 10 {
			fmt.Println()
			break
		}
	}
	fmt.Println("條件循環")
	i = 1
	for i < 11 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println("標準循環")
	for j := 1; j < 11; j++ {
		fmt.Print(i, "\t")
	}
}
