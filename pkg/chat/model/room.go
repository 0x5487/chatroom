package model

import (
	"encoding/json"
)

type Room struct {
	Name string

	// Registered clients.
	clients map[*Client]bool

	// broadcast messages to all clients.
	broadcast chan Payload

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	maxPeople int
}

func NewRoom(name string) *Room {
	room := &Room{
		Name:       name,
		broadcast:  make(chan Payload, 100),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}

	go room.start()
	return room
}

func (rm *Room) Join(c *Client) error {
	c.room = rm
	rm.register <- c

	go c.readPump()
	go c.writePump()

	payload := Payload{
		Kind:   "join",
		Member: c.member,
	}

	rm.broadcast <- payload

	return nil
}

func (rm *Room) Count() int {
	return len(rm.clients)
}

func (rm *Room) SetMaxPeople(max int) {
	rm.maxPeople = max
}

func (rm *Room) start() {
	for {
		select {
		case client := <-rm.register:
			rm.clients[client] = true
		case client := <-rm.unregister:
			if _, ok := rm.clients[client]; ok {
				delete(rm.clients, client)
				close(client.send)
			}
		case payload := <-rm.broadcast:

			jsonB, err := json.Marshal(payload)
			if err != nil {
				return
			}

			for client := range rm.clients {
				select {
				case client.send <- jsonB:
				default:
					close(client.send)
					delete(rm.clients, client)
				}
			}
		}
	}
}
