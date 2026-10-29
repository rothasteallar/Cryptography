//Lab #2: Logical Operators

package main

import "fmt"

func main()  {

	var a, b int
	fmt.Print("Enter the first number(a): ")
	fmt.Scan(&a)
	
	fmt.Print("Enter the first number(b): ")
	fmt.Scan(&b)

	fmt.Println("Both positive: ", a >= 0 && b>= 0)
	fmt.Println("One greater than the other:  ", a > b || b>a)
	fmt.Println("Not equal: ", a != b)
	
}