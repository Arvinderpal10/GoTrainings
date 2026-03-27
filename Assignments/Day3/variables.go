package main

//import "fmt"

//var salary = 10000

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
//
//	slice, and pointer without initialization.
//	Print their values and check which ones are nil.

//Q7 to Q10 solutions
/*func varUse() {
	a := 50
	fmt.Println("Inside varUse(), a =", a)
}
func main() {


	x := 10
	fmt.Println("Inside main(), x =", x)

	{
		y := 20
		fmt.Println("Inside block, y =", y)
	}

	fmt.Println("Global variable salary =", salary)
	varUse()
} */

//Q1 to Q6 solutions
/*func main() {
	// Write solutions here
	// A1
	var x int = 200
	fmt.Println(x)
	// A2
	y := 4000
	fmt.Println(y)
	// A3
	var a int
	a = 300
	fmt.Println(a)
	// A4
	slogan := "Variables in Go - Day 3"
	fmt.Println(slogan)
	// A5
	var x2 int = 10
	fmt.Println(x2)
	// x2 = "Hello" // compiler error
	// A6
	y1 := "Go Language"
	fmt.Println(y1)
	// y1 = 100 // compiler error

} */

// Q11 to Q14 solutions
/*var (
	age    int = 21
	salary int = 100000000
	id     int = 12345
)

func main() {
	var (
		name string = "Swaroop"
		city string = "Hyderabad"
		pin  int    = 500032
	)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("ID:", id)

	fmt.Println("Name:", name)
	fmt.Println("City:", city)
	fmt.Println("PIN:", pin)
}*/

// Q15 to Q17 solutions
/*func main() {
	b, c := 10, 20
	fmt.Println(" b = ", b, " c = ", c)
	b, c = 30, 40
	fmt.Println("Updated b = ", b, " Updated c = ", c)
	b, c = c, b
	fmt.Println("Swapped b = ", b, " Swapped c = ", c)

}*/

// Q18 and Q19 solutions
/*func main() {
	x, y := 100, 2000
	//x, y := 300, 4000 // error
	fmt.Println("x :=", x, " y :=", y)
	x, z := 500, 6000
	fmt.Println("x :=", x, " z :=", z)

}*/

// Q20 solution
/*func main() {
	var i int
	var f float64
	var s string
	var b bool
	var sl []int
	var p *int
	fmt.Println("int=", i)
	fmt.Println("flaot64=", f)
	fmt.Println("string=", s)
	fmt.Println("bool=", b)
	fmt.Println("slice=", sl)
	fmt.Println("pointer=", p)

}*/
