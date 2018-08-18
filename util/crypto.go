package util

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"os"
)

func LoadKeys(path string) (*rsa.PrivateKey, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	line, _, _ := reader.ReadLine()

	privKey, err := x509.ParsePKCS1PrivateKey(line)
	if err != nil {
		fmt.Println(err)
	}

	return privKey, nil
}

func WriteKeys(path string, privKey *rsa.PrivateKey) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	if err != nil {
		return err
	}

	privBytes := x509.MarshalPKCS1PrivateKey(privKey)
	file.Write(privBytes)

	return nil
}
