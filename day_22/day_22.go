package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// func main() {
// 	// 獲取用戶輸入的路徑
// 	fmt.Println("請輸入帶查詢的目錄")
// 	var path string
// 	fmt.Scan(&path)

// 	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
// 	if err != nil {
// 		fmt.Println("OpenFile error: ", err)
// 		return
// 	}
// 	defer f.Close()
// 	// 讀取目錄項
// 	info, err := f.ReadDir(-1) // 讀取目錄中所有的目錄項

//		// 變量返回的切片
//		for _, filedInfo := range info {
//			if !filedInfo.IsDir() {
//				if strings.HasSuffix(filedInfo.Name(), "jpg") {
//					fmt.Println("jpg文件有", filedInfo.Name())
//				}
//			}
//		}
//	}

// 拷貝mp3的見到指定目錄
func cpMp3Dir(src, dst string) {
	// 打開讀文件
	f_r, err := os.Open(src)
	if err != nil {
		fmt.Println("Open err", err)
		return
	}
	defer f_r.Close()

	// 創建寫文件
	f_w, err := os.Create(dst)
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

func main() {
	// 獲取用戶輸入的路徑
	fmt.Println("請輸入帶查詢的目錄")
	var path string
	fmt.Scan(&path)

	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile error: ", err)
		return
	}
	defer f.Close()
	// 讀取目錄項
	info, err := f.ReadDir(-1) // 讀取目錄中所有的目錄項

	// 變量返回的切片
	for _, filedInfo := range info {
		if !filedInfo.IsDir() {
			if strings.HasSuffix(filedInfo.Name(), "jpg") {
				fmt.Println("jpg文件有", filedInfo.Name())
				cpMp3Dir(path+"/"+filedInfo.Name(), "./"+filedInfo.Name())
			}
		}
	}
}
