// fetches several URLs concurrently. prints how long it took to fetch each one.
// prints how long the program took in total
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main()  {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts a goroutine
	}
	// create a file
	f, err := os.Create("out")
	check(err)
	defer f.Close()

	// for each goroutine print the results
	for range os.Args[1:] {
		res := <-ch + "\n"
		_, err := f.WriteString(res) // writes results from each goroutine into the file
		check(err)

		fmt.Print(res) // prints into the console as well
	}
	elapsed := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Println(elapsed)
	_, err = f.WriteString(elapsed)
	check(err)

	_ = f.Sync() // flush contents of a file to a stable storage
}

func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // result is sent to channel
	}
	nbytes, err := io.Copy(ioutil.Discard, res.Body)
	_ = res.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}