package main

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {

		case client := <-r.join:
			// joining
			r.clients[client] = true

		case client := <-r.leave:
			//leaving
			delete(r.clients, client)
			close(client.send)

		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
			}

		}
	}
}
