package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	nats "github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("kaa.v1.events.%s.endpoint.traffic-reporting.payload", ch)

	var Stop = make(chan os.Signal, 1)
	signal.Notify(Stop,
		syscall.SIGTERM,
		syscall.SIGINT,
		// syscall.SIGKILL,
		syscall.SIGABRT,
	)

	for {
		select {
		case msg := <-ch:
			log.Println(msg)
		case <-Stop:
			// Unsubscribe
			sub.Unsubscribe()
			// Drain
			sub.Drain()
		}
	}
}
