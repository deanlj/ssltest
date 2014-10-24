package main

import (
	"log"
	"time"

	"github.com/jamescun/ssltest/ssltest"
)

type Task interface {
	Execute()
}

type TestVersion struct {
	SSLTest *SSLTest // ssl test parameters/results
	Name    string   // tls version name
	Version uint16   // tls version
}

func (t TestVersion) Execute() {
	log.Printf("Test: Begin: %s\n", t.Name)

	t0 := time.Now()
	v, err := ssltest.Version(t.SSLTest.Addr, t.Version)
	if err != nil {
		log.Printf("Test: Error: %s: %s\n", t.Name, err)
		return
	}
	t1 := time.Now()

	t.SSLTest.Versions[t.Name] = v

	log.Printf("Test: Complete: %s %t %s\n", t.Name, v, t1.Sub(t0))
}

type TestCipher struct {
	SSLTest *SSLTest // ssl test parameters/results
	Name    string   // tls cipher name
	Cipher  uint16   // tls cipher
}

func (t TestCipher) Execute() {
	log.Printf("Test: Begin: %s\n", t.Name)

	t0 := time.Now()
	c, err := ssltest.Cipher(t.SSLTest.Addr, t.Cipher)
	if err != nil {
		log.Printf("Test: Error: %s: %s\n", t.Name, err)
		return
	}
	t1 := time.Now()

	t.SSLTest.Ciphers[t.Name] = c

	log.Printf("Test: Complete: %s %t %s\n", t.Name, c, t1.Sub(t0))
}
