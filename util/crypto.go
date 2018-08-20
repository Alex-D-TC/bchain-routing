package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"io/ioutil"
	"math/big"
	"os"
)

type EcdsaSignature struct {
	R *big.Int
	S *big.Int
}

func (sig EcdsaSignature) Equal(oth EcdsaSignature) bool {
	return sig.R.Cmp(oth.R) == 0 && sig.S.Cmp(oth.S) == 0
}

func GenerateECDSAKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
}

func Sign(key *ecdsa.PrivateKey, hash [sha256.Size]byte) (EcdsaSignature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, key, hash[:])
	if err != nil {
		return EcdsaSignature{}, err
	}

	return EcdsaSignature{
		R: r,
		S: s,
	}, err
}

func Verify(pubKey *ecdsa.PublicKey, hashData []byte, signature EcdsaSignature) bool {
	return ecdsa.Verify(pubKey, hashData, signature.R, signature.S)
}

func PubKeysEqual(k1 *ecdsa.PublicKey, k2 *ecdsa.PublicKey) bool {
	return k1.X.Cmp(k2.X) == 0 && k1.Y.Cmp(k2.Y) == 0
}

func LoadKeys(path string) (*ecdsa.PrivateKey, error) {
	rawKey, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privKey, err := x509.ParseECPrivateKey(rawKey)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

func WriteKeys(path string, privKey *ecdsa.PrivateKey) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()

	if err != nil {
		return err
	}

	privBytes, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		return err
	}

	file.Write(privBytes)

	return nil
}
