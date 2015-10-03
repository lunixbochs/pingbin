package main

import (
	"encoding/json"
	"github.com/googollee/go-socket.io"
	"log"
)

func subscribe(so socketio.Socket, topic string) {
	events := Subscribe(topic)
	so.On("disconnection", func() {
		Unsubscribe(topic, events)
	})
	for e := range events {
		v, err := json.Marshal(e)
		if err != nil {
			log.Println(err)
		} else {
			so.Emit(topic, v)
		}
	}
}
