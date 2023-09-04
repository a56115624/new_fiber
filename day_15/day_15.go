package main

import (
	"fmt"
	"strings"
)

// map 做函數傳參數,返回值,傳引用
// func mapDelete(m map[int]string, key int) map[int]string {
// 	delete(m, key)
// 	return m
// }

//	func main() {
//		m := map[int]string{1: "luffy,", 2: "ok", 250: "Zoro", 543: "88"}
//		m2 := mapDelete(m, 1)
//		fmt.Println(m2)
//	}
func main() {
	str := "I love my work and I I I I love my family too"
	m2 := checkStr(str)
	for k, v := range m2 {
		fmt.Printf("%q: %d\n", k, v) // 使用 Printf 函数来格式化输出
	}
}

func checkStr(data string) map[string]int {
	s := strings.Fields(data) // 將字符串,拆分成字符串切片
	m := make(map[string]int) // 創建一個用於存儲word 出現次數的map
	for i := 1; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}
	return m
}
