package main

import (
	"fmt"
	"io"
	"os"
)

// func main() {
// 	f, err := os.Create("/Users/shane/Desktop/Golang_Project/new_fiber/testFile.csv")
// 	if err != nil {
// 		fmt.Println("create err", err)
// 	}
// 	defer f.Close()
// 	fmt.Println("successful")
// }

// func main() {
// 	f, err := os.OpenFile("/Users/shane/Desktop/Golang_Project/new_fiber/testFile.csv", os.O_RDWR|os.O_CREATE, 0644)
// 	if err != nil {
// 		fmt.Println("open err", err)
// 		return
// 	}
// 	defer f.Close()
// 	_, err = f.WriteString("#########")
// 	if err != nil {
// 		fmt.Println("write err", err)
// 	}
// 	fmt.Println("successful")
// }

func main() {
	f, err := os.OpenFile("/Users/shane/Desktop/Golang_Project/new_fiber/testFile.csv", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err", err)
		return
	}
	defer f.Close()
	fmt.Println("successfully")
	n, err := f.WriteString("helloworld\r\n")
	if err != nil {
		fmt.Println("WriteString err", err)
		return
	}
	fmt.Println("writeString", n)

	off, err := f.Seek(-5, io.SeekEnd)
	if err != nil {
		fmt.Println("Seek err", err)
		return
	}
	fmt.Println("off", off)

	n, _ = f.WriteAt([]byte("11111"), off)
	fmt.Println("writeAt", n)
}
