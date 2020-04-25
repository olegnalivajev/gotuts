package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main()  {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://"+url
		}
		resp, err := http.Get(url)
		if resp != nil {
			_, _ = fmt.Fprintf(os.Stdout, "Status code: %s\n", resp.Status)
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		b, err := io.Copy(os.Stdout,resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}