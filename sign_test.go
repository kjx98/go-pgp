package pgp

import (
	"fmt"
	"testing"
)

func TestSignature(t *testing.T) {
	fmt.Println("Signature test: START")
	entity, err := GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created private key entity.")

	signature, err := Sign(entity, []byte(TestMessage))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created signature of test message with private key entity.")

	publicKeyEntity, err := GetEntity([]byte(TestPublicKey), []byte{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created public key entity.")

	err = Verify(publicKeyEntity, []byte(TestMessage), signature)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Signature verified using public key entity.")
	fmt.Println("Signature test: END")
}
