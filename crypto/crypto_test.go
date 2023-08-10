package crypto

import (
	"fmt"
	"testing"
	"time"
)

func TestEnDe(t *testing.T) {

	SetIV([]byte{0x2a, 0x92})

	origin := "OK"
	key := []byte(fmt.Sprintf("%d", time.Now().UnixNano())[3:19])

	secret := Encrypt(origin, key)
	fmt.Println(secret)
	fmt.Println(string(secret))
	fmt.Printf("%x\n", secret)

	plain := Decrypt(secret, key)
	// fmt.Println(plain)

	if plain == origin {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR")
	}
}

func TestEnDeStr(t *testing.T) {

	SetIV([]byte{0x2a, 0x92})

	origin := "OK"
	encoded, key := EncodeStr(origin, nil)
	fmt.Println("encoded:", encoded, "key:", key)

	decoded := DecodeStr(encoded, key)
	fmt.Println("origin:", decoded)
}
