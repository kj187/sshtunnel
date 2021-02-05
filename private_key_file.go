package sshtunnel

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func PrivateKeyFile(file string, passphrase []byte) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	var key ssh.Signer
	if passphrase != nil {
		key, err = ssh.ParsePrivateKeyWithPassphrase(buffer, passphrase)
	} else {
		key, err = ssh.ParsePrivateKey(buffer)
	}

	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}
