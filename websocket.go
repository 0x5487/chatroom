package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

type WSClient struct {
}

type WSHub struct {
	clients []*WSClient
}

func (hub *WSHub) BroadCast(message []byte, ignore *WSClient) error {
	return nil
}

func (hub *WSHub) Add(client *WSClient) error {
	return nil
}

func (hub *WSHub) Remove(client *WSClient) error {
	return nil
}

func (hub *WSHub) Count() (int, error) {
	return 0, nil
}

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
