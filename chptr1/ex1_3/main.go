package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	stop := time.Since(start).Nanoseconds()
	fmt.Println(stop)
}

// func main() {
// 	start := time.Now()
// 	s, spc := "", " "
// 	for _, i := range os.Args[1:] {
// 		s += i + spc
// 		spc = " "
// 	}
// 	fmt.Println(s)
// 	stop := time.Since(start).Nanoseconds()
// 	fmt.Println(stop)
// }
