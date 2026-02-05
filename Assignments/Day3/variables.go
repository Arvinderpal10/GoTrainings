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
var salary int = 10000 //9

func varUse() { //10
	var localVar string = "I am a local variable"
	fmt.Println(localVar)
}

var age int = 30 //11
var salary2 float64 = 50000.50
var id string = "EMP123"

// 12 Grouped global variable declaration
var (
	age2    int     = 25
	salary3 float64 = 60000
	id2     string  = "EMP102"
)

func main() {
	// Write solutions here
	//1
	var x int = 200
	fmt.Println(x)
	//2
	y := 4000
	fmt.Println(y)
	//3
	var a int
	a = 300
	fmt.Println(a)
	//4
	slogan := "Variables in Go - Day 3"
	fmt.Println(slogan)
	//5
	//var x int = 10
	//x = "Hello"
	//fmt.Println(x) //compile time error
	//6
	//p := "prasad"
	//p = 123 //compile time error
	//fmt.Println(p)
	//7
	var x1 int = 50
	fmt.Println(x1)
	//8
	{
		y1 := 100
		fmt.Println(y1)
	}
	//fmt.Println(y1) //compile time error because y1 is not accessible outside the block
	//9
	fmt.Println(salary)
	//10varUse()
	varUse()
	//13
	var (
		q = 444
		r = "sai"
		s = 44.4
	)
	fmt.Println(q, r, s)
	//14
	fmt.Println(age2, salary3, id2)
	//15
	b, c := 10, 20
	fmt.Println(b, c)
	//16
	b, c = 30, 40
	fmt.Println(b, c)
	//17
	b, c = c, b
	fmt.Println(b, c)
	//18
	x2, y2 := 5, 10
	fmt.Println(x2, y2)
	//x2, y2 := 100, 200 //compile time error
	//fmt.Println(x2, y2)
	//19
	x2, z := 100, 300
	fmt.Println(x2, z)
	//20
	var i int
	var f float64
	var str string
	var b1 bool
	var s1 []int
	var p1 *int
	fmt.Println(i, f, str, b1, s1, p1)

}
