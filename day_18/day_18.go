package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串按指定分隔符
	str := "I love my work and I I I I love my family too"
	// ret := strings.Split(str, "I")
	// for _, v := range ret {
	// 	fmt.Println(v)
	// }
	// 字符串按空格拆分
	res := strings.Fields(str)
	for _, s := range res {
		fmt.Println(s)
	}
	// 判斷字符串結束標記
	flg := strings.HasSuffix("test.abc", "mp2")
	fmt.Println(flg)
	//判斷字符串起始標記
	flgs := strings.HasSuffix("test.abc", "abc")
	fmt.Println(flgs)
}
