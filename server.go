package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gorilla/websocket"
)

type handler struct{}

var config = oauth1.NewConfig(os.Getenv("TWITTER_API_KEY_BRANDLESS"), os.Getenv("TWITTER_API_SECRET_BRANDLESS"))
var token = oauth1.NewToken(os.Getenv("TWITTER_TOKEN"), os.Getenv("TWITTER_SECRET"))
var httpClient = config.Client(oauth1.NoContext, token)
var client = twitter.NewClient(httpClient)
var params = &twitter.StreamFilterParams{
	Track:         []string{"kitten"},
	StallWarnings: twitter.Bool(true),
}
var stream, err = client.Streams.Filter(params)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(messageType)
		log.Println(string(p))
		p = []byte("hey can you read this?")
		if err = conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)
	var sokcetHandler handler
	http.Handle("/hey", sokcetHandler)
	log.Println("listening...")
	http.ListenAndServe(":3000", nil)

	// ==================

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}

	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	log.Println("Stopping Stream...")
	stream.Stop()

}
