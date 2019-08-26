package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "Write with os.Stdout at %v", time.Now())
}
