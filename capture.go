package main

import (
	"bytes"
	"encoding/hex"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"os"
	"regexp"
)

func newCapture(device string) (*gopacket.PacketSource, error) {
	if handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever); err != nil {
		return nil, err
	} else if err := handle.SetBPFFilter("icmp or icmp6 or port 53"); err != nil {
		return nil, err
	} else if err := handle.SetDirection(pcap.DirectionIn); err != nil {
		return nil, err
	} else {
		return gopacket.NewPacketSource(handle, handle.LinkType()), nil
	}
}

func findToken(p []byte) string {
	for i := 0; i < len(p)-14; i++ {
		part := p[i : i+14]
		if bytes.Count(p, []byte("\x00"+string(part)+"\x00")) >= 2 {
			return hex.EncodeToString(part)
		}
	}
	return ""
}

var tokenRe = regexp.MustCompile(`^[a-zA-Z0-9]{28}$`)

func Capture(device string) (<-chan Record, error) {
	capture, err := newCapture(os.Args[1])
	if err != nil {
		return nil, err
	}
	ret := make(chan Record)
	go func() {
		for packet := range capture.Packets() {
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer == nil {
				continue
			}
			ip, _ := ipLayer.(*layers.IPv4)
			for _, layer := range packet.Layers() {
				switch layer := layer.(type) {
				case *layers.DNS:
					if len(layer.Answers) > 0 {
						continue
					}
					for _, q := range layer.Questions {
						domain := string(q.Name)
						token := domain
						if !tokenRe.MatchString(domain) {
							token = ""
						}
						header := NewRecordHeader(ip.SrcIP.String(), token, "dns", packet.Data())
						ret <- &DnsRecord{header, domain}
					}
				case *layers.ICMPv4:
					if (layer.TypeCode >> 8) == layers.ICMPv4TypeEchoRequest {
						token := findToken(layer.Payload)
						header := NewRecordHeader(ip.SrcIP.String(), token, "icmp", packet.Data())
						ret <- &IcmpRecord{header}
					}
				}
			}
		}
	}()
	return ret, nil
}
