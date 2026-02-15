/*
	1. Declare variables of types int8, int16, int32, int64, uint8, uint16, uint32, and uint64. Assign them their maximum values and print.

	2. Demonstrate integer overflow by assigning uint8 variable with 255 then adding 1. Print before and after values.

	3. Show precision differences between float32 and float64 by computing 1.0/3.0 with both types and printing with 10 decimal places.

	4. Create two bool variables and print.

	5. Write two functions - one returning a value, one returning a pointer to local variable. Build with -gcflags="-m" to analyze escape.

	6. Use unsafe.Sizeof() to print sizes of all integer types, float types, bool, and string.

	7. Create exported and unexported variables in a separate package. Try accessing both from main package and handle the error.

	8. Declare x := 10 in main. Inside a block, redeclare x := 20. Print both values to demonstrate shadowing.

	9. Use _ to ignore multiple return values from a function. Try to declare _ as a variable and assign to it.

	10. Create package-level variables using grouped declaration at global scope. Initialize some, leave others with zero values.

	11. Declare variable inside for loop: for i := 0; i < 3; i++ { x := i*2 }. Try to access x outside loop and handle error.

	12. Compare performance concept: Create slice with pre-allocated capacity vs dynamic appends (just write both versions).

	13. Create large array [1000000]int locally in function. Check if it escapes with -gcflags="-m".

	14. Declare variables without initialization: var a int; var b float64; var c string; var d bool; var e *int; var f []int; var g map[string]int. Print them to show zero values.

	15. Declare multiple variables in single line: var a, b, c int = 1, 2, 3. Then update them in single assignment: a, b, c = 4, 5, 6.

	16. Swap two variables without temporary: x, y := 10, 20; x, y = y, x. Print before and after.

	17. Try redeclaring variables with := when no new variables: x, y := 10, 20; x, y := 30, 40. Handle error.

	18. Redeclare with one new variable: x, z := 50, 60. Print both.

	19. Use math.MaxInt32, math.MinInt32, math.MaxUint32 constants. Assign them to variables and print.

	20. Declare variables with zero values and check if they equal nil for pointer, slice, and map types.*/

package main

import "fmt"

