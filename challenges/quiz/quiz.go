package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main()  {

	// prompt user for a file name
	breader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file name: ")
	fn, _, _ := breader.ReadLine()
	fileName := string(fn)

	// open the file
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	creader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := creader.Read()
		if err == io.EOF {
			return
		}
		check(err)

		def := line[0]
		sol := line[1]

		fmt.Println(def)
		ans, _, _ := breader.ReadLine()
		eval(def, sol, string(ans))
	}
}

func eval(def, sol, ans string) {
	if sol == ans {
		fmt.Println("good job")
	} else {
		fmt.Println("nop. that's wrong")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}