package main

import (
	"goexample/simpleclient"
	"goexample/simpleserver"
	"os"
)

func main() {
	var doRun string
	if len(os.Args) > 1 {
		doRun = os.Args[1]
	}
	switch doRun {
	case "server":
		simpleserver.RunServer()
	default:
		simpleclient.RunClient()
	}
}
