package main

import "github.com/clinotes/client/cmd"

var (
	// VERSION is set during build
	VERSION = "latest"
)

func main() {
	cmd.Execute(VERSION)
}
