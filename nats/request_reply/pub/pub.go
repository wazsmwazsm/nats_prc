package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://192.168.108.100:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for i := 0; i < 10; i++ {
		// Send the request
		msg, err := nc.Request("time", nil, time.Second)
		if err != nil {
			log.Fatal(err)
		}

		// Use the response
		log.Printf("Reply: %s", msg.Data)
		time.Sleep(time.Second)
	}

	// Close the connection
	nc.Close()
}
