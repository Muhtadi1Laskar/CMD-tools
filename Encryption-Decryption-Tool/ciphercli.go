package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func encrypt(plainText, secretKey string) (string, error) {
	hashedKey := sha256.Sum256([]byte(secretKey))

	block, err := aes.NewCipher([]byte(hashedKey[:]))
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %v", err)
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return hex.EncodeToString(cipherText), nil

}

func decrypt(cipherTextHex, secretKey string) (string, error) {
	// Derive the fixed-size key using SHA-256
	hashedKey := sha256.Sum256([]byte(secretKey))

	ciphertext, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return "", fmt.Errorf("invalid hex-encoded ciphertext: %v", err)
	}

	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt data
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %v", err)
	}

	return string(plaintext), nil
}

func main() {
	data := flag.String("data", "", "the data")
	key := flag.String("key", "", "the secret key")
	cipherType := flag.String("type", "", "the cipher type")

	flag.Parse()

	if *cipherType != "encrypt" && *cipherType != "decrypt" {
		fmt.Println("Invalid cipher type! Use 'encrypt' or 'decrypt'.")
		os.Exit(1)
	}
	var output string
	var err error

	switch *cipherType {
	case "encrypt":
		output, err = encrypt(*data, *key)
	case "decrypt":
		output, err = decrypt(*data, *key)
	default:
		return 
		
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	fmt.Println(output)
}