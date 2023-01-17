package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"

	"github.com/klauspost/compress/zstd"
)

func main() {
	var b bytes.Buffer

	file, _ := os.Open("test.pdf")

	read := bufio.NewReader(file)

	data, _ := ioutil.ReadAll(read)

	// w := zlib.NewWriter(&b)
	// w.Write(data)
	// w.Close()
	w, _ := zstd.NewWriter(&b)
	w.Write(data)
	w.Close()
	ioutil.WriteFile("test.zstd", b.Bytes(), 0644)
}
