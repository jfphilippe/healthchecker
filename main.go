/*
Program to test and HTTP end point
Copyright (c) 2018 Jean-François PHILIPPE

Try to call some local URL (127.0.0.1) , return 0 if the URL returns a 2xx code, 1 otherwise
*/
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	debug := flag.Bool("debug", false, "print debug infos")
	port := flag.Uint("port", 80, "HTTP port of URL check")
	endpoint := flag.String("endpoint", "/health", "local url to check (start with /)")
	timeout := flag.Duration("timeout", time.Second*10, "HTTP timeout in seconds")
	ip := flag.String("ip", "127.0.0.1", "IP Address of end point")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Health checker\n")
		fmt.Fprintf(os.Stderr, "    version: %s\n", "0.1")
		fmt.Fprintf(os.Stderr, "    copyright: %s\n", "jeff")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// build the URL to call
	url := "http://" + *ip
	if 80 != *port {
		url = url + ":" + strconv.Itoa(int(*port))
	}
	url = url + *endpoint
	if *debug {
		print("URL : " + url + "\n")
	}
	// Create a Client & Specify a timeout
	var netClient = &http.Client{
		Timeout: *timeout,
	}

	resp, err := netClient.Get(url)

	// If there is an error or non-2xx status, exit with 1 signaling unsuccessful check.
	if nil != err || (200 > resp.StatusCode || 300 <= resp.StatusCode) {
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
