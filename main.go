package main

import (
	"flag"
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
		print("URL : " + url + "\n")
	}
	resp, err := http.Get(url) // Note pointer dereference using *

	// If there is an error or non-200 status, exit with 1 signaling unsuccessful check.
	if nil != err || 200 != resp.StatusCode {
		if *debug {
			if nil != err {
				print("FAIL Error : ")
				print(err.Error())
				print("\n")
			} else {
				print("FAIL StatusCode : ")
				print(resp.StatusCode)
				print("\n")
			}
		}
		os.Exit(1)
	}
	os.Exit(0)
}
