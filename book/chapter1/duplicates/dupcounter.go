// counts number of duplicate words across the files or standard input.
// prints names of those file containing duplicates.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int) // pass by reference
	dupFiles := make(map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // could've used ioutil.ReadFile (would load entire file into memory)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupFiles)
			_ = f.Close()
		}
	}
	fmt.Println("Duplicates were found in: ")
	for file, _ := range dupFiles {
		fmt.Printf("%s ", file)
	}
	fmt.Println()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int, dupFiles map[string]bool)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] > 1 {
			dupFiles[(*f).Name()] = true
		}
		counts[input.Text()]++
	}
}