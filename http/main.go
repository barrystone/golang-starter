package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println(resp)

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	fmt.Printf("\n\n=== Response (1) ===\n\n")
	bs := make([]byte, 99999)
	tee.Read(bs)
	fmt.Println(string(bs))

	fmt.Printf("\n\n=== Response (2-1 os.Stdout) ===\n\n")
	_, err = io.Copy(os.Stdout, bytes.NewReader(buf.Bytes()))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("\n\n=== Response (2-2 Custom Writer) ===\n\n")
	lw := logWriter{}
	_, err = io.Copy(lw, bytes.NewReader(buf.Bytes()))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
