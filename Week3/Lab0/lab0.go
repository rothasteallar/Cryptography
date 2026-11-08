//Hashing

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func toMd5(input1, input2 string) {
	var result string
	hash1 := md5.Sum([]byte(input1))
	hash2 := md5.Sum([]byte(input2))
	if hash1 == hash2 {
		result = "Match!"
	} else {
		result = "No Match!"
	}
	fmt.Println("Hash (MD5):")
	fmt.Println("Output A =", hex.EncodeToString(hash1[:]))
	fmt.Println("Output B =", hex.EncodeToString(hash2[:]))
	fmt.Println("=> ", result)

}
func toSha1(input1, input2 string) {
	var result string
	hash1 := sha1.Sum([]byte(input1))
	hash2 := sha1.Sum([]byte(input2))
	if hash1 == hash2 {
		result = "Match!"
	} else {
		result = "No Match!"
	}
	fmt.Println("Hash (SHA1):")
	fmt.Println("Output A =", hex.EncodeToString(hash1[:]))
	fmt.Println("Output B =", hex.EncodeToString(hash2[:]))
	fmt.Println("=> ", result)
}
func toSha256(input1, input2 string) {
	var result string
	hash1 := sha256.Sum256([]byte(input1))
	hash2 := sha256.Sum256([]byte(input2))
	if hash1 == hash2 {
		result = "Match!"
	} else {
		result = "No Match!"
	}
	fmt.Println("Hash (SHA256):")
	fmt.Println("Output A =", hex.EncodeToString(hash1[:]))
	fmt.Println("Output B =", hex.EncodeToString(hash2[:]))
	fmt.Println("=> ", result)
}
func toSha512(input1, input2 string) {
	var result string
	hash1 := sha512.Sum512([]byte(input1))
	hash2 := sha512.Sum512([]byte(input2))
	if hash1 == hash2 {
		result = "Match!"
	} else {
		result = "No Match!"
	}
	fmt.Println("Hash (SHA512):")
	fmt.Println("Output A =", hex.EncodeToString(hash1[:]))
	fmt.Println("Output B =", hex.EncodeToString(hash2[:]))
	fmt.Println("=> ", result)
}
func toSha3(input1, input2 string) {
	var result string
	hash1 := sha3.Sum256([]byte(input1))
	hash2 := sha3.Sum256([]byte(input2))
	if hash1 == hash2 {
		result = "Match!"
	} else {
		result = "No Match!"
	}
	fmt.Println("Hash (SHA3):")
	fmt.Println("Output A =", hex.EncodeToString(hash1[:]))
	fmt.Println("Output B =", hex.EncodeToString(hash2[:]))
	fmt.Println("=> ", result)
}

func main() {
	var input1 string
	var input2 string

	fmt.Print("Please input value 1: ")
	fmt.Scanln(&input1)
	fmt.Print("Please input value 2: ")
	fmt.Scanln(&input2)

	toMd5(input1, input2)
	toSha1(input1, input2)
	toSha256(input1, input2)
	toSha512(input1, input2)
	toSha3(input1, input2)

}
