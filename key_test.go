package pgp

import (
	"fmt"
	"testing"
)

func TestKeyRing(t *testing.T) {
	fmt.Println("KeyRing test: START")
	entity, err := getKeyRing()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("List key entity.")
	for idx, ee := range entity {
		bitLen, _ := ee.PrimaryKey.BitLength()
		fmt.Printf("Entity %d: %s(%d) PUB: %s\n", idx,
			ee.PrimaryKey.PubKeyAlgo, bitLen, ee.PrimaryKey.KeyIdString())
		for idN, _ := range ee.Identities {
			fmt.Println("ID:", idN)
		}
	}

	fmt.Println("KeyRing test: END")
}
