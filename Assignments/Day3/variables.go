package main

/*
=====================================================
Day 3 – Golang Training - Coding Questions
=====================================================
*/

/*
-----------------------------------------------------
1. BASIC VARIABLE DECLARATION
-----------------------------------------------------
*/

// Q1. Declare an integer variable x with value 200 and print it.

// Q2. Declare a variable y using type inference with value 4000 and print it.

// Q3. Declare a variable a of type int, initialize it later with value 300, and print it.

// Q4. Use short declaration to create a string variable slogan with value
//     "Variables in Go - Day 3" and print it.

/*
-----------------------------------------------------
2. STATIC TYPING
-----------------------------------------------------
*/

// Q5. Declare a variable x as int and assign it a number.
//     Then try assigning a string to x and observe the compiler error.

// Q6. Declare a variable using := with a string value.
//     Then try assigning an integer to the same variable.

/*
-----------------------------------------------------
3. VARIABLE SCOPE
-----------------------------------------------------
*/

// Q7. Declare a variable x inside main() and print it.

// Q8. Create a block {} inside main().
//     Declare a variable y inside the block and print it inside the block.
//     Try printing y outside the block.

// Q9. Declare a global variable salary = 10000.
//     Access and print it inside main().

// Q10. Create a function varUse() that declares and prints a local variable.
//      Try accessing that variable from main().

/*
-----------------------------------------------------
4. GLOBAL AND GROUPED DECLARATIONS
-----------------------------------------------------
*/

// Q11. Declare global variables age, salary, and id using individual declarations.

// Q12. Rewrite Q11 using grouped variable declaration.

// Q13. Declare grouped variables inside main() and print them.

// Q14. Write code to demonstrate that grouped declarations work
//      both at global scope and function scope.

/*
-----------------------------------------------------
5. MULTIPLE DECLARATION AND ASSIGNMENT
-----------------------------------------------------
*/

// Q15. Declare two variables b and c in a single line with values 10 and 20.
//      Print both variables.

// Q16. Update variables b and c using a single assignment statement.

// Q17. Swap two variables using multiple assignment
//      without using a temporary variable.

// Q18. Declare variables x and y using :=
//      Then try redeclaring x and y again using := and observe the error.

// Q19. Redeclare x using := by introducing one new variable.
//      Print both variables.

/*
-----------------------------------------------------
6. ZERO VALUES
-----------------------------------------------------
*/

// Q20. Declare variables of type int, float64, string, bool,
//      slice, and pointer without initialization.
//      Print their values and check which ones are nil.

import (
	"fmt"
)
var salary=10000
func varUse(){
		var x8=8746
		fmt.Println(x8)
	}
	var(
		age= 22
		salary2=25000
		id=13
	)
func main() {
	fmt.Println(salary)
	{
		z:=15
		fmt.Println(z)
	}
	
	varUse()
	fmt.Println(id,age,salary2)
	//fmt.Println(z)

	var(
		age= 22
		salary2=25000
		id=13
	)
	fmt.Println(age,salary2,id)
	var x =200
	fmt.Println(x)
	y:=4000
	fmt.Println(y)
	var g int=300
	fmt.Println(g)
	var slogan="variables in go -day3"
	fmt.Println(slogan)
	/*var x1=05
	var x1="DOB"
	fmt.Println(x1)
	y1:="hello"
	y1=100
	fmt.Println(y1)*/
	var b,c=10,20
	fmt.Println(b,c)
	b,c=30,40
	fmt.Println("Before Swapping:",b,c)
	b,c=c,b
	fmt.Println("After swapping:",b,c)
	/*x9:=58
	y9:=96
	fmt.Println(x9,y9)
	x9:=96
	y9:=58
	fmt.Println(x9,y9)*/
	x9:=53
	y9:=35
	fmt.Println(x9,y9)
	var i int
    var f float64
    var s string
    var b2 bool
    var sl []int
    var p *int
	fmt.Printf("Zero value of int: %d\n", i)
	fmt.Printf("Zero value of float64: %f\n", f)
	fmt.Printf("Zero value of string: %s\n", s)
	fmt.Printf("Zero value of bool: %t\n", b2)
	fmt.Printf("Zero value of slice: %v\n", sl)
	fmt.Printf("Zero value of pointer: %v\n", p)

}
