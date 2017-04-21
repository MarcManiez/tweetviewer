package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type we struct{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (yo we) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		} else {
			log.Println(messageType)
			log.Println(string(p))
		}
		p = []byte("hey can you read this?")
		if err = conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)
	var lol we
	http.Handle("/hey", lol)
	log.Println("listening...")
	http.ListenAndServe(":3000", nil)
}
