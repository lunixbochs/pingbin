package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Record interface {
	JSON() ([]byte, error)
	Header() *RecordHeader
}

func NewRecordHeader(ip, token, typ string, data []byte) *RecordHeader {
	return &RecordHeader{
		IP:    ip,
		Token: token,
		Type:  typ,
		Time:  time.Now().UTC().Unix(),
		Data:  data,
	}
}

type RecordHeader struct {
	IP    string `json:"ip"`
	Token string `json:"token"`
	Type  string `json:"type"`
	Time  int64  `json:"time"`
	Data  []byte `json:"data"`
}

func (r *RecordHeader) JSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RecordHeader) Header() *RecordHeader {
	return r
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
