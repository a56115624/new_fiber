package day2

import (
	"fmt"
	"strconv"
)

// SortNumber 编写函数，返回其（两个）参数正确的（自然）数字顺序：
func SortNumber(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a

}

// FoundError 下面的程序有什么错误？
func FoundError() {

	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	// fmt.Printf("%v\n", i) //作用域错误。这里会提示I未定义
}

// 定义栈结构
type stack struct {
	node int     //当前的节点
	data [10]int //实际数据
}

// 往栈中增加数据
func (stack *stack) Push(number int) {
	if stack.node < len(stack.data) {
		stack.data[stack.node] = number
		stack.node++
	}

}

// 从栈中取出数据
func (stack *stack) Pop() int {
	if stack.node > -1 {
		stack.node--

	}
	return stack.data[stack.node]
}

// 更进一步。编写一个 String 方法将栈转化为字符串形式的表达。可以这样的
// 方式打印整个栈： fmt.Printf("My stack %v\n", stack)
// 栈可以被输出成这样的形式： [0:m] [1:l] [2:k]
func (stack *stack) ToString() string {
	var s string
	for i := 0; i < stack.node; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(stack.data[i])
		s += "[" + k + ":" + v + "] "

	}

	return s
}

// Tooargs 编写函数接受整数类型变参，并且每行打印一个数字。
func Tooargs(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}

}

// Gblq 斐波那契数列以：1,1,2,3,5,8,13,... 开始
func Gblq(number int) []int {
	s := make([]int, number)
	for i := 1; i <= number; i++ {
		if i < 3 {
			s[i-1] = 1
		} else {
			s[i-1] = s[i-2] + s[i-3]

		}

	}
	return s

}

// ArrayMap 函数 map() 函数是一个接受一个函数和一个列表作为参数的函数。函
// 数应用于列表中的每个元素，而一个新的包含有计算结果的列表被返回
func ArrayMap(numbers []int, f func(int) int) []int {
	for i, v := range numbers {
		numbers[i] = f(v)
	}

	return numbers
}

// MaxOf 编写一个函数，找到 int slice ( []int ) 中的最大值。
func MaxOf(numbers []int) int {
	var max = 0
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max

}

// MinOf 编写一个函数，找到 int slice ( []int ) 中的最小值。
func MinOf(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var min = numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min

}

// ArrSort 冒泡排序
func ArrSort(numbers []int) {
	var temp int
	len := len(numbers)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if numbers[i] > numbers[j] {
				temp = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
				//或者
				//	numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	//	return numbers 因为是引用类型。这里不需要返回
}

// PlusTwo 编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2。函数的名称叫做 plusTwo
func PlusTwo() func(int) int {
	v := func(number int) int {
		return number + 2
	}
	return v
}

// PlusX 使 PlusTwo 中的函数更加通用化，创建一个 plusX(y) 函数，返回一个函数用于对整数加上 x 。
func PlusX(x int) func(int) int {
	v := func(y int) int {
		return x + y
	}
	return v

}
