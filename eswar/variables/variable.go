// Q1. Declare an integer variable x with value 200 and print it.

// package main
// import "fmt"
// func main() {
// 	var x int = 200
// 	fmt.Println(x)
// }



// Q2. Declare a variable y using type inference with value 4000 and print it.

// package main
// import "fmt"
// func main() {
	
// 		var y = 4000
// 		fmt.Println(y)

// }


//Q3. Declare a variable a of type int, initialize it later with value 300, and print it.

// package main
// import "fmt"
// func main() {
// 	var a int
// 	a = 300
// 	fmt.Println(a)
// }


// Q4. Use short declaration to create a string variable slogan with value
//     "Variables in Go - Day 3" and print it.

// package main
// import "fmt"
// func main () {
// 	slogan := "variables in go - day 3"
// 	fmt.Println(slogan)
// }

// Q5. Declare a variable x as int and assign it a number.
//     Then try assigning a string to x and observe the compiler error.

// package main
// import "fmt"
// func main() {
// 	var x int = 10
// 	x = "Hello"
// 	fmt.Println(x)
// }

// Q6. Declare a variable using := with a string value.
//     Then try assigning an integer to the same variable.

// package main
// import "fmt"
// func main () {
// 	name := "golang"
// 	name = 100
// 	fmt.Println(name)
// }

// Q7. Declare a variable x inside main() and print it.

// package main
// import "fmt"
// func main() {
// 	var x int = 1206
// 	fmt.Println(x)
// }

// Q8. Create a block {} inside main().
//     Declare a variable y inside the block and print it inside the block.
//     Try printing y outside the block.

// package main
// import "fmt"
// func main() {
// 	{
// 		var y int = 1206
// 		fmt.Println(y)
// 	}
// 	fmt.Println(y) 
// }

// Q9. Declare a global variable salary = 10000.
//     Access and print it inside main().

// package main
// import "fmt"
// var salary int = 10000
// func main() {
// 	fmt.Println(salary)
// }

// Q10. Create a function varUse() that declares and prints a local variable.
//      Try accessing that variable from main().

// package main
// import "fmt"
// func varuse() {
// 	x := 1206
// 	fmt.Println(x)
// }
// func main() {
// 	varuse()
// 	fmt.Println(x)
// }

// Q11. Declare global variables age, salary, and id using individual declarations.

// package main
// import "fmt"
// var age int = 23
// var salary int = 50000
// var id int = 1206
// func main (){
// 	fmt.Println(age,salary,id)
// }


// Q12. Rewrite Q11 using grouped variable declaration.

// package main

// import "fmt"

// var (
//     age    int     = 25
//     salary float64 = 50000
//     id     int     = 101
// )

// func main() {
//     fmt.Println(age, salary, id)
// }

// Q13. Declare grouped variables inside main() and print them.
// package main

// import "fmt"

// func main() {
//     var (
//         a int    = 10
//         b string = "Go"
//         c float64 = 3.14
//     )

//     fmt.Println(a, b, c)
// }


// Q14. Write code to demonstrate that grouped declarations work
//      both at global scope and function scope.

// package main

// import "fmt"


// var (
//     age    int     = 30
//     salary float64 = 60000
// )

// func main() {

//     var (
//         name string = "Go"
//         id   int    = 12
//     )

//     fmt.Println(age, salary)
//     fmt.Println(name, id)
// }


// Q15. Declare two variables b and c in a single line with values 10 and 20.
//      Print both variables.

// package main
// import "fmt"
// func main(){
// 	var b,c = 10,20
// 	fmt.Println(b,c)
// }

// Q16. Update variables b and c using a single assignment statement.

// package main
// import "fmt"
// func main () {
// 	b,c := 12,6
// 	b,c = 20,22
// 	fmt.Println(b,c)
// }

// Q17. Swap two variables using multiple assignment
//      without using a temporary variable.

// package main
// import "fmt"
// func main(){
// 	a,b := 12,6
// 	a,b = b,a
// 	fmt.Println(a,b)
// }

// Q18. Declare variables x and y using :=
//      Then try redeclaring x and y again using := and observe the error.

// package main
// import "fmt"
// func main() {
// 	x,y := 10,20
// 	x,y := 10,20
// 	fmt.Println(x,y)
// }

 //Q19. Redeclare x using := by introducing one new variable.
//      Print both variables.

// package main
// import "fmt"
// func main() {
// 	x:= 10
// 	x,y := 20,30
// 	fmt.Println(x,y)
// }


