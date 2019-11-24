package pgp

import (
	"bytes"
	"compress/gzip"
	_ "crypto/sha256"
	"errors"
	"fmt"
	"github.com/kjx98/openpgp"
	"github.com/kjx98/openpgp/armor"
	_ "golang.org/x/crypto/ripemd160"
	"io/ioutil"
)

func Decrypt(entity *openpgp.Entity, encrypted []byte) ([]byte, error) {
	// Decode message
	block, err := armor.Decode(bytes.NewReader(encrypted))
	if err != nil {
		return []byte{}, fmt.Errorf("Error decoding: %v", err)
	}
	if block.Type != "Message" {
		return []byte{}, errors.New("Invalid message type")
	}

	// Decrypt message
	entityList := openpgp.EntityList{entity}
	messageReader, err := openpgp.ReadMessage(block.Body, entityList, nil, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Error reading message: %v", err)
	}
	read, err := ioutil.ReadAll(messageReader.UnverifiedBody)
	if err != nil {
		return []byte{}, fmt.Errorf("Error reading unverified body: %v", err)
	}

	// Uncompress message
	reader := bytes.NewReader(read)
	uncompressed, err := gzip.NewReader(reader)
	if err != nil {
		return []byte{}, fmt.Errorf("Error initializing gzip reader: %v", err)
	}
	defer uncompressed.Close()

	out, err := ioutil.ReadAll(uncompressed)
	if err != nil {
		return []byte{}, err
	}

	// Return output - an unencoded, unencrypted, and uncompressed message
	return out, nil
}
