// min prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const HttpPrefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(HttpPrefix, url) == false {
			url = HttpPrefix + url
		}
		resp, err := http.Get(url)
		fmt.Printf(resp.Status)
		if err != nil {
			fmt.Printf("status: %s\n", resp.Status)
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Printf("Fatal error: %v\n", err)
		}
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
