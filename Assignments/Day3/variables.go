package main

import "fmt"

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

//Q9
var salary = 10000

//Q11
var salaryy = 10000
var age = 22
var id = 1

//Q12
var (
	salary1 = 20000
	age1    = 23
	id1     = 2
)

func main() {
	// Write solutions here
	//Q1. Declare an integer variable x with value 200 and print it.
	var x int = 200
	fmt.Println(x)

	// Q2. Declare a variable y using type inference with value 4000 and print it.
	var y = 4000
	fmt.Println(y)

	// Q3. Declare a variable a of type int, initialize it later with value 300, and print it.
	var a int
	a = 300
	fmt.Println(a)

	// Q4. Use short declaration to create a string variable slogan with value
	//     "Variables in Go - Day 3" and print it.
	slogan := "Variables in Go - Day 3"
	fmt.Println(slogan)

	// Q5. Declare a variable x as int and assign it a number.
	//     Then try assigning a string to x and observe the compiler error.
	var xy int = 20
	fmt.Println(xy)
	//xy="Hello"  //cannot use "Hello" (untyped string constant) as int value in assignmentcompilerIncompatibleAssign

	// Q6. Declare a variable using := with a string value.
	//     Then try assigning an integer to the same variable.
	val := "welcome to go"
	fmt.Println(val)
	//val=20 //cannot use 20 (untyped int constant) as string value in assignmentcompilerIncompatibleAssign

	// Q7. Declare a variable x inside main() and print it.
	var xyz = 467
	fmt.Println(xyz)

	// Q8. Create a block {} inside main().
	//     Declare a variable y inside the block and print it inside the block.
	//     Try printing y outside the block.
	{
		var yz = 500
		fmt.Println(yz)
	}
	//fmt.Println(yz) //undefined: yzcompilerUndeclaredName

	// Q9. Declare a global variable salary = 10000.
	//     Access and print it inside main().
	fmt.Println(salary)

	//Q10
	varUse()
	// fmt.Println(value) //undefined: value

	// Q11. Declare global variables age, salary, and id using individual declarations.
	fmt.Println(salaryy)
	fmt.Println(age)
	fmt.Println(id)

	// Q12. Rewrite Q11 using grouped variable declaration.
	fmt.Println(salary1)
	fmt.Println(age1)
	fmt.Println(id1)

	// Q13. Declare grouped variables inside main() and print them.
	var (
		name    = "Malleswari"
		address = "Hyderabad"
	)
	fmt.Println(name)
	fmt.Println(address)

	// Q14. Write code to demonstrate that grouped declarations work
	//      both at global scope and function scope.

	// grouped declarations
	fmt.Println(salary1)
	fmt.Println(age1)
	fmt.Println(id1)
	// local declarations
	fmt.Println(name)
	fmt.Println(address)

	// Q15. Declare two variables b and c in a single line with values 10 and 20.
	//      Print both variables.
	var b, c = 10, 20
	fmt.Println(b, c)
	//fmt.Println(b)
	//fmt.Prinltn(c)

	// Q16. Update variables b and c using a single assignment statement.
	b, c = 20, 30
	fmt.Println(b, c)

	// Q17. Swap two variables using multiple assignment
	//      without using a temporary variable.
	var a1 = 10
	var b1 = 20
	fmt.Println("before swap", a1, b1)
	a1, b1 = b1, a1
	fmt.Println("After swap", a1, b1)

	// Q18. Declare variables x and y using :=
	//      Then try redeclaring x and y again using := and observe the error.
	x1 := 20
	y1 := 30
	fmt.Println(x1, y1)
	//x1:=30 //no new variables on left side of :=compiler
	//y1:=40 //no new variables on left side of :=compiler

	// Q19. Redeclare x using := by introducing one new variable.
	//      Print both variables.
	x2 := 4
	fmt.Println(x2)
	x2, y2 := 10, 20
	fmt.Println(x2, y2)

	// Q20. Declare variables of type int, float64, string, bool,
	//      slice, and pointer without initialization.
	//      Print their values and check which ones are nil.

	var i int
	var f float64
	var s string
	var b2 bool
	var s1 []int
	var p *int
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b2)
	fmt.Println(s1) //nil
	fmt.Println(p)  //nil
	fmt.Println(s1 == nil)
	fmt.Println(p == nil)

}

// Q10. Create a function varUse() that declares and prints a local variable.
//      Try accessing that variable from main().
func varUse() {
	var value = 100
	fmt.Println(value)
}
