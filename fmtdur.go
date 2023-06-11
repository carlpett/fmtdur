package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: fmtdur [seconds]")
		os.Exit(1)
	}
	var d float64
	_, err := fmt.Sscanf(os.Args[1], "%f", &d)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse %q\n", os.Args[1])
		os.Exit(1)
	}
	dur := time.Duration(d * float64(time.Second))
	fmt.Printf("%v\n", dur)
}
