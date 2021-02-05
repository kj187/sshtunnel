package sshtunnel

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func PrivateKeyFile(file string, passphrase []byte) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var key ssh.Signer
	if passphrase != nil {
		key, err = ssh.ParsePrivateKeyWithPassphrase(buffer, passphrase)
	} else {
		key, err = ssh.ParsePrivateKey(buffer)
	}

	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(key), nil
}