// Q20. Declare variables of type int, float64, string, bool,
//      slice, and pointer without initialization.
//      Print their values and check which ones are nil.

// package main

// import "fmt"

// func main() {
//     var i int
//     var f float64
//     var s string
//     var b bool
//     var sl []int
//     var p *int

//     fmt.Println("int:", i)
//     fmt.Println("float64:", f)
//     fmt.Println("string:", s)
//     fmt.Println("bool:", b)
//     fmt.Println("slice:", sl)
//     fmt.Println("pointer:", p)

//     fmt.Println("slice is nil?", sl == nil)
//     fmt.Println("pointer is nil?", p == nil)
// }


//Day -2

// Q1.Declare variables of types int8, int16, int32, int64, uint8, uint16, uint32, and uint64. Assign them their maximum values and print.
// package  main
// import "fmt"
// func main() {
//     var i8 int8 = 127
//     var i16 int16 = 32767
//     var i32 int32 = 2147483647
//     var i64 int64 = 9223372036854775807

    
//     var u8 uint8 = 255
//     var u16 uint16 = 65535
//     var u32 uint32 = 4294967295
//     var u64 uint64 = 18446744073709551615

//     fmt.Println("int8  :", i8)
//     fmt.Println("int16 :", i16)
//     fmt.Println("int32 :", i32)
//     fmt.Println("int64 :", i64)

//     fmt.Println("uint8  :", u8)
//     fmt.Println("uint16 :", u16)
//     fmt.Println("uint32 :", u32)
//     fmt.Println("uint64 :", u64)

// }

// 2. Demonstrate integer overflow by assigning uint8 variable with 255 then adding 1. Print before and after values.

// package main

// import "fmt"
// func main (){
//     var x uint8 = 255
//     fmt.Println("before:",x)
//      x = x + 1

//     fmt.Println("After :", x)
// }

// 3. Show precision differences between float32 and float64 by computing 1.0/3.0 with both types and printing with 10 decimal places.

// package main

// import "fmt"

// func main() {
//     var f32 float32 = 1.0 / 3.0
//     var f64 float64 = 1.0 / 3.0

//     fmt.Printf("float32: %.10f\n", f32)
//     fmt.Printf("float64: %.10f\n", f64)
// }



// 4. Create two bool variables and print.

// package main
// import "fmt"
// func main(){
//     var a bool = true
//     var b bool = false
//     fmt.Println(a)
//     fmt.Println(b)
// }
	
// 5. Write two functions - one returning a value, one returning a pointer to local variable. Build with -gcflags="-m" to analyze escape.

// package main

// import "fmt"

// func pointerFunc() *int {
//     x := 10
//     return &x
// }

// func main() {
//     p := pointerFunc()
//     fmt.Println(*p)
// }

// 6. Use unsafe.Sizeof() to print sizes of all integer types, float types, bool, and string.
// package main

// import (
//     "fmt"
//     "unsafe"
// )

// func main() {
//     var (
//         i8  int8
//         i16 int16
//         i32 int32
//         i64 int64
//         i   int

//         u8  uint8
//         u16 uint16
//         u32 uint32
//         u64 uint64
//         u   uint

//         f32 float32
//         f64 float64

//         b bool
//         s string
//     )

//     fmt.Println("int8:", unsafe.Sizeof(i8))
//     fmt.Println("int16:", unsafe.Sizeof(i16))
//     fmt.Println("int32:", unsafe.Sizeof(i32))
//     fmt.Println("int64:", unsafe.Sizeof(i64))
//     fmt.Println("int:", unsafe.Sizeof(i))

//     fmt.Println("uint8:", unsafe.Sizeof(u8))
//     fmt.Println("uint16:", unsafe.Sizeof(u16))
//     fmt.Println("uint32:", unsafe.Sizeof(u32))
//     fmt.Println("uint64:", unsafe.Sizeof(u64))
//     fmt.Println("uint:", unsafe.Sizeof(u))

//     fmt.Println("float32:", unsafe.Sizeof(f32))
//     fmt.Println("float64:", unsafe.Sizeof(f64))

//     fmt.Println("bool:", unsafe.Sizeof(b))
//     fmt.Println("string:", unsafe.Sizeof(s))
// }

// 7. Create exported and unexported variables in a separate package. Try accessing both from main package and handle the error.

// package main

// import "fmt"

// var ExportedVar = "I am exported"
// var unexportedVar = "I am unexported"

