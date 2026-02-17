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
// 1. Do NOT change function signatures
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
// b) The type of value the function returns     >>>**correct
// c) The number of parameters
// d) The memory location of the function

// 2. (True/False)
// A function in Go can return more than one value.  >>>*True

// 3. (MCQ)
// Which keyword is used to return a value from a function?
// a) break
// b) return                           >>>*correct
// c) func
// d) yield

// 4. (True/False)
// A recursive function must have a base case.    >>>*True

// 5. (MCQ)
// What is a variadic function?
// a) A function with no parameters
// b) A function that returns multiple values
// c) A function that takes a variable number of arguments     >>>*correct
// d) A function that calls itself

// 6. (True/False)
// Variadic parameters are treated as slices inside the function.     >>>*True

// 7. (MCQ)
// What is an anonymous function?
// a) A function without a name         >>>*correct
// b) A function inside a package only
// c) A function that returns nothing
// d) A recursive function

// 8. (True/False)
// Functions in Go can be assigned to variables.     >>>*True

// 9. (MCQ)
// Which of the following is a correct function type?
// a) func(int) int       >>>*correct
// b) function(int) int
// c) fn(int) int
// d) int func(int)

// 10. (True/False)
// Only one return statement executes during a single function call.    >>>*False

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
	if n%2 == 0 {
		return true
	}
	return false
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
	for _, n := range nums {
		total += n
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
	fmt.Println(add(4, 5))            //Q1
	fmt.Println(isEven(10))           //Q2
	countdown(5)                      //Q3
	fmt.Println(factorial(5))         //Q4
	fmt.Println(sum(1, 2, 3))         //Q5
	result := apply(func(n int) int { //Q6
		return n * n
	}, 5)
	fmt.Println(result)

	func(a int, b int) { //Q7
		fmt.Println("Anonymous function Multilication", a*b)

	}(3, 5)

	fmt.Println(divide(10, 3)) //Q8

	add10 := makeAdder(10) //Q9
	fmt.Println(add10(5))
	add20 := makeAdder(20)
	fmt.Println(add20(5))

	nums := []int{1, 2, 3, 4} //Q10
	doubled := doubleSlice(nums)
	fmt.Println("Original:", nums)
	fmt.Println("Doubled:", doubled)

}
