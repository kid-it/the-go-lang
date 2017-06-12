package main

import "fmt"
import "bufio"
import "os"

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Println(files)
	if len(files) == 0 {
		fmt.Println("reading from input")
		countlines(os.Stdin, counts)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error opening files %v\n", err)
				continue
			}
			countlines(f, counts)
			f.Close()
		}
	}
	fmt.Println()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%v %v \n", line, n)
		}
	}
}

func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "quit" {
			break
		}
	}
}
