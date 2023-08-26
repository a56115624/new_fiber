package main

import (
	"fmt"
	"sync"
	"unicode/utf8"
)

// First  创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值。
func First() {
	fmt.Println("first func start")
	for i := 1; i < 100; i++ {
		fmt.Println(i)
	}
	fmt.Println("first func end")
}

// Second  用 goto 改写 1 的循环。关键字 for 不可使用。
func Second() {
	fmt.Println("second func start")
	i := 1
HERE:
	fmt.Println(i)
	if i < 10 {
		i++
		goto HERE
	}
	fmt.Println("second func end")
}

// Third  再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上。
func Third() {
	fmt.Println("third func start")
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, v := range arr {
		fmt.Printf("第%d元素為%d\n", i+1, v)
	}
	fmt.Println("third func end")
}

func main() {
	const start = 1
	const end = 100

	// 创建一个等待组，用于同步 Goroutine
	var wg sync.WaitGroup

	// 启动多个 Goroutine 进行并发处理
	for i := start; i <= end; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			result := fizzBuzz(num)
			fmt.Println(result)
		}(i)
	}

	// 等待所有 Goroutine 完成
	wg.Wait()
}

// FizzBuzz 函数，根据数字判断并返回对应的字符串
func fizzBuzz(num int) string {
	switch {
	case num%3 == 0 && num%5 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", num)
	}
}

// PrintA 建立一个 Go 程序打印下面的内容（到 100 个字符）：  A
// AA
// AAA
// AAAA
// AAAAA
// AAAAAA
// AAAAAAA
func PrintA(number int) {
	a := "A"
	for i := 0; i < number; i++ {
		fmt.Println(a)
		a = a + "A"
	}
}

// CountStr 建立一个程序统计字符串里的字符数量：
// asSASA ddd dsjkdsjs dk你好
// 同时输出这个字符串的字节数。
func CountStr() {
	str := "asSASA ddd dsjkdsjs dk你好"
	strRune := []rune(str)
	fmt.Printf("len长度为:%d\n", len(str))
	fmt.Printf("rune长度为:%d\n", len(strRune))
	fmt.Printf("RuneCountInString长度为:%d\n", utf8.RuneCountInString(str))
}

// ReverseStr  编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”
func ReverseStr(str string) {
	rs := []rune(str)
	var newStr []rune
	for i := len(rs) - 1; i >= 0; i-- {
		newStr = append(newStr, rs[i])
	}
	fmt.Println(newStr)
}

// AverageSlice 编写计算一个类型是 float64 的 slice 的平均值的代码。
func AverageSlice() {
	s := []float64{10, 20, 30, 40, 50, 60, 80, 100}
	var total float64
	for _, v := range s {
		total += v
	}
	fmt.Println(total / float64(len(s)))
}
