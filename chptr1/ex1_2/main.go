package main

import (
	"fmt"
	"os"
)

func main() {
	j := 1
	for _, i := range os.Args[1:] {
		fmt.Printf("%v %v \n", i, j)
		j++
	}
}
