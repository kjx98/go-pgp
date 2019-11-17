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
		fmt.Printf("Entity %d: \n", idx)
		for idN, _ := range ee.Identities {
			fmt.Println("ID:", idN)
		}
	}

	fmt.Println("KeyRing test: END")
}
