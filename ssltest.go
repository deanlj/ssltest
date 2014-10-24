package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNotSSL = errors.New("not ssl host")
	ErrNoAddr = errors.New("no address returned from dns query")
)

type SSLTest struct {
	Hostname string `json:"hostname"` // hostname
	Addr     string `json:"addr"`     // address

	Versions map[string]bool `json:"versions"` // supported TLS versions
	Ciphers  map[string]bool `json:"ciphers"`  // supported TLS ciphers

	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
}

// create a new ssl test struct filled with certificate/host information
func NewSSLTest(hostname string, port int) (test *SSLTest) {
	test = &SSLTest{
		Hostname: hostname,
		Addr:     fmt.Sprintf("%s:%d", hostname, port),
		Versions: make(map[string]bool),
		Ciphers:  make(map[string]bool),
	}

	return
}

// output an SSLTest struct as JSON
func (t SSLTest) ToJSON() []byte {
	dump, _ := json.Marshal(t)
	return dump
}
