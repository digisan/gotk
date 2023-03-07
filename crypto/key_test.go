package crypto

import (
	"fmt"
	"log"
	"testing"
)

func TestKey(t *testing.T) {

	// Create the keys
	prv, pub := GenerateRsaKeyPair()

	// Export the keys to pem string
	prv_pem, err := ExportRsaPrivateKeyAsPemStr(prv, "private.pem")
	if err != nil {
		log.Fatalln(err)
	}
	pub_pem, err := ExportRsaPublicKeyAsPemStr(pub, "public.pem")
	if err != nil {
		log.Fatalln(err)
	}

	// Import the keys from pem string
	prv_parsed, err := ParseRsaPrivateKeyFromPemStr(prv_pem)
	if err != nil {
		log.Fatalln(err)
	}
	pub_parsed, err := ParseRsaPublicKeyFromPemStr(pub_pem)
	if err != nil {
		log.Fatalln(err)
	}

	// Export the newly imported keys
	prv_parsed_pem, err := ExportRsaPrivateKeyAsPemStr(prv_parsed, "")
	if err != nil {
		log.Fatalln(err)
	}
	pub_parsed_pem, err := ExportRsaPublicKeyAsPemStr(pub_parsed, "")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(prv_parsed_pem)
	fmt.Println(pub_parsed_pem)

	// Check that the exported/imported keys match the original keys
	if prv_pem != prv_parsed_pem || pub_pem != pub_parsed_pem {
		fmt.Println("Failure: Export and Import did not result in same Keys")
	} else {
		fmt.Println("Success")
	}
}
