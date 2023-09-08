package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 創建一個帶有緩衝區的(用戶緩衝)reader
func main() {

	f, err := os.OpenFile("/Users/shane/Desktop/Golang_Project/new_fiber/testFile.csv", os.O_RDWR, 0644)
	if err != nil && err == io.EOF {
		fmt.Println("OpenFile err", err)
		return
	}
	defer f.Close()
	fmt.Println("successfully")
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("ReadBytes err", err)
			return
		}
		fmt.Println(string(buf))
	}
}
