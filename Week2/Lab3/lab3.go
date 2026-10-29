//Lab #3: Bitwise and Assignment Opr.

package main

import "fmt"

//both differ = 1
func myXor(a, b int)int {
	return a ^ b
}

//reverse it
func myNOT(a, b int)int {
	return ^a
}

func myOR(a, b int)int {
	return a | b
}
func myAND(a, b int)int {
	return a &  b
}

//a * (2 square b), or binary go to left by the b times and ADD 0 at end by b times
func myLeftShift(a, b int)int{
	return a << b
}

//a % (2 square b), or binary go to right by the b times and ADD 0 at end by b times
func myRightShift(a, b int)int{
	return a >> b
}

func main()  {
	var a, b int 

	fmt.Print("Enter the first number(a): ")
	fmt.Scan(&a)
	
	fmt.Print("Enter the first number(b): ")
	fmt.Scan(&b)

	fmt.Printf("XOR: %d\n",myXor(a, b) )
	fmt.Printf("NOT: %d\n",myNOT(a, b) )
	fmt.Printf("OR: %d\n",myOR(a, b) )
	fmt.Printf("AND: %d\n",myAND(a, b) )

	fmt.Printf("Left shift: %d\n", myLeftShift(a, b))	
	fmt.Printf("Right shift: %d\n",myRightShift(a, b) )
	

	


}