package main

import (
	"fmt"
	"math"
)

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

	var (
		i8  int8   = math.MaxInt8
		i16 int16  = math.MaxInt16
		i32 int32  = math.MaxInt32
		i64 int64  = math.MaxInt64
		u8  uint8  = math.MaxUint8
		u16 uint16 = math.MaxUint16
		u32 uint32 = math.MaxUint32
		u64 uint64 = math.MaxUint64
	)
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n", i8, i16, i32, i64)
	fmt.Printf("uint8: %d, uint16: %d, uint32: %d, uint64: %d\n", u8, u16, u32, u64)
	var ov uint8 = 255
	fmt.Printf(" Before: %d", ov)
	ov = ov + 1
	fmt.Printf(" After: %d\n", ov)
	var f32 float32 = 1.0 / 3.0
	var f64 float64 = 1.0 / 3.0
	fmt.Printf("float32: %.10f, float64: %.10f\n", f32, f64)
	var bool1 bool = true
	var bool2 bool = false
	fmt.Printf("bool1: %t, bool2: %t\n", bool1, bool2)
	val := returnValue()
	ptr := returnPointer()
	fmt.Printf("Value: %d, Pointer: %d\n", val, *ptr)
	x := 10
	fmt.Println("x before block:", x)
	{
		x := 20
		fmt.Println("x inside block:", x)
	}
	fmt.Println("x after block:", x)
	a, _, c := multiReturn()
	fmt.Println("a =", a, "c =", c)
	fmt.Println("Global variables:", gA, gB, gC, gD)
	

	for i := 0; i < 3; i++ {
		x := i * 2
		fmt.Println("Inside loop x =", x)
	}


	
	s1 := make([]int, 0)
	for i := 0; i < 5; i++ {
		s1 = append(s1, i)
	}
	fmt.Println("Dynamic slice:", s1)

	s2 := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		s2 = append(s2, i)
	}
	fmt.Println("Pre-allocated slice:", s2)

	createLargeArray()


	var zva int
	var zvb float64
	var zvc string
	var zvd bool
	var zve *int
	var zvf []int
	var zvg map[string]int

	fmt.Println("int:", zva)
	fmt.Println("float64:", zvb)
	fmt.Println("string:", zvc)
	fmt.Println("bool:", zvd)
	fmt.Println("pointer:", zve)
	fmt.Println("slice:", zvf)
	fmt.Println("map:", zvg)

	var p, q, r int = 1, 2, 3
	fmt.Println("Before:", p, q, r)

	p, q, r = 4, 5, 6
	fmt.Println("After:", p, q, r)


	x, y := 10, 20
	fmt.Println("Before swap:", x, y)

	x, y = y, x
	fmt.Println("After swap:", x, y)

	
	a1, b1 := 10, 20
	fmt.Println("a1 =", a1, "b1 =", b1)

	
	a1, b1 = 30, 40
	fmt.Println("Updated a1 =", a1, "b1 =", b1)

	m := 50
	m, z := 60, 70 
	fmt.Println("m =", m, "z =", z)

	maxInt32 := math.MaxInt32
	minInt32 := math.MinInt32
	maxUint32 := uint32(math.MaxUint32)

	fmt.Println("MaxInt32:", maxInt32)
	fmt.Println("MinInt32:", minInt32)
	fmt.Println("MaxUint32:", maxUint32)

	
	var nilPtr *int
	var nilSl []int
	var nilMp map[string]int

	fmt.Println("ptr == nil:", nilPtr == nil)
	fmt.Println("slice == nil:", nilSl == nil)
	fmt.Println("map == nil:", nilMp == nil)
}

func createLargeArray() {
	var arr [1000000]int
	arr[0] = 99
	fmt.Println("Large array created, arr[0] =", arr[0])

	
}

func multiReturn() (any, any, any) {
	return 42, "ignored", true
}

func returnValue() int {
	return 42
}

func returnPointer() *int {
	v := 100
	return &v
}

var (
	gA int    = 10
	gB string = "Hello"
	gC bool
	gD float64
)
