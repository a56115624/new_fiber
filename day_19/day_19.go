package main

import (
	"fmt"
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

func main() {
	f, err := os.OpenFile("/Users/shane/Desktop/Golang_Project/new_fiber/testFile.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open err", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString("#########")
	if err != nil {
		fmt.Println("write err", err)
	}
	fmt.Println("successful")
}
