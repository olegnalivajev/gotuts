package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	fmt.Println("Commands as follows: ")
	for i, arg := range os.Args[0:] {
		fmt.Println(strconv.Itoa(i)	 + " : " + arg)
	}
}
