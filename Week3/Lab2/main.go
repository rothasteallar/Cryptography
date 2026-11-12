//Lab #1: Made the Password Cracking

package main

import (
	"lab2/utils/crack"
)

func main() {
	targetSHA1 := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

	crack.CrackSha1(targetSHA1)

}
