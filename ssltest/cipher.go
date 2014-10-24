package ssltest

import (
	"crypto/tls"
)

// test if a server supports a particular cipher
func Cipher(addr string, cipher uint16) (bool, error) {
	c, err := tls.Dial("tcp", addr, &tls.Config{
		CipherSuites:             []uint16{cipher},
		PreferServerCipherSuites: false,
	})
	if err != nil {
		if err.Error() == "remote error: handshake failure" {
			return false, nil
		}

		return false, err
	}
	defer c.Close()

	return true, nil
}