func main() {

	// 1. and 2.

	var a int8 = 127
	fmt.Println(a)
	var b int16 = 32767
	fmt.Println(b)
	var c int32 = 2147483647
	fmt.Println(c)
	var d int64 = 9223372036854775807
	fmt.Println(d)
	var e uint8 = 255
	fmt.Println(e)
	var f uint16 = 65535
	fmt.Println(f)
	var g uint32 = 4294967295
	fmt.Println(g)
	var h uint64 = 18446744073709551615
	fmt.Println(h)
	// 2.
	/*var x uint8 = 255
	  fmt.Println("Before overflow:", x)

	  x = x + 1
	  fmt.Println("After overflow:", x)*/

	// 3.
	/*var f32 float32 = 1.0 / 3.0
	var f64 float64 = 1.0 / 3.0
	fmt.Printf("float32 (1/3) = %.10f\n", f32)
	fmt.Printf("float64 (1/3) = %.10f\n", f64)*/

	//4.
	/*var b1 bool = true
	var b2 bool = false
	fmt.Println("Bool 1:", b1)
	fmt.Println("Bool 2:", b2)*/

	// 5.
	/*func returnValue() int {
	  	x := 44
	  	return x
	  }

	  func returnPointer() *int {
	  	y := 13
	  	return &y
	  }
	  	a := returnValue()
	  	b := returnPointer()

	  	fmt.Println(a)
	  	fmt.Println(*b)*/

	// 6.
	// import "unsafe"
	/*var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	var f32 float32
	var f64 float64

	var b bool
	var s string

	fmt.Println("int8:", unsafe.Sizeof(i8))
	fmt.Println("int16:", unsafe.Sizeof(i16))
	fmt.Println("int32:", unsafe.Sizeof(i32))
	fmt.Println("int64:", unsafe.Sizeof(i64))

	fmt.Println("uint8:", unsafe.Sizeof(ui8))
	fmt.Println("uint16:", unsafe.Sizeof(ui16))
	fmt.Println("uint32:", unsafe.Sizeof(ui32))
	fmt.Println("uint64:", unsafe.Sizeof(ui64))

	fmt.Println("float32:", unsafe.Sizeof(f32))
	fmt.Println("float64:", unsafe.Sizeof(f64))

	fmt.Println("bool:", unsafe.Sizeof(b))
	fmt.Println("string:", unsafe.Sizeof(s))*/

	// 7.
	// import mypkg and create a seperate package with exported and unexported variables better to create a folder named mypkg and
	//inside that create a file named mypkg.go and declare the variables there and then import it here and try to access both variables
	/*fmt.Println(mypkg.ExportedVar)*/
	// Uncommenting this will cause a compile-time error
	// fmt.Println(mypkg.unexportedVar)

	// 8.
	/*x := 10
	fmt.Println("Outer x:", x)

	{
		x := 20
		fmt.Println("Inner x:", x)
	}

	fmt.Println("Outer x after block:", x)*/

	// 9.
	//func getValues() (int, int, int) {
	//return 1, 2, 3
	/*a, _, c := getValues()
	fmt.Println(a, c)*/

	// 10.
	//var (
	//a int    = 10
	//b string = "hello"
	//c float64
	//d bool
	//)
	//code here
	/*fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)*/

	// 11.
	/*for i := 0; i < 3; i++ {
		x := i * 2
		fmt.Println("Inside loop x:", x)

	}
	// undefined x */

	// 12.
	/*func dynamicAppend() {
		var s []int
		for i := 0; i < 1_000_000; i++ {
			s = append(s, i)
		}
		fmt.Println(len(s))
	}

	func preAllocatedAppend() {
		s := make([]int, 0, 1_000_000)
		for i := 0; i < 1_000_000; i++ {
			s = append(s, i)
		}
		fmt.Println(len(s))
	dynamicAppend()
	preAllocatedAppend()*/

	// 13.
	/*func createArray() {
		var arr [1_000_000]int
		_ = arr
	}
	createArray()*/

	// 14.
	/*var a int
	var b float64
	var c string
	var d bool
	var e *int
	var f []int
	var g map[string]int
	fmt.Println("int:", a)
	fmt.Println("float64:", b)
	fmt.Println("string:", c)
	fmt.Println("bool:", d)
	fmt.Println("pointer:", e)
	fmt.Println("slice:", f)
	fmt.Println("map:", g)*/

	// 15.
	/*var a, b, c int = 1, 2, 3
	fmt.Println("Before:", a, b, c)
	a, b, c = 4, 5, 6
	fmt.Println("After:", a, b, c)*/

	// 16.
	/*x, y := 10, 20
	fmt.Println("Before swap:", x, y)
	x, y = y, x
	fmt.Println("After swap:", x, y)*/

	// 17.
	/*x, y := 10, 20
	fmt.Println(x, y)

	//x, y := 30, 40
	//fmt.Println(x, y)*/

	// 18.
	/*x := 10

	x, z := 50, 60
	fmt.Println(x, z)*/

	// 19.
	/*// import "math"
	var maxInt32 int32 = math.MaxInt32
	var minInt32 int32 = math.MinInt32
	var maxUint32 uint32 = math.MaxUint32
	fmt.Println("Max int32:", maxInt32)
	fmt.Println("Min int32:", minInt32)
	fmt.Println("Max uint32:", maxUint32)*/

	// 20.
	/*var p *int
	var s []int
	var m map[string]int

	fmt.Println("Pointer is nil:", p == nil)
	fmt.Println("Slice is nil:", s == nil)
	fmt.Println("Map is nil:", m == nil)*/

}
