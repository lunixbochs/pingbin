package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Record interface {
	JSON() ([]byte, error)
	Header() *RecordHeader
	String() string
}

func NewRecordHeader(ip, token, typ string, data []byte) *RecordHeader {
	return &RecordHeader{
		IP:    ip,
		Token: token,
		Type:  typ,
		Time:  time.Now().UTC(),
		Data:  data,
	}
}

type RecordHeader struct {
	IP    string    `json:"ip"`
	Token string    `json:"token"`
	Type  string    `json:"type"`
	Time  time.Time `json:"time"`
	Data  []byte    `json:"data"`
}

func (r *RecordHeader) JSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RecordHeader) Header() *RecordHeader {
	return r
}

func (r *RecordHeader) String() string {
	return fmt.Sprintf("%s %s %s", r.IP, r.Type, r.Token)
}

type HttpRecord struct {
	*RecordHeader
	Path    string      `json:"path"`
	Domain  string      `json:"domain"`
	Headers http.Header `json:"headers"`
}

type DnsRecord struct {
	*RecordHeader
	Domain string `json:"domain"`
}

type IcmpRecord struct {
	*RecordHeader
}
