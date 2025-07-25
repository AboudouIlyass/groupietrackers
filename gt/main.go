package main

import (
	"os"

	"gt/server"
)

func main() {
	if len(os.Args) != 1 {
		return
	}
	server.Server()
}
