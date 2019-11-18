package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

// This is a sending channel.
// It sends the bytes to the forward channel on the room struct
func (c *client) read() {
	// this should be done before ending function
	defer c.socket.Close()
	// loop indefinitely
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		// send the bytes to the forward channel.
		c.room.forward <- msg
	}
}

// This is a method that sends data through the
// websocket. It doesn't deal with any channel funcitonality
func (c *client) write() {
	// Close the websocket connection before returning
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
