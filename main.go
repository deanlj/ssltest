package main

import (
	"crypto/tls"
	"flag"
	"log"
	"os"
	"sync"
	"time"
)

var (
	Hostname = flag.String("hostname", "localhost", "hostname to test")
	Port     = flag.Int("port", 443, "port to test")
	Workers  = flag.Int("workers", 10, "concurrency to execute tests")
	JSON     = flag.Bool("json", false, "output to JSON")
)

func main() {
	flag.Parse()
	log.SetOutput(os.Stderr)

	test := NewSSLTest(*Hostname, *Port)
	log.Printf("Testing host %s...\n", test.Addr)

	// tests to execute against host
	var tasks = []Task{
		TestVersion{SSLTest: test, Name: "SSLv3", Version: tls.VersionSSL30},
		TestVersion{SSLTest: test, Name: "TLS1.0", Version: tls.VersionTLS10},
		TestVersion{SSLTest: test, Name: "TLS1.1", Version: tls.VersionTLS11},
		TestVersion{SSLTest: test, Name: "TLS1.2", Version: tls.VersionTLS12},

		TestCipher{SSLTest: test, Name: "TLS_RSA_WITH_RC4_128_SHA", Cipher: tls.TLS_RSA_WITH_RC4_128_SHA},
		TestCipher{SSLTest: test, Name: "TLS_RSA_WITH_3DES_EDE_CBC_SHA", Cipher: tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA},
		TestCipher{SSLTest: test, Name: "TLS_RSA_WITH_AES_128_CBC_SHA", Cipher: tls.TLS_RSA_WITH_AES_128_CBC_SHA},
		TestCipher{SSLTest: test, Name: "TLS_RSA_WITH_AES_256_CBC_SHA", Cipher: tls.TLS_RSA_WITH_AES_256_CBC_SHA},
		TestCipher{SSLTest: test, Name: "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA", Cipher: tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA},
		TestCipher{SSLTest: test, Name: "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA", Cipher: tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA},
		TestCipher{SSLTest: test, Name: "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA", Cipher: tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA},
		TestCipher{SSLTest: test, Name: "TLS_ECDHE_RSA_WITH_RC4_128_SHA", Cipher: tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA},
		TestCipher{SSLTest: test, Name: "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", Cipher: tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA},
		TestCipher{SSLTest: test, Name: "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA", Cipher: tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA},
		TestCipher{SSLTest: test, Name: "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA", Cipher: tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA},
		TestCipher{SSLTest: test, Name: "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", Cipher: tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256},
		TestCipher{SSLTest: test, Name: "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", Cipher: tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256},
	}

	test.StartedAt = time.Now()
	work(tasks, *Workers)
	test.FinishedAt = time.Now()

	// output format
	if *JSON == true {
		JSONOutput(test)
	} else {
		HumanOutput(test)
	}
}

// execute an array of tasks
func work(tasks []Task, workers int) {
	tchan := make(chan Task)
	wg := new(sync.WaitGroup)

	// launch workers
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go worker(tchan, wg)
	}

	// add tasks to queue
	for _, task := range tasks {
		tchan <- task
	}

	close(tchan)
	wg.Wait()
}

// task execution worker
func worker(tasks chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		task.Execute()
	}
}
