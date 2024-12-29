package main

import (
	"encryptcli/encryptor"
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.String("mode", "", "Mode: 'encrypt' or 'decrypt'")
	text := flag.String("text", "", "Text to encrypt/decrypt")
	key := flag.String("key", "", "32-character encryption key")
	flag.Parse()

	fmt.Println("Mode:", *mode)
	fmt.Println("Text:", *text)
	fmt.Println("Key:", *key)

	if len(*key) != 32 {
		fmt.Println("Error: Key must be 32 characters long.")
		os.Exit(1)
	}
	if *mode == "" || *text == "" {
		fmt.Println("Error: Both mode and text are required.")
		os.Exit(1)
	}

	switch *mode {
	case "encrypt":
		encryptedText, err := encryptor.Encrypt(*text, *key)
		if err != nil {
			fmt.Println("Encryption failed:", err)
			os.Exit(1)
		}
		fmt.Println("Encrypted text:", encryptedText)
	case "decrypt":
		decryptedText, err := encryptor.Decrypt(*text, *key)
		if err != nil {
			fmt.Println("Decryption failed:", err)
			os.Exit(1)
		}
		fmt.Println("Decrypted text:", decryptedText)
	default:
		fmt.Println("Invalid mode. Use 'encrypt' or 'decrypt'.")
		os.Exit(1)
	}
}
