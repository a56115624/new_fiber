package main

import (
	"fmt"
	"new_fiber/homework"
	"new_fiber/util"
)

var A = util.F("main A")

func init() {
	util.F("main.init1")
}
func init() {
	util.F("main.init2")
}

func main() {
	homework.DeferRecover()
	fmt.Println("我還在運行")
}

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	// 创建一个 Fiber 应用程序
// 	app := fiber.New()

// 	// 定义路由和处理函数
// 	app.Get("/", helloHandler)

// 	// 启动服务器并监听指定的地址和端口
// 	err := app.Listen(":8080")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // 处理函数
// func helloHandler(c *fiber.Ctx) error {
// 	return c.SendString("Hello, World!")
// }
// import (
// 	"fmt"
// 	"new_fiber/homework"
// )

// func main() {
// 	homework.DeferRecover()
// 	fmt.Println("我還在運行")
// }

// 4.1 數組
func Array() {
	var a [3]int = [3]int{1, 456, 789}
	a[0] = 123
}
