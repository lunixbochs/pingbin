package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <capture interface>\n", os.Args[0])
		os.Exit(1)
	}
	capture, err := Capture(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	Autopub(capture)
	httpcap, err := Http()
	if err != nil {
		log.Fatal(err)
	}
	Autopub(httpcap)
	<-make(chan int)
}
