package util

import (
	"crypto/ecdsa"
	"crypto/sha256"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateECDSAKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func Sign(key *ecdsa.PrivateKey, hash [sha256.Size]byte) ([]byte, error) {
	return crypto.Sign(hash[:], key)
}

func Verify(pubKey ecdsa.PublicKey, hashData []byte, signature []byte) bool {
	return crypto.VerifySignature(crypto.CompressPubkey(&pubKey), hashData, signature)
}

func PubKeysEqual(k1 ecdsa.PublicKey, k2 ecdsa.PublicKey) bool {
	return k1.X.Cmp(k2.X) == 0 && k1.Y.Cmp(k2.Y) == 0
}

func LoadKeys(path string) (*ecdsa.PrivateKey, error) {
	return crypto.LoadECDSA(path)
}

func WriteKeys(path string, privKey *ecdsa.PrivateKey) error {
	return crypto.SaveECDSA(path, privKey)
}

func MarshalPubKey(key *ecdsa.PublicKey) []byte {
	return crypto.FromECDSAPub(key)
}

func UnmarshalPubKey(rawKey []byte) (*ecdsa.PublicKey, error) {
	return crypto.UnmarshalPubkey(rawKey)
}
