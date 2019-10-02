package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	// stream := io.MultiReader(computer, system, programming)
	// ここにioパッケージを使ったコードを書く
	io.Copy(os.Stdout, stream)
}
