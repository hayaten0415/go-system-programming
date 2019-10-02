package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	b := make([]byte, 1024)
	rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	_, err2 := newFile.Write(b)
	if err2 != nil {
		fmt.Println("file write err:", err2)
		return
	}

	fmt.Println("file write ok.")

}
