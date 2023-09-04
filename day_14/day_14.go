package main

import "fmt"

// func main() {
// 	arr := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	s := arr[1:5:7]
// 	fmt.Println("a", arr, "s", s)
// 	fmt.Println("s", cap(s)) // 5-1=4
// 	fmt.Println("s", len(s)) // 7-1=6
// 	s2 := s[0:6]
// 	fmt.Println("s2", s2)
// 	fmt.Println("s2", cap(s2)) // 6-0=6
// 	fmt.Println("s2", len(s2))
// 	s1 := s2[0:3]
// 	fmt.Println("s1", s1)
// 	fmt.Println("s1", cap(s1)) // 6-0=6
// }

// func main() {
// 	// 1. 自動推倒賦予初值
// 	s1 := []int{1, 2, 4, 6}
// 	fmt.Printf("s1 = %d", s1)

// 	s2 := make([]int, 5, 10)
// 	fmt.Println("len", len(s2), "cap", cap(s2))

//		s3 := make([]int, 7)
//		fmt.Println("len", len(s3), "cap", cap(s3))
//	}

// 使用append

// func noEmptyString(s []string) []string {
// 	out := s[:0]
// 	for _, str := range s {
// 		if str != "" {
// 			out = append(out, str)
// 		}
// 	}
// 	return out
// }

// func main() {
// 	a := []string{"red", "", "green", "pink", "green"}
// 	aferDate := noEmptyString(a)
// 	fmt.Println(aferDate)
// }

// 不使用append

// func main() {
// 	a := []string{"black", "red", "orange", "pink", "yellow", "yellow", "", "red", "green", ""}
// 	newData := sameString(a)
// 	fmt.Println(newData)
// }

// func sameString(a []string) []string {
// 	newa := a[:1]
// 	for _, word := range a {
// 		i := 0
// 		for ; i < len(newa); i++ {
// 			if word == newa[i] {
// 				break
// 			}
// 		}
// 		if i == len(newa) {
// 			newa = append(newa, word)
// 		}
// 	}
// 	return newa
// }

// func main() {
// 	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	s1 := data[:8]
// 	s2 := data[0:5]
// 	copy(s1, s2)
// 	fmt.Println(s1, s2)
// }

// 練習刪除slice中原有的元素,並且保存元素的順序
// {5,6,7,8,9}    -> {5,6,8,9}

// func main() {
// 	ask := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	ans := remove(ask, 2)
// 	fmt.Println(ans)
// }
// func remove(data []int, index int) []int {
// 	copy(data[index:], data[index+1:])
// 	data = data[:len(data)-1]
// 	return data
// }

// func main() {
// 	ask := []int{5, 6, 7, 8, 9}
// 	ans := takeout7(ask, 2)
// 	fmt.Println(ans)

// }
// func takeout7(data []int, index int) []int {
// 	copy(data[index:], data[index+1:])
// 	data = data[:len(data)-1]
// 	return data
// }

func main() {
	// var m1 map[int]string // map聲明他沒有空間不能存儲ｋｅｙ
	// // m1[100] = "hello"
	// if m1 == nil {
	// 	fmt.Println("我是空的")
	// }
	// m2 := map[int]string{}
	// fmt.Println(len(m2))
	// m2[100] = "好帥"
	// fmt.Println(m2)
	// m2[4] = "red"
	// fmt.Println(m2)

	// m3 := make(map[int]string)
	// m3[1] = "在幹嘛"
	// fmt.Println(m3)

	// m4 := make(map[int]string, 5)
	// fmt.Println(len(m4))
	// fmt.Println(m4)

	// 	var m5 map[int]string = map[int]string{1: "luffy,", 2: "ok", 250: "Zoro", 543: "88"}
	// 	fmt.Println(m5)
	// m6 := map[int]string{1: "luffy,", 2: "ok", 250: "Zoro", 34: "09"}
	// fmt.Println(m6)
	// m7 := make(map[int]string)
	// m7[100] = "ok"
	// m7[101] = "hello"
	// m7[101] = "good"
	// fmt.Println(m7)
	//遍歷
	var m5 map[int]string = map[int]string{1: "luffy,", 2: "ok", 250: "Zoro", 543: "88"}
	for _, v := range m5 {
		fmt.Printf("value: %q\n", v)
	}
	for k, v := range m5 {
		fmt.Printf("key: %d, value: %q\n", k, v)
	}
	if v, ok := m5[1]; ok {
		fmt.Println("vlaue", v, "ok", ok)
	}
}
