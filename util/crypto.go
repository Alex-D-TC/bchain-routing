package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"io/ioutil"
	"os"
)

func Sign(key *rsa.PrivateKey, hashFunc crypto.Hash, hash [sha256.Size]byte) ([]byte, error) {
	return rsa.SignPSS(rand.Reader, key, hashFunc, hash[:], nil)
}

func Verify(pubKey *rsa.PublicKey, hash crypto.Hash, hashData []byte, signature []byte) error {
	return rsa.VerifyPSS(pubKey, hash, hashData, signature, nil)
}

func PubKeysEqual(k1 *rsa.PublicKey, k2 *rsa.PublicKey) bool {
	return k1.N.Cmp(k2.N) == 0 && k1.E == k2.E
}

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
