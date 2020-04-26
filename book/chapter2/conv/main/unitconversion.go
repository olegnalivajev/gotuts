// program that takes numbers as cmd arguments and prints aout various unit conversions with the numbers.
// if no cmd arguments are present, cmd standard input is used to promt user for numbers
package main

import (
	"bufio"
	"fmt"
	"github.com/olegnalivajev/gotuts/book_chapter2/conv"
	"os"
	"strconv"
	"strings"
)

func main()  {
	if args := os.Args[1:]; !(len(args)==0) {
		for _, arg := range args {
			val, _ := strconv.ParseFloat(arg, 64)
			printConversions(val)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Println("Enter a number to convert or 'q' to exit")
			text, _ := reader.ReadString('\n')
			text = strings.ToLower(strings.Trim(text, " \r\n"))
			if text != "q" {
				val, _ := strconv.ParseFloat(text, 64)
				printConversions(val)
			} else {
				fmt.Println("Thanks")
				break
			}
		}
	}
}

func printConversions(val float64) {
	fmt.Printf("%s = %s\n", conv.Centimeter(val), conv.CmToM(conv.Centimeter(val)))
	fmt.Printf("%s = %s\n", conv.Meter(val), conv.MToKm(conv.Meter(val)))
	fmt.Printf("%s = %s\n", conv.Kilometer(val), conv.KmToMile(conv.Kilometer(val)))

	fmt.Printf("%s = %s\n", conv.Celsius(val), conv.CToF(conv.Celsius(val)))
	fmt.Printf("%s = %s\n", conv.Fahrenheit(val), conv.FToC(conv.Fahrenheit(val)))

	fmt.Printf("%s = %s\n", conv.Kilogram(val), conv.KgToTon(conv.Kilogram(val)))
	fmt.Println()
}