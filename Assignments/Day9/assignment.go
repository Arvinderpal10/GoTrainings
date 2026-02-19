package main

import "fmt"

///////////////////////////////////////////////////////////////
// GO ASSIGNMENT – BASIC LEVEL
// Topics Covered:
// Functions, Function Types, Recursion,
// Variadic Functions, Anonymous Functions,
// Multiple Return Values
//
// Instructions:
// 1. Do NOT change function signatures.
// 2. Do NOT use global variables.
// 3. Write your answers for theory in comments.
// 4. Write test cases inside main().
///////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////
// SECTION A – THEORY
///////////////////////////////////////////////////////////////

// 1. (MCQ)
// What does the return type of a function indicate?
// a) The name of the function
// b) The type of value the function returns
// c) The number of parameters
// d) The memory location of the function

// Ans:b. The type of value the function returns

// 2. (True/False)
// A function in Go can return more than one value.
// Ans:True

// 3. (MCQ)
// Which keyword is used to return a value from a function?
// a) break
// b) return
// c) func
// d) yield
// Ans:b. return

// 4. (True/False)
// A recursive function must have a base case.
// Ans:True

// 5. (MCQ)
// What is a variadic function?
// a) A function with no parameters
// b) A function that returns multiple values
// c) A function that takes a variable number of arguments
// d) A function that calls itself
// Ans: c. A function that takes a variable number of arguments


// 6. (True/False)
// Variadic parameters are treated as slices inside the function.
// Ans:True

// 7. (MCQ)
// What is an anonymous function?
// a) A function without a name
// b) A function inside a package only
// c) A function that returns nothing
// d) A recursive function

// Ans: a. A function without a name

// 8. (True/False)
// Functions in Go can be assigned to variables.
// Ans:True

// 9. (MCQ)
// Which of the following is a correct function type?
// a) func(int) int
// b) function(int) int
// c) fn(int) int
// d) int func(int)
// Ans: a. func(int)int

// 10. (True/False)
// Only one return statement executes during a single function call.
// Ans:True


///////////////////////////////////////////////////////////////
// SECTION B – CODING (Basic Level)
///////////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////////
// 1. Simple Add Function
// Write a function add that takes two integers
// and returns their sum.
// /////////////////////////////////////////////////////////////
func add(a int, b int) int {
	return a + b
}

// /////////////////////////////////////////////////////////////
// 2. Even or Odd
// Write a function isEven that takes an integer
// and returns true if the number is even,
// otherwise false.
// /////////////////////////////////////////////////////////////
func isEven(n int) bool {
	return n%2 == 0
}

// /////////////////////////////////////////////////////////////
// 3. Recursive Countdown
// Write a recursive function countdown(n int)
// that prints numbers from n down to 1.
// Stop when n becomes 0.
// /////////////////////////////////////////////////////////////
func countdown(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	countdown(n - 1)
}

// /////////////////////////////////////////////////////////////
// 4. Recursive Factorial
// Write a recursive function factorial(n int)
// that returns the factorial of n.
// (Assume n >= 0)
// /////////////////////////////////////////////////////////////
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// /////////////////////////////////////////////////////////////
// 5. Variadic Sum
// Write a variadic function sum that accepts
// multiple integers and returns their total.
// /////////////////////////////////////////////////////////////
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// /////////////////////////////////////////////////////////////
// 6. Function as Parameter
// Write a function apply that takes:
// - a function of type func(int) int
// - an integer value
// It should return the result of calling the function
// with the integer.
// /////////////////////////////////////////////////////////////
func apply(f func(int) int, x int) int {
	return f(x)
}

///////////////////////////////////////////////////////////////
// 7. Anonymous Function
// Inside main(), create an anonymous function
// that multiplies two integers and print the result
// of multiplying 3 and 4.
///////////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////////
// 8. Multiple Return Values
// Write a function divide that takes two integers
// and returns quotient and remainder.
// /////////////////////////////////////////////////////////////
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

// /////////////////////////////////////////////////////////////
// 9. Closure
// Write a function makeAdder(x int) that returns
// a function. The returned function should take
// one integer y and return x + y.
// /////////////////////////////////////////////////////////////
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// /////////////////////////////////////////////////////////////
// 10. Simple Map Function
// Write a function doubleSlice that takes a slice
// of integers and returns a new slice where each
// element is doubled.
// /////////////////////////////////////////////////////////////
func doubleSlice(slice []int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = v * 2
	}
	return result
}

// /////////////////////////////////////////////////////////////
// MAIN FUNCTION
// Add test cases here to check your functions.
// /////////////////////////////////////////////////////////////
func main() {

	
	fmt.Println("1. Add:", add(3, 5))

	fmt.Println("2. isEven(4):", isEven(4))
	fmt.Println("   isEven(7):", isEven(7))

	fmt.Println("\n3. Countdown from 5:")
	countdown(5)

	fmt.Println("\n4. Factorial of 5:", factorial(5))

	fmt.Println("\n5. Variadic Sum:", sum(1, 2, 3, 4))

	// helper function for apply
	square := func(n int) int {
		return n * n
	}
	fmt.Println("\n6. Apply square to 6:", apply(square, 6))

	// 7. Anonymous Function
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("\n7. Anonymous multiply 3*4:", multiply(3, 4))

	// 8. Divide
	q, r := divide(10, 3)
	fmt.Println("\n8. Divide 10 by 3 → Quotient:", q, "Remainder:", r)

	// 9. Closure
	addFive := makeAdder(5)
	fmt.Println("\n9. Closure addFive(10):", addFive(10))

	// 10. doubleSlice
	nums := []int{1, 2, 3, 4}
	fmt.Println("\n10. Double Slice:", doubleSlice(nums))
}