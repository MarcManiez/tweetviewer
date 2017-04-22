package main

import (
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gorilla/websocket"
)

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
		p, err := json.Marshal(message)
		log.Println(p)
		if err != nil {
			log.Println(err)
		}
		if err = conn.WriteMessage(1, p); err != nil {
			log.Println(err)
		}
	}
}

func getStream(params *twitter.StreamFilterParams) (*twitter.Stream, error) {
	config := oauth1.NewConfig(os.Getenv("TWITTER_API_KEY_BRANDLESS"), os.Getenv("TWITTER_API_SECRET_BRANDLESS"))
	token := oauth1.NewToken(os.Getenv("TWITTER_TOKEN"), os.Getenv("TWITTER_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client.Streams.Filter(params)
}

func main() {
	params := &twitter.StreamFilterParams{
		Track:         []string{"brands && advertising", "consumerism"},
		Follow:        []string{"12480582", "22151553", "267399199", "308452020", "17540485", "26787673", "713033114092761088", "21778607", "21778607", "485071945", "138845026", "23085995", "570715775", "109224937", "17137891", "17475575", "280557152", "2874230781", "17899654", "534249758", "92793164", "2347131241", "12405142", "19720440", "19720019", "57013560", "398942686", "831488280", "213299248", "126084292", "51119925", "31143489", "347958019", "71026122", "151913390", "20758087", "20094535", "92294003"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := getStream(params)
	if err != nil {
		panic(err)
	}

	log.Println("Control flow...")
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", fs)
	http.Handle("/tweets", &handler{stream: stream})
	log.Println("listening...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println(port)
	e := http.ListenAndServe(port, nil)
	log.Println(":" + port)
	if e != nil {
		panic(e)
	} else {
		log.Println(port)
	}
}
