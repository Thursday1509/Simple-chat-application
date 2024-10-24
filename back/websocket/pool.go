package customwebsocket

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case Client := <-pool.Register:
			pool.Clients[Client] = true
			fmt.Println("Total Connection POOL :- ", len(pool.Clients))
			for K, _ := range pool.Clients {
				fmt.Println(K)
				K.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined"})
			}
		case Client := <-pool.Unregister:
			delete(pool.Clients, Client)
			fmt.Println("Total COnnection POOL :- ", len(pool.Clients))
			for K, _ := range pool.Clients {
				fmt.Println(K)
				K.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnection"})
			}
		case msg := <-pool.Broadcast:
			fmt.Println("broadcasting a message")
			for K, _ := range pool.Clients {
				if err := K.Conn.WriteJSON(msg); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
