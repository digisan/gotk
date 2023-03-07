package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"log"
	"sync"
)

var (
	commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	mu       sync.Mutex
)

func SetIV(iv []byte) {

	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < len(commonIV); i++ {
		if i < len(iv) {
			commonIV[i] = iv[i]
		}
	}
}

func Encrypt(plain string, key []byte) []byte {

	mu.Lock()
	defer mu.Unlock()

	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error: NewCipher(%d bytes) = %s", len(key), err)
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherBuf := make([]byte, len(plain))
	cfb.XORKeyStream(cipherBuf, []byte(plain))
	// fmt.Printf("%s => %x\n", []byte(plain), cipherBuf)

	return cipherBuf // fmt.Sprintf("%x", cipherBuf)
}

func Decrypt(cipherBuf, key []byte) string {

	mu.Lock()
	defer mu.Unlock()

	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error: NewCipher(%d bytes) = %s", len(key), err)
	}

	cfbDec := cipher.NewCFBDecrypter(c, commonIV)
	plainBuf := make([]byte, 1024)
	cfbDec.XORKeyStream(plainBuf, cipherBuf)
	plainBuf = bytes.TrimRight(plainBuf, "\x00")
	// fmt.Printf("%x => %s\n", cipherBuf, plainBuf)
	return string(plainBuf)
}
