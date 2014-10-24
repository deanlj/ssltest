package ssltest

import (
	"crypto/tls"
)

// test if a server supports a particular SSL/TLS version
func Version(addr string, version uint16) (bool, error) {
	c, err := tls.Dial("tcp", addr, &tls.Config{
		MinVersion: version,
		MaxVersion: version,
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
