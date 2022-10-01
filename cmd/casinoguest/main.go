package main

import (
	"flag"
	"fmt"
)

var (
	version = "0.0.0"
	commit  = "none"
	date    = "none"
)

const (
	httpPort = "8080"
	baseURL  = ":" + httpPort
)

func main() {
	printVersion := flag.Bool("version", false, "Get version")
	flag.Parse()
	if *printVersion {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Build Date: %s\n", date)
		fmt.Printf("Build Commit: %s\n", commit)
		return
	}

}
