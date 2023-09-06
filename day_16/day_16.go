package main

import (
	"fmt"
	"unsafe"
)

type Persons struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	p     []Persons
	id    []int
	score []int
}

func test(man Persons) {
	man.name = "bod"
	man.sex = 1
	man.age = 33
	fmt.Println("test", man)
}

func main() {
	// var man = Persons{name: "John", sex: 0, age: 18} //順序初始化，必須全部都定義不含位噴錯
	// fmt.Println(man)
	// man2 := Persons{name: "羅成員", age: 30} // 部分初始化
	// fmt.Println(man2)
	// fmt.Printf("man2:%q", man2.name) //索引成員變量
	// var man3 Persons
	// man3.name = "小王"
	// man3.age = 18
	// man3.sex = 1

	// 結構體比較
	var p1 Persons = Persons{name: "shane", sex: 0, age: 33}
	var p2 Persons = Persons{name: "jack", sex: 0, age: 13}
	var p3 Persons = Persons{name: "shane", sex: 0, age: 33}
	fmt.Println(p1 == p2)
	fmt.Println(p1 == p3)
	// 在函數內部使用結構體傳參
	var temp Persons
	fmt.Println(unsafe.Sizeof(temp))
	test(temp) // 值傳遞,將實參的值拷貝一份給行參

	fmt.Printf("%p\n", &temp)
	fmt.Printf("%p\n", &temp.name)
	fmt.Println("main temp size", unsafe.Sizeof(temp))

}
