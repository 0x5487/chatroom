package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

func websocketEndpoint() websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		for {
			// Read
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", msg)

			// Write
			err = websocket.Message.Send(ws, msg)
			if err != nil {
				log.Fatal(err)
			}

		}
	})
}
