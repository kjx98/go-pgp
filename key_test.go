package pgp

import (
	"fmt"
	"testing"
)

func TestKeyRing(t *testing.T) {
	fmt.Println("KeyRing test: START")
	fmt.Println("test pubring")
	entity, err := getKeyRing("")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("List key entity.")
	for idx, ee := range entity {
		fmt.Printf("Entity %d: PUB: %X\n", idx, ee.PrimaryKey.KeyId)
		for idN, _ := range ee.Identities {
			fmt.Println("ID:", idN)
		}
	}
	fmt.Println("test secring")
	entity, err = getKeyRing("secring.gpg")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("List key entity.")
	for idx, ee := range entity {
		fmt.Printf("Entity %d: PUB: %X\n", idx, ee.PrimaryKey.KeyId)
		for idN, _ := range ee.Identities {
			fmt.Println("ID:", idN)
		}
		if ee.PrivateKey != nil {
			fmt.Println("Private Encrypted:", ee.PrivateKey.Encrypted)
		}
	}

	fmt.Println("KeyRing test: END")
}
