package main
import "fmt"

var salary = 10000

// Q11. Declare global variables age, salary, and id using individual declarations.
var age int
var id int
var salary1 int

// Q12. Rewrite Q11 using grouped variable declaration.
var(
	age1 int
	id1 int
	salary2 int
)

func main() {
	// Write solutions here
	
	/*
	-----------------------------------------------------
	1. BASIC VARIABLE DECLARATION
	-----------------------------------------------------
	*/

	// Q1. Declare an integer variable x with value 200 and print it.
	var x int=200
	fmt.Println(x)
	
	// Q2. Declare a variable y using type inference with value 4000 and print it.
	var y =4000
	fmt.Println(y)
	
	// Q3. Declare a variable a of type int, initialize it later with value 300, and print it.
	var a int
	a=300
	fmt.Println(a)
	
	/*Q4. Use short declaration to create a string variable slogan with value
    "Variables in Go - Day 3" and print it.*/
	slogan:="Variables in Go - Day 3"
	fmt.Println(slogan)

	/*
	-----------------------------------------------------
	2. STATIC TYPING
	-----------------------------------------------------
	*/

	// Q5. Declare a variable x as int and assign it a number.
	//     Then try assigning a string to x and observe the compiler error.
	// var x int =100
	// x = "Hello"
	
	// Q6. Declare a variable using := with a string value.
	//     Then try assigning an integer to the same variable.
	
	str:="Hello"
	str="world"
	fmt.Println(str)

	/*
	-----------------------------------------------------
	3. VARIABLE SCOPE
	-----------------------------------------------------
	*/
	
	// Q7. Declare a variable x inside main() and print it.
	
	x1:=60
	fmt.Println(x1)

	// Q8. Create a block {} inside main().
	//     Declare a variable y inside the block and print it inside the block.
	//     Try printing y outside the block.
	{
		y1:=20
		fmt.Println(y1)
	}
	// fmt.Println(y1)
	
	// Q9. Declare a global variable salary = 10000.
	//     Access and print it inside main().
		fmt.Println(salary)

	// Q10. Create a function varUse() that declares and prints a local variable.
	//      Try accessing that variable from main().
		varUse()

	/*
	-----------------------------------------------------
	4. GLOBAL AND GROUPED DECLARATIONS
	-----------------------------------------------------
	*/

	// Q13. Declare grouped variables inside main() and print them.
		
	var(
		age2 int =20
		id2 int =10
		salary3 int=30
		)
	fmt.Println(age2)
	fmt.Println(id2)
	fmt.Println(salary3)

	/*
	-----------------------------------------------------
	5. MULTIPLE DECLARATION AND ASSIGNMENT
	-----------------------------------------------------
	*/
		
	// Q15. Declare two variables b and c in a single line with values 10 and 20.
	//      Print both variables.
				
		b,c:=10,20
		fmt.Println(b,c)
		
	// Q16. Update variables b and c using a single assignment statement.

		b,c=30,40
		fmt.Println(b,c)
		
	// Q17. Swap two variables using multiple assignment
	//      without using a temporary variable.
	m,n:=5,10
	fmt.Println("Before Swap:",m,n)
	m,n=n,m
	fmt.Println("After Swap:",m,n)

	// Q18. Declare variables x and y using :=
	//      Then try redeclaring x and y again using := and observe the error.

	r,s:=1,2
	// r,s:=3,4
	fmt.Println(r,s)

	// Q19. Redeclare x using := by introducing one new variable.
	//      Print both variables.
	r,t:=5,6
	fmt.Println(r,t)

	/*
	-----------------------------------------------------
	6. ZERO VALUES
	-----------------------------------------------------
	*/

	// Q20. Declare variables of type int, float64, string, bool,
	//      slice, and pointer without initialization.
	//      Print their values and check which ones are nil.
	var zeroint int
	var zerofloat float64
	var zeroslice []int
	var zeroptr *int
	var zerosbool bool
	var zerosstring string
	fmt.Println("Zero int:",zeroint)
	fmt.Println("Zero float:",zerofloat)
	fmt.Println("Zero slice:",zeroslice)
	fmt.Println("Zero pointer:",zeroptr)
	fmt.Println("Zero bool:",zerosbool)
	fmt.Println("Zero string:",zerosstring)
}
func varUse(){
	c:=29
	fmt.Println(c)
}