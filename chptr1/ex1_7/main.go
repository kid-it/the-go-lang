package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
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
	}
}
