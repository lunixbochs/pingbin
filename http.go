package main

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"regexp"
)

var pathPingRe *regexp.Regexp = regexp.MustCompile(`^/p/([a-f0-9]{32})$`)

func Http() (<-chan Record, error) {
	ret := make(chan Record)
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}
	server.On("connection", func(so socketio.Socket) {
		so.On("subscribe", func(channel string) {
			go subscribe(so, channel)
		})
	})
	http.HandleFunc("/p/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		matches := pathPingRe.FindStringSubmatch(path)
		var token string
		if len(matches) > 1 {
			token = matches[1]
		} else {
			token = "public"
		}
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
		err := http.ListenAndServe(":5007", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return ret, nil
}
