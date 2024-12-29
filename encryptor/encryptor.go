package encryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func Encrypt(plaintext, key string) (string, error) {
	// Debug: Start encryption
	fmt.Println("Encrypt function called with plaintext:", plaintext)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("Error creating AES cipher:", err)
		return "", err
	}

	plainBytes := []byte(plaintext)
	padding := aes.BlockSize - len(plainBytes)%aes.BlockSize
	paddedPlaintext := append(plainBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)

	ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("Error generating IV:", err)
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlaintext)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("Encryption successful, ciphertext:", encoded)
	return encoded, nil
}

func Decrypt(ciphertext, key string) (string, error) {
	// Debug: Start decryption
	fmt.Println("Decrypt function called with ciphertext:", ciphertext)

	encryptedBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		fmt.Println("Error decoding base64 ciphertext:", err)
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("Error creating AES cipher:", err)
		return "", err
	}

	if len(encryptedBytes) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := encryptedBytes[:aes.BlockSize]
	encryptedBytes = encryptedBytes[aes.BlockSize:]

	if len(encryptedBytes)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptedBytes, encryptedBytes)

	padding := int(encryptedBytes[len(encryptedBytes)-1])
	if padding > aes.BlockSize || padding > len(encryptedBytes) {
		return "", errors.New("invalid padding")
	}

	result := string(encryptedBytes[:len(encryptedBytes)-padding])
	fmt.Println("Decryption successful, plaintext:", result)
	return result, nil
}
