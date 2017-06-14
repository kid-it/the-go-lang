package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("test")
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		_, err1 := io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err1 != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(res.Status)
	}
}
