package crack

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"time"

	// "crypto/md5"
	"fmt"
	// "go/scanner"
	"os"
)

// use uppercase if we want to export the function, lowercast for private function
func CrackSha512(targetSHA512 string) {
	startTime := time.Now()
	//Create verbose file
	logFile, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating verbose file: ", err)
		return
	}
	defer logFile.Close()

	// Starting information
	fmt.Fprintf(logFile, "[INFO] Starting SHA512 cracking...\n") //Fprintf = Formatted Print to a specified destination.
	fmt.Fprintf(logFile, "[INFO] Target SHA512: %s\n", targetSHA512)

	//JUST open file
	file, err := os.Open("wordlists/nord_vpn.txt") //os.Open needs two values
	if err != nil {
		fmt.Println("Error opening file: ", err)
		fmt.Fprintln(logFile, "[ERROR] Failed to open wordlist.")
		return
	}
	//CLOSE THE FILE
	defer file.Close()
	//defer delays execution until the surrounding function returns.
	fmt.Fprintln(logFile, "[INFO] Wordlist opened successfully.")

	//connect to the opened file
	scanner := bufio.NewScanner(file)

	isFound := false
	attempts := 0

	//GO doesn't have "while", only "for" type
	for scanner.Scan() {
		attempts++
		word := scanner.Text()

		hashWord := sha512.Sum512([]byte(word))        //this got [16]byte(if sha256, it's [32]byte)
		encodedHash := hex.EncodeToString(hashWord[:]) //[:] == [0:16]. encodetostring only accepts as Slice not a whole byte so

		if attempts <= 3 {
			fmt.Fprintf(logFile, "[TRY] Attempt %d\n", attempts)
		} else if attempts == 4 {
			fmt.Fprintln(logFile, "[TRY] .....")
		}

		if targetSHA512 == encodedHash {
			isFound = true
			fmt.Println("Password: ", word)
			fmt.Fprintf(logFile, "[SUCCESS] Password found at attempt %d\n", attempts)
			fmt.Fprintf(logFile, "[SUCCESS] Password is: %s\n", word)
			break
		}

	}

	// Check scanning error
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning file: ", err)
		fmt.Fprintln(logFile, "[ERROR] Failed while scanning wordlist.")
		return
	}

	// Report result
	if !isFound {
		fmt.Println("There's no corresponse password in this wordlist!!!")
		fmt.Fprintln(logFile, "[WARNING] No corresponding password found in this wordlist!!!!!")
	}

	// Execution summary
	duration := time.Since(startTime)
	fmt.Fprintf(logFile, "[INFO] Total attempts: %d\n", attempts)
	fmt.Fprintf(logFile, "[INFO] Execution time: %s\n", duration)
	fmt.Fprintln(logFile, "[INFO] Cracking process finished.")

}
