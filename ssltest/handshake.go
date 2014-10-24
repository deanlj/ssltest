package ssltest

// simulate the default handshake of a platform, given then TLS version and cipher
func Handshake(addr string, version, cipher uint16) (bool, error) {
	v, err := Version(addr, version)
	if err != nil {
		return false, err
	}
	if v == false {
		return false, nil
	}

	c, err := Cipher(addr, cipher)
	if err != nil {
		return false, err
	}
	if c == false {
		return false, nil
	}

	return true, nil
}
