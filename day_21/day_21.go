package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打開讀文件
	f_r, err := os.Open("/Users/shane/Desktop/Golang_Project/new_fiber/testFile.csv")
	if err != nil {
		fmt.Println("Open err", err)
		return
	}
	defer f_r.Close()

	// 創建寫文件
	f_w, err := os.Create("/Users/shane/Desktop/Golang_Project/new_fiber/test2.txt")
	if err != nil {
		fmt.Println("Open err", err)
		return
	}
	defer f_w.Close()

	// 讀文件中獲取數據,放到緩衝區中
	buf := make([]byte, 4096)

	// 循環從讀文件中,獲取數據,"原封不動的,寫到寫文件中"
	for {
		n, err := f_r.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("讀取錯誤：%v\n", err)
			return
		}
		if n == 0 {
			break // 已經讀取完畢，退出循環
		}
		// 寫入到寫文件中
		f_w.Write(buf[:n]) // 讀多少寫多少
	}
	fmt.Println("成功複製文件")
}
