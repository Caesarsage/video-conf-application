package main

import (
	"log"
	"videochat/internals/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
