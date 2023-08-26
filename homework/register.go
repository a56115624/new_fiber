package homework

import (
	"fmt"
	"new_fiber/util"
)

var A = util.F("register A")

func init() {
	util.F("register.init1")
}
func init() {
	util.F("register.init2")
}

// func Register() {
// 	for {
// 		var username, pwd, cpwd string
// 		fmt.Println("歡迎註冊")
// 		fmt.Println("用戶名")
// 		fmt.Scanf("%s\n", &username)
// 		fmt.Println("密碼")
// 		fmt.Scanf("%s\n", &pwd)
// 		fmt.Println("確認密碼")
// 		fmt.Scanf("%s\n", &cpwd)
// 		if username == "" || pwd == "" || cpwd == "" {
// 			fmt.Println("輸入不得為空")
// 			continue
// 		}
// 		if pwd != cpwd {
// 			fmt.Println("輸入密碼不一樣")
// 			continue
// 		}
// 		fmt.Println("註冊成功")
// 		break
// 	}
// }

//	func Register() {
//		for {
//			var username, pwd, cpwd string
//			fmt.Println("請輸入用戶名")
//			fmt.Scanf("%s", &username)
//			fmt.Println("請輸入用密碼")
//			fmt.Scanf("%s", &pwd)
//			fmt.Println("請輸入確認密碼")
//			fmt.Scanf("%s", &cpwd)
//			if username == "" || pwd == "" || cpwd == "" {
//				fmt.Println("輸入不得為空")
//				continue
//			}
//			if pwd != cpwd {
//				fmt.Println("兩次帳密不相同")
//				continue
//			}
//			fmt.Println("註冊成功")
//			break
//		}
//	}

// 3.4 lable and goto
// func LableAndGoto() {
// outside:
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			fmt.Print("+")
// 			if i == 9 && j == 4 {
// 				break outside
// 			}
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println("goto")
// 	fmt.Print('1')
// 	if i := 1; i != 1 {

// 		fmt.Print('2')
// 		fmt.Print('3')
// 	}
// 	fmt.Print('4')
// 	fmt.Print('4')
// }

// 函數

func Function() {
	res1, res2 := func(n1 int, n2 int) (sum, difference int) {
		sum = n1 + n2
		difference = n1 - n2
		return
	}(2, 3)
	fmt.Println("res1", res1, "res2", res2)

}

// defer
func deferUtil() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次調用接收到n = %v\n", n)
		i++
		fmt.Printf("工具被第%v次調用\n", i)
		return i
	}

}
func Defer() int {
	f := deferUtil()
	result := f(3)
	defer f(result)
	return f(2)
}

func DeferRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

// init函數
