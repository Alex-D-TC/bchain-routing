package util

import (
	"crypto/rsa"
	"crypto/x509"
	"io/ioutil"
	"os"
)

func LoadKeys(path string) (*rsa.PrivateKey, error) {
	rawKey, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privKey, err := x509.ParsePKCS1PrivateKey(rawKey)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

func WriteKeys(path string, privKey *rsa.PrivateKey) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()

	if err != nil {
		return err
	}

	privBytes := x509.MarshalPKCS1PrivateKey(privKey)
	file.Write(privBytes)

	return nil
}
