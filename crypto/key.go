package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// `openssl genrsa -out cert/id_rsa 4096`
// `openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub`

/////////////////////////////////////////////////////////////////////

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	prvKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	return prvKey, &prvKey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(prvKey *rsa.PrivateKey, fOutPath string) (string, error) {
	prvKey_bytes := x509.MarshalPKCS1PrivateKey(prvKey)
	prvKey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: prvKey_bytes,
		},
	)
	ret := string(prvKey_pem)
	if len(fOutPath) != 0 {
		if err := os.WriteFile(fOutPath, prvKey_pem, os.ModePerm); err != nil {
			return "", err
		}
	}
	return ret, nil
}

func ParseRsaPrivateKeyFromPemStr(prvPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(prvPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}
	prvKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return prvKey, nil
}

func ExportRsaPublicKeyAsPemStr(pubKey *rsa.PublicKey, fOutPath string) (string, error) {
	pubKey_bytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	pubKey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubKey_bytes,
		},
	)
	ret := string(pubKey_pem)
	if len(fOutPath) != 0 {
		if err := os.WriteFile(fOutPath, pubKey_pem, os.ModePerm); err != nil {
			return "", err
		}
	}
	return ret, nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch pub := pubKey.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}
