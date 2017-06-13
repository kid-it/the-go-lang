package main

import "fmt"
import "bufio"
import "os"

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	fmt.Println(files)
	if len(files) == 0 {
		fmt.Println("Must suply file name")
		return
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error opening files %v\n", err)
				continue
			}
			countlines(f, counts, arg)
			f.Close()
		}
	}
	fmt.Println()
	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%v %v \n", line, files)
		}
	}
}

func countlines(f *os.File, counts map[string][]string, file string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if file != "" {
			counts[input.Text()] = appendIfUnique(counts[input.Text()], file)
		}
	}
}

func appendIfUnique(strSlice []string, tar string) []string {
	for _, j := range strSlice[0:] {
		if j == tar {
			return strSlice
		}
	}
	strSlice = append(strSlice, tar)
	return strSlice
}
