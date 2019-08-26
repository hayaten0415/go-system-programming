package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	//書き出し先のテキストファイル生成
	file, err := os.Create("csv.txt")
	if err != nil {
		panic(err)
	}
	//読み込むcsvファイル
	readfile, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer readfile.Close()
	//https://golang.org/pkg/encoding/csv/#pkg-examples
	reader := csv.NewReader(readfile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	writer.WriteAll(records)
	writer.Flush()
}
