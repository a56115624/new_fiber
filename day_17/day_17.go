package main

import "fmt"

type Person2 struct {
	name string
	sex  byte
	age  int
}

func test2(p *Person2) {
	p.name = "luffy"
	p.sex = 'm'
	p.age = 10
}

func main() {
	var p1 *Person2 = &Person2{name: "n1", sex: 'f', age: 13}
	fmt.Println("p1", p1)
	var tmp Person2 = Person2{name: "n2", sex: 'm', age: 19}
	var p2 *Person2
	p2 = &tmp
	fmt.Println("p2", p2)
	p3 := new(Person2)
	p3.name = "n3"
	p3.sex = 'f'
	p3.age = 22
	fmt.Println("p3", p3)
	fmt.Printf("&p3.name = %p\n", &p3.name)
	test2(p3)
	fmt.Println("p3", p3)
}
