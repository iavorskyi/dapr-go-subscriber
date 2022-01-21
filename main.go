package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/gorilla/mux"
)

//code
var sub = &common.Subscription{
	PubsubName: "pubsub",
	Topic:      "project-lifecycle",
	Route:      "/checkout",
}

func main() {
	log.Println("Listening :8000")
	r := mux.Router(*handlers())
	s := daprd.NewServiceWithMux(":8000", &r)

	err := s.AddTopicEventHandler(sub, eventHandler)
	if err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("Subscriber received: %s", e.Data)
	return false, nil
}

func myOtherHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In the other endpoint")
}

func handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", myOtherHandler).Methods("GET")

	return r
}
