// Task 1: Write a Go program to demonstrate the use of assignment operators. The program should take two integer inputs and perform various assignment operations
// such as =, +=, -=, *=, /=, and %=. Display the result after each operation


package main

import "fmt"

func main()  {
	fmt.Println("Hello World!")

	var a, b int

	fmt.Print("Enter the first number (a): ")
	fmt.Scan(&a)
	fmt.Print("Enter the second number (b): ")
	fmt.Scan(&b)

	result := a
	fmt.Println("= : ", result)

	result += b
	fmt.Println("+= : ", result)

	result -= b
	fmt.Println("-= : ", result)

	result *= b
	fmt.Println("*= : ", result)

	result /= b
	fmt.Println("/= : ", result)

	result %= b
	fmt.Println("%= : ", result)

}
