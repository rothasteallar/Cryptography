//Lab #4: Mini Calculator

package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}
func sub(a, b float64) float64 {
	return a - b
}
func mul(a, b float64) float64 {
	return a * b
}
func div(a, b float64) float64 {
	return a / b
}
func mod(a, b int) int {
	return a % b
}

func main() {

	var choice int
	var a, b float64

	for {
		fmt.Println("\n\t==Mini Calculator==\t")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")
		fmt.Printf("Choose: ")
		fmt.Scan(&choice)
		if choice == 6 {
			break
		}
		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice!")
			continue
		}

		fmt.Printf("Enter a: ")
		fmt.Scan(&a)
		fmt.Printf("Enter b: ")
		fmt.Scan(&b)

		if (choice == 4 || choice == 5) && b == 0 {
			fmt.Println("Error: division by zero")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Result: ", add(a, b))
		case 2:
			fmt.Println("Result: ", sub(a, b))

		case 3:
			fmt.Println("Result: ", mul(a, b))

		case 4:
			fmt.Println("Result: ", div(a, b))
		case 5:
			fmt.Println("Result: ", mod(int(a), int(b)))

		}

	}

}
