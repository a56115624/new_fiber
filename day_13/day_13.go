package main

import "fmt"

// func test(m int) {
// 	var b int = 1000
// 	b += m
// }

//	func main() {
//		var a int = 10
//		var p *int = &a
//		a = 100
//		fmt.Println("a", a)
//		test(12)
//		*p = 250 // 借助a變量的地,去操控a空間
//		fmt.Println("a", a)
//		fmt.Println("*P", *p)
//		a = 1000
//		fmt.Println("*P", *p)
//	}
// func test() {
// 	p := new(int)
// 	*p = 1000
// 	fmt.Println(*p)
// }

// func main() {
// 	var p *string
// 	// 在 heap 上申请一片内存地址空间
// 	p = new(string)        // 默认类型的默认值
// 	fmt.Printf("%s\n", *p) // 使用格式化占位符 %s
// 	fmt.Printf("%q\n", *p) // 使用格式化占位符 %q
// 	test()
// }

// var a int = 10
// var b int  = 20
// a = b // 取出b對應的內容空間放入a內

func swap(a, b int) {
	a, b = b, a
	fmt.Println("swap a", a, "b", b)
}
func swap2(x, y *int) {
	*x, *y = *y, *x // * 在等號左邊,取空間,在等號右邊取值
}
func main() {
	a, b := 10, 20
	swap(a, b)    // 傳變量值
	swap2(&a, &b) // 傳地址值
	fmt.Println("main a", a, "b", b)
	fmt.Println("swap2 a", a, "b", b)
}
