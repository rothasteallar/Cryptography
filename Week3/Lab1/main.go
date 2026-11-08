//Lab #1: Made the Password Cracking

package main

import (
	"lab1/utils/crack"
)

func main() {
	targetMD5 := "6a85dfd77d9cb35770c9dc6728d73d3f"

	crack.CrackMd5(targetMD5)

}
