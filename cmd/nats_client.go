package main

import (
	"log"

	"github.com/nats-io/nats.go@v1.10.0"
)

func main() {
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
}
