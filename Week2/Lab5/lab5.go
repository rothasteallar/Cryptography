//Lab #5: Binary, Hex, and Base64 Encoding

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var text string
	fmt.Print("Enter text: ")
	fmt.Scan(&text)

	fmt.Println("\n\t==Display==\t")
	fmt.Println("String: ", text)
	fmt.Print("Binary: ")
	for i := 0; i < len(text); i++ {
		//pads the binary number with zeros until it has 8 digits
		fmt.Printf("%08b ", text[i])
	}
	fmt.Print("\nHex: ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%02x ", text[i])
	}
	fmt.Println("\nBase64: ", base64.StdEncoding.EncodeToString([]byte(text)))

}
