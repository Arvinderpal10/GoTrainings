package main

import (
	"fmt"
	"math"
	"unsafe"
)

//import "GoTrainings/exportunexport"

// Function returning a VALUE //5 answer
func returnValue() int {
	x := 10
	return x
}

// Function returning a POINTER to local variable //5 answer
func returnPointer() *int {
	y := 20
	return &y
}

func divide(a, b int) (int, int) { //as part of 9
	return a / b, a % b
}

var (
	age        int     = 22 //   part of 10
	salary     float64 = 45000.50
	name       string
	isEmployed bool
	count      int
)

func dynamicAppend() { //part of 12
	var nums []int

	for i := 0; i < 1000000; i++ {
		nums = append(nums, i)
	}
	fmt.Println("Length of dynamically appended slice:", len(nums))
}
func preAllocatedSlice() {
	nums := make([]int, 0, 1000000) //Always pre-allocate slice capacity when size is known.

	for i := 0; i < 1000000; i++ {
		nums = append(nums, i)
	}
	fmt.Println("Length of preallocated slice:", len(nums))
}

func createArray() { //part of 13
	var arr [1000000]int
	_ = arr
	fmt.Println("Created large array of size:", len(arr))
}
func main() {

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

		20. Declare variables with zero values and check if they equal nil for pointer, slice, and map types.


	*/
	//1
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807
	var e uint8 = 255
	var f uint16 = 65535
	var g uint32 = 4294967295
	var h uint64 = 18446744073709551615
	fmt.Println(a, b, c, d, e, f, g, h)
	//2
	var u uint8 = 255 //255 is the maximum value uint8 can store
	fmt.Println("Before overflow:", u)
	u = u + 1                         //Now Go tries to store 256 in a uint8.
	fmt.Println("After overflow:", u) //uint8 cannot store 256, so it wraps around to 0.
	//3
	var f32 float32 = 1.0 / 3.0         //float32 has limited precision That’s why you see 3433 instead of repeating 3s
	var f64 float64 = 1.0 / 3.0         //float64 is more precise than float32 when dealing with decimal values.
	fmt.Printf("float32: %.10f\n", f32) //Prints 10 digits after decimal
	fmt.Printf("float64: %.10f\n", f64)
	//4
	var bool1 bool = true
	var bool2 bool = false
	fmt.Println("Bool1:", bool1, "Bool2:", bool2)
	//5
	a1 := returnValue() //Escape analysis decides whether a variable is allocated on the stack or must escape to the heap.
	b1 := returnPointer()

	fmt.Println(a1)
	fmt.Println(*b1) //&Address-of operator (gives address), Dereference operator (gives value at address)
	//6 unsafe.Sizeof() returns How many bytes the variable x occupies in memory.
	fmt.Println("Size of int8:", unsafe.Sizeof(int8(0)))
	fmt.Println("Size of int16:", unsafe.Sizeof(int16(0)))
	fmt.Println("Size of int32:", unsafe.Sizeof(int32(0)))
	fmt.Println("Size of int64:", unsafe.Sizeof(int64(0)))

	fmt.Println("Size of uint8:", unsafe.Sizeof(uint8(0)))
	fmt.Println("Size of uint16:", unsafe.Sizeof(uint16(0)))
	fmt.Println("Size of uint32:", unsafe.Sizeof(uint32(0)))
	fmt.Println("Size of uint64:", unsafe.Sizeof(uint64(0)))

	fmt.Println("Size of float32:", unsafe.Sizeof(float32(0)))
	fmt.Println("Size of float64:", unsafe.Sizeof(float64(0)))

	fmt.Println("Size of bool:", unsafe.Sizeof(bool(true)))

	fmt.Println("Size of string:", unsafe.Sizeof(string("")))
	//7
	//fmt.Println("Age:", exportunexport.Age) //The Rule Starts with CAPITAL letter → Exported (public) Starts with small letter → Unexported (private to package)
	//fmt.Println("Salary:", exportunexport.salary) //compile time error because salary is unexported
	//8
	x := 10
	fmt.Println("Outer x:", x)
	{
		x := 20
		fmt.Println("Inner x:", x)
	}
	fmt.Println("Outer x:", x)
	//9
	quotient, _ := divide(10, 3) //divide returns two values We only care about the first one ,_ is used to ignore the second value

	fmt.Println("Quotient:", quotient) //_ is the blank identifier in Go. It means: “I know a value is returned, but I don’t want to use it.”
	//var _ int = 10  Go does NOT allow _ to behave like a normal variable.
	//_ := 10
	//10
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Name:", name)
	fmt.Println("Is Employed:", isEmployed)
	fmt.Println("Count:", count)

	//11
	for i := 0; i < 3; i++ {
		z := i * 2
		fmt.Println("Inside loop x:", z)
	}
	//fmt.Println("Outside loop x:", z) not accessible outside the loop
	//12
	dynamicAppend()
	preAllocatedSlice()
	//13
	createArray()
	//14
	var i int
	var f1 float64
	var str string
	var b2 bool
	var p *int
	var s []int
	var m map[string]int
	fmt.Println("int:", i)
	fmt.Println("float64:", f1)
	fmt.Println("string:", str)
	fmt.Println("bool:", b2)
	fmt.Println("pointer:", p)
	fmt.Println("slice:", s)
	fmt.Println("map:", m)
	//15
	var a2, b3, c3 int = 1, 2, 3
	fmt.Println("Before:", a2, b3, c3)
	a2, b3, c3 = 4, 5, 6
	fmt.Println("After:", a2, b3, c3)
	//16
	x1, y1 := 10, 20
	fmt.Println("Before Swap: x1 =", x1, "y1 =", y1)
	x1, y1 = y1, x1
	fmt.Println("After Swap: x1 =", x1, "y1 =", y1)
	//17
	x2, y2 := 30, 40
	fmt.Println("Before  x2 =", x2, "y2 =", y2)
	//x2, y2 := 50, 60 //compile time error: no new variables on left side of :=
	//fmt.Println("After x2 =", x2, "y2 =", y2)
	//18
	x2, y3 := 100, 200
	fmt.Println("x2 =", x2, "y3 =", y3)
	//19
	var maxInt32 int32 = math.MaxInt32
	var minInt32 int32 = math.MinInt32
	var maxUint32 uint32 = math.MaxUint32

	fmt.Println("Max Int32:", maxInt32)
	fmt.Println("Min Int32:", minInt32)
	fmt.Println("Max Uint32:", maxUint32)
	//20
	var p1 *int
	var s1 []int
	var m1 map[string]int

	// nil checks
	fmt.Println("Pointer is nil:", p1 == nil)
	fmt.Println("Slice is nil:", s1 == nil)
	fmt.Println("Map is nil:", m1 == nil)

}
