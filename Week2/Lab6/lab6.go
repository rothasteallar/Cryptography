package main

import "fmt"

func xorEncrypt(text string, key byte) string {
	result := []byte(text)

	for i := 0; i < len(result); i++ {
		result[i] = result[i] ^ key
	}

	return string(result)
}

func main() {
	var text string
	var key byte

	fmt.Print("\nEnter plaintext: ")
	fmt.Scanln(&text)

	fmt.Print("Enter key (single character): ")
	fmt.Scanf("%c", &key)

	encrypted := xorEncrypt(text, key)
	fmt.Printf("Ciphertext: %x\n", encrypted)

	decrypted := xorEncrypt(encrypted, key)
	fmt.Println("Decrypted text:", decrypted)
}
