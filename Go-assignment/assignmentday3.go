package main

import "fmt"
//1. Q1. Declare an integer variable x with value 200 and print it.

// func main() {
//     var x int = 200
//     fmt.Println(x)
// }

// Q2. Declare a variable y using type inference with value 4000 and print it.
// func main() {
// 	x := 200
// 	fmt.Printf("Value of x is %v\n",x)
// 	fmt.Printf("Type of x is %T\n",x)
// }


// Q3. Declare a variable a of type int, initialize it later with value 300, and print it.
// func main() {
// 	var a int
// 	a = 300
// 	fmt.Println(a)
// }

// Q4. Use short declaration to create a string variable slogan with value
// func main() {
// 	slogan := "Hello, Jaya!"
// 	fmt.Println(slogan)
// }

// Q5. Declare a variable x as int and assign it a number. Then, reassign x to a string value and print both the value and type of x. What happens?
//  func main() {
//  	var x int
//  	x = "jaya"
//  	fmt.Printf("Value of x is %v\n",x)
//  	fmt.Printf("Type of x is %T\n",x)
// }

// Q6. Declare a variable using := with a string value.Then try assigning an integer to the same variable.
// func main() {
// 	x := "hi"
// 	x = 20
// 	fmt.Printf("Value of x is %v\n",x)
// 	fmt.Printf("Type of x is %T\n",x)
// }

// Q7. Declare a variable x inside main() and print it.
// func main() {
// 	{
// 		x := 100
// 		fmt.Println(x)
// 	}
// }


// Q8. Create a block {} inside main() and declare a variable inside that block. Try to access that variable outside the block and observe what happens.
// func main() {
// 	{
// 		y :=20

// 	}
// 	fmt.Println(y)
// }


// Q9. Declare a global variable salary = 10000. Access and print it inside main().
// var salary int = 10000
// func main() {
// 	fmt.Println(salary)
// }

// Q10. Create a function varUse() that declares and prints a local variable.Try accessing that variable from main().
// func varUse() {
// 	x := 50
// 	//y := 100 *error 
// 	fmt.Println(x)
// }

// func main() {
// 	varUse()
// 	y := 200
// 	//fmt.Println(x) *error
// 	fmt.Println(y)
// }

// Q11. Declare global variables age, salary, and id using individual declarations
// var age int = 22
// var salary int = 15000
// var id int = 217
// func main() {
// 	fmt.Println(age, salary, id)
// }

// Q12. Rewrite Q11 using grouped variable declaration.
// var (
// 	age = 22
// 	salary = 15000
// 	id = 217
// )
// func main() {
// 	fmt.Println(age, salary, id)
// }

// Q13. Declare grouped variables inside main() and print them.
// func main() {
// 	var (
// 		age = 22
// 		salary = 15000
// 		id = 217
// 	)
// 	fmt.Println(age, salary, id)
// }

//Q14. Write code to demonstrate that grouped declarations work, both at global scope and function scope.
// var (
// 	age    int = 22
// 	salary int = 15000
// 	id     int = 217
// )
// func main() {
// 	var (
// 		name   string = "jaya"
// 		city  string = "hyd"
// 	)
// 	fmt.Println(age, salary, id)
// 	fmt.Println(name, city)
// }

// Q15. Declare two variables b and c in a single line with values 10 and 20.Print both variables.
// func main() {
// 	var b,c = 10,20
// 	fmt.Println(b,c)
// }


// Q16. Update variables b and c using a single assignment statement.
// func main() {
// 	var b,c=10,20
// 	b,c=50,100
// 	fmt.Println(b,c)
// }

// Q17. Swap two variables using multiple assignment, without using a temporary variable.
// func main() {
// 	var b,c=10,20
// 	fmt.Println(b,c)
// 	b,c=c,b
// 	fmt.Println(b,c)
// }

// Q18. Declare variables x and y using := ,Then try redeclaring x and y again using := and observe the error.
// func main() {
// 	x,y := 1,2
// 	fmt.Println(x,y)
// 	// x,y := 3,4  *error
// 	fmt.Println(x,y)
// }	


// Q19. Redeclare x using := by introducing one new variable.Print both variables.

// func main() {
// 	x:= 1
// 	fmt.Println(x)
// 	x,y := 3,4
// 	fmt.Println(x,y)
// }


// Q20. Declare variables of type int, float64, string, bool, slice, and pointer without initializing them.Print their zero values. Print their values and check which ones are nil.

func main() {

	var i int
	var f float64
	var s string
	var b bool
	var sl []int
	var p *int

	fmt.Println(i, f, s, b, sl, p, sl == nil, p == nil)
}
