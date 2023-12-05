package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"io"
)

func EncryptMessage(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("could not create new cipher: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)
	data, err := serializeToBinary(cipherText)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DecryptMessage(key []byte, cipherText []byte, data interface{}) (string, error) {
	// First, deserialize data into the expected format
	err := deserializeFromBinary(cipherText, data)
	if err != nil {
		return "", fmt.Errorf("could not deserialize data: %v", err)
	}

	// Assert that data is of type []byte after deserialization
	decryptedData, ok := data.([]byte)
	if !ok {
		return "", fmt.Errorf("data is not of type []byte")
	}

	// Create a new cipher using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	// Check if the decrypted data has the minimum block size
	if len(decryptedData) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	// Extract the IV from the decrypted data
	iv := decryptedData[:aes.BlockSize]
	decryptedData = decryptedData[aes.BlockSize:]

	// Decrypt the data
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decryptedData, decryptedData)

	return string(decryptedData), nil
}

func deserializeFromBinary(in []byte, target interface{}) error {
	buf := bytes.NewBuffer(in)
	decoder := gob.NewDecoder(buf)
	return decoder.Decode(target)
}

func serializeToBinary(in any) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(in)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
