package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "applicatoin/json")
	//json化するもとのデータ
	source := map[string]string{
		"Hello": "World",
	}
	gz := gzip.NewWriter(w)
	writer := io.MultiWriter(gz, os.Stdout)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", " ")
	encoder.Encode(source)
	gz.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
