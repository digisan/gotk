package crypto

import (
	"fmt"
	"testing"
	"time"
)

func TestEnDe(t *testing.T) {

	SetIV([]byte{0x2a, 0x92})

	original := "OK"
	key := []byte(fmt.Sprintf("%d", time.Now().UnixNano())[3:19])

	secret := Encrypt(original, key)
	fmt.Println(secret)
	fmt.Println(string(secret))
	fmt.Printf("%x\n", secret)

	plain := Decrypt(secret, key)
	// fmt.Println(plain)

	if plain == original {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR")
	}
}