// func main() {
//     fmt.Println(ExportedVar)
//     fmt.Println(unexportedVar)
// }

// 8. Declare x := 10 in main. Inside a block, redeclare x := 20. Print both values to demonstrate shadowing.
// package main
// import "fmt"
// func main(){
//     x:= 10
//     fmt.Println("outer x:",x)
//     {
//         x:= 20
//         fmt.Println("inner x:", x)
//     }
//     fmt.Println("outer x again:", x)
// }

// 9. Use _ to ignore multiple return values from a function. Try to declare _ as a variable and assign to it.
// package main

// import "fmt"

// func giveTwo() (int, int) {
// 	return 10, 20
// }

// func main() {
// 	a, _ := giveTwo()
// 	fmt.Println(a)
// }

// 10. Create package-level variables using grouped declaration at global scope. Initialize some, leave others with zero values.
// package main

// import "fmt"

// var (
// 	x int = 5
// 	y int
// 	z string = "Go"
// 	w bool
// )

// func main() {
// 	fmt.Println(x, y, z, w)
// }

// 11. Declare variable inside for loop: for i := 0; i < 3; i++ { x := i*2 }. Try to access x outside loop and handle error.

// package main

// import "fmt"

// func main() {
// 	var x int

// 	for i := 0; i < 3; i++ {
// 		x = i * 2
// 		fmt.Println("inside:", x)
// 	}

// 	fmt.Println("outside:", x)
// }

// 12. Compare performance concept: Create slice with pre-allocated capacity vs dynamic appends (just write both versions).
//pre-located capacity
// package main

// import "fmt"

// func main() {
// 	n := 5

// 	s := make([]int, 0, n) 
// 	for i := 0; i < n; i++ {
// 		s = append(s, i)
// 	}

// 	fmt.Println(s)
// }

// dynamic  appends
// package main

// import "fmt"

// func main() {
// 	n := 10

// 	var s []int 

// 	for i := 0; i < n; i++ {
// 		s = append(s, i)
// 	}

// 	fmt.Println(s)
// }


// 13. Create large array [1000000]int locally in function. Check if it escapes with -gcflags="-m".

// package main

// func bigArray() {
// 	var arr [1000000]int
// 	arr[0] = 42
// }

// func main() {
// 	bigArray()
// }

// 14. Declare variables without initialization: var a int; var b float64; var c string; var d bool; var e *int; var f []int; var g map[string]int. Print them to show zero values.

// package main

// import "fmt"

// func main() {
// 	var a int
// 	var b float64
// 	var c string
// 	var d bool
// 	var e *int
// 	var f []int
// 	var g map[string]int

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// 	fmt.Println(f)
// 	fmt.Println(g)
// }

// 15. Declare multiple variables in single line: var a, b, c int = 1, 2, 3. Then update them in single assignment: a, b, c = 4, 5, 6.
// package main

// import "fmt"

// func main() {
// 	var a, b, c int = 1, 2, 3
// 	fmt.Println(a, b, c)

// 	a, b, c = 4, 5, 6
// 	fmt.Println(a, b, c)
// }

// 16. Swap two variables without temporary: x, y := 10, 20; x, y = y, x. Print before and after.
// package main

// import "fmt"

// func main() {
// 	x, y := 10, 20

// 	fmt.Println("before:", x, y)

// 	x, y = y, x

// 	fmt.Println("after:", x, y)
// }

// 17. Try redeclaring variables with := when no new variables: x, y := 10, 20; x, y := 30, 40. Handle error.

// package main

// func main() {
// 	x, y := 10, 20
// 	x, y := 30, 40
// }

// 18. Redeclare with one new variable: x, z := 50, 60. Print both.
// package main

// import "fmt"

// func main() {
// 	x := 10

// 	x, z := 50, 60

// 	fmt.Println(x, z)
// }

// 19. Use math.MaxInt32, math.MinInt32, math.MaxUint32 constants. Assign them to variables and print.

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var a int32 = math.MaxInt32
// 	var b int32 = math.MinInt32
// 	var c uint32 = math.MaxUint32

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// 20. Declare variables with zero values and check if they equal nil for pointer, slice, and map types.

package main

import "fmt"

func main() {
	var p *int
	var s []int
	var m map[string]int

	fmt.Println("pointer nil?", p == nil)
	fmt.Println("slice nil?", s == nil)
	fmt.Println("map nil?", m == nil)
}

