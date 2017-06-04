package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <listen address> <capture interface>\n", os.Args[0])
		os.Exit(1)
	}
	capture, err := Capture(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	Autopub(capture)
	httpcap, err := Http(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	Autopub(httpcap)
	<-make(chan int)
}
