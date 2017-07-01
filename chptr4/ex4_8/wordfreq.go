package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PrintResults(mmap map[string]int) {
	for name, freq := range mmap {
		fmt.Printf("%q %d \n", name, freq)
	}
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	count := make(map[string]int)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		line := input.Text()
		count[line]++
	}
	PrintResults(count)
}
