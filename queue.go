package main

import (
	"log"
	"sync"
)

type Topic struct {
	History     []Record
	Subscribers []chan Record
}

var topics = make(map[string]*Topic)
var topicLock sync.Mutex

func getTopic(name string) *Topic {
	t, ok := topics[name]
	if !ok {
		t = &Topic{}
		topics[name] = t
	}
	return t
}

func Autopub(c <-chan Record) {
	go func() {
		for r := range c {
			topic := r.Header().Token
			if topic == "" {
				topic = "public"
			}
			Publish(topic, r)
		}
	}()
}

func Publish(name string, r Record) {
	log.Printf("%s\n", r)
	topicLock.Lock()
	defer topicLock.Unlock()
	topic := getTopic(name)
	for _, c := range topic.Subscribers {
		c <- r
	}
	topic.History = append(topic.History, r)
	if len(topic.History) > 999 {
		topic.History = topic.History[:999]
	}
}

func Subscribe(name string) chan Record {
	topicLock.Lock()
	defer topicLock.Unlock()
	c := make(chan Record)
	topic := getTopic(name)
	topic.Subscribers = append(topic.Subscribers, c)
	return c
}

func Unsubscribe(name string, remove chan Record) {
	topicLock.Lock()
	defer topicLock.Unlock()
	topic := getTopic(name)
	min := len(topic.Subscribers) - 1
	if min < 0 {
		min = 0
	}
	newSubs := make([]chan Record, 0, min)
	for _, c := range topic.Subscribers {
		if c != remove {
			newSubs = append(newSubs, c)
		}
	}
	topic.Subscribers = newSubs
}

func History(name string) []Record {
	topicLock.Lock()
	defer topicLock.Unlock()
	topic := getTopic(name)
	tmp := make([]Record, len(topic.History))
	for i := 0; i < len(tmp); i++ {
		tmp[i] = topic.History[len(tmp)-i-1]
	}
	return tmp
}
