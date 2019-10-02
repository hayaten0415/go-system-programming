package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("0")
	duration := 10 * time.Second
	fmt.Println("Waiting 10Seconds")
	<-time.After(duration)
	fmt.Println("10seconds")
}
