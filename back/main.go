package main

import (
	customwebsocket "chatapplication/websocket"
	"log"
	"net/http"
)

func serverWs(pool *customwebsocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := customwebsocket.Upgrader(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	client := &customwebsocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	log.Println("Working")
	pool := customwebsocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(pool, w, r)
	})

}

func main() {
	setupRoutes()
	http.ListenAndServe(":7000", nil)
}
