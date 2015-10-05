package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

func generateToken() string {
	var token [14]byte
	rand.Read(token[:])
	return hex.EncodeToString(token[:])
}

var historyPathRe = regexp.MustCompile(`^/(public|[a-fA-F0-9]{28})/history$`)
var tokenPathRe = regexp.MustCompile(`^/(public|[a-fA-F0-9]{28})(/nojs)?$`)
var pathPingRe = regexp.MustCompile(`^/p/(public|[a-fA-F0-9]{28})$`)

func Http() (<-chan Record, error) {
	ret := make(chan Record)
	sockio, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}
	sockio.On("connection", func(so socketio.Socket) {
		so.On("subscribe", func(channel string) {
			go subscribe(so, channel)
		})
	})
	http.Handle("/socket.io/", sockio)
	indexHtml, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		log.Fatal(err)
	}
	nojsHtml, err := template.ParseFiles("frontend/index-nojs.html")
	if err != nil {
		log.Fatal(err)
	}
	compiled := true
	if _, err := os.Stat("frontend/main.full.js"); os.IsNotExist(err) {
		compiled = false
	}
	fileServer := http.FileServer(http.Dir("frontend"))
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
		if path == "/" {
			token := generateToken()
			http.Redirect(w, r, "/"+token, 302)
		} else if match := tokenPathRe.FindStringSubmatch(path); len(match) < 2 {
			fileServer.ServeHTTP(w, r)
		} else {
			token := match[1]
			template := indexHtml
			if len(match) >= 3 && match[2] == "/nojs" {
				template = nojsHtml
			}
			err = template.Execute(w, &struct {
				Compiled bool
				Token    string
				History  []Record
			}{compiled, token, History(token)})
			if err != nil {
				log.Println(err)
			}
		}
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
