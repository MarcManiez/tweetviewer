package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gorilla/websocket"
)

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type handler struct {
	stream *twitter.Stream
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for message := range h.stream.Messages {
		demux := twitter.NewSwitchDemux()
		demux.Tweet = func(tweet *twitter.Tweet) {
			fmt.Println(tweet.Text)
			p := []byte(tweet.Text)
			if err = conn.WriteMessage(1, p); err != nil {
				log.Println(err)
			}
		}
		demux.Handle(message)
	}
}

func getStream(params *twitter.StreamFilterParams) (*twitter.Stream, error) {
	config := oauth1.NewConfig(os.Getenv("TWITTER_API_KEY_BRANDLESS"), os.Getenv("TWITTER_API_SECRET_BRANDLESS"))
	token := oauth1.NewToken(os.Getenv("TWITTER_TOKEN"), os.Getenv("TWITTER_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client.Streams.Filter(params)
}

func readTweets(s *twitter.Stream) {
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}
	go demux.HandleChan(s.Messages)
}

func main() {
	params := &twitter.StreamFilterParams{
		Track:         []string{"kitten"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := getStream(params)
	if err != nil {
		panic(err)
	}
	go readTweets(stream)

	// ==================

	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)
	http.Handle("/tweets", &handler{stream: stream})
	log.Println("listening...")
	http.ListenAndServe(":3000", nil)
}
