package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

func generateToken() string {
	var token [14]byte
	rand.Read(token[:])
	return hex.EncodeToString(token[:])
}

var historyPathRe = regexp.MustCompile(`^/([a-fA-F0-9]{28})/history$`)
var tokenPathRe = regexp.MustCompile(`^/([a-fA-F0-9]{28})$`)
var pathPingRe = regexp.MustCompile(`^/p/([a-fA-F0-9]{28})$`)

func Http(listen string) (<-chan Record, error) {
	ret := make(chan Record)
	sockio, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}
	sockio.OnEvent("/", "subscribe", func(s socketio.Conn, topic string) {
		events := Subscribe(topic)
		sockio.OnDisconnect("/", func(s socketio.Conn, msg string) {
			Unsubscribe(topic, events)
		})
		go func() {
			for e := range events {
				v, err := json.Marshal(e)
				if err != nil {
					log.Println(err)
				} else {
					s.Emit(topic, string(v))
				}
			}
		}()
	})
	http.Handle("/socket.io/", sockio)
	appHtml, err := template.ParseFiles("templates/app.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		path := r.URL.Path
		if path == "/favicon.ico" {
			return
		}
		matches := historyPathRe.FindStringSubmatch(path)
		if len(matches) > 1 {
			token := matches[1]
			history := History(token)
			if history == nil {
				history = []Record{}
			}
			j, err := json.Marshal(history)
			if err != nil {
				log.Println(err)
			} else {
				w.Write(j)
			}
			return
		}
		if !tokenPathRe.MatchString(path) {
			token := generateToken()
			http.Redirect(w, r, "/"+token, 302)
		} else {
			token := path[1:]
			err = appHtml.Execute(w, &struct {
				Token   string
				History []Record
			}{token, History(token)})
			if err != nil {
				log.Println(err)
			}
		}
	})
	http.HandleFunc("/p/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		matches := pathPingRe.FindStringSubmatch(path)
		var token string
		if len(matches) < 2 {
			return
		}
		token = matches[1]
		ip := r.RemoteAddr
		if v, ok := r.Header["X-Forwarded-For"]; ok && len(v) > 0 {
			ip = v[0]
		}
		host := r.Host
		if v, ok := r.Header["X-Real-Host"]; ok && len(v) > 0 {
			host = v[0]
		}
		header := NewRecordHeader(ip, token, "http", nil)
		ret <- &HttpRecord{
			RecordHeader: header,
			Domain:       host,
			Path:         path,
			Headers:      r.Header,
		}
		if err != nil {
			log.Println(err)
		}
	})
	go func() {
		err := http.ListenAndServe(listen, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return ret, nil
}
