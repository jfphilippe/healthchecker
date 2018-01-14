package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	debug := flag.Bool("debug", false, "print debug infos")
	port := flag.String("port", "80", "port on localhost to check")
	endpoint := flag.String("endpoint", "/health", "local url to check (start with /)")
	flag.Parse()

	url := "http://127.0.0.1:" + *port + *endpoint
	if *debug {
		fmt.Printf("URL : %s\n", url)
	}
	resp, err := http.Get(url) // Note pointer dereference using *

	// If there is an error or non-200 status, exit with 1 signaling unsuccessful check.
	if err != nil || resp.StatusCode != 200 {
		if *debug {
			if nil != err {
				fmt.Printf("FAIL Error : %s\n", err)
			} else {
				fmt.Printf("FAIL StatusCode : %d\n", resp.StatusCode)
			}
		}
		os.Exit(1)
	}
	os.Exit(0)
}
