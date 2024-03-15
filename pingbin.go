package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <listen address>:<listen port> <capture interface> <http host>\n", os.Args[0])
		os.Exit(1)
	}
	capture, err := Capture(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	Autopub(capture)
	httpcap, err := Http(os.Args[1], os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	Autopub(httpcap)
	<-make(chan int)
}
