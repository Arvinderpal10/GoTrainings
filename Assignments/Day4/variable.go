package main

import (
	"fmt"
	"math"
	"unsafe"
	"day4/test"
)

//Q10
var (
	p int
	q float32=234.35
	r string ="Welcome to Go"
	s bool
)
func main() {

	// 1. Declare variables of types int8, int16, int32, int64, uint8, uint16, uint32, and uint64. Assign them their maximum values and print.

	var a int8 = math.MaxInt8
	fmt.Println(a)

	var b int16 = math.MaxInt16
	fmt.Println(b)

	var c int32 = math.MaxInt32
	fmt.Println(c)

	var d int64 = math.MaxInt64
	fmt.Println(d)

	var x uint8 = math.MaxUint8
	fmt.Println(x)

	var y uint16 = math.MaxUint16
	fmt.Println(y)

	var z uint32 = math.MaxUint32
	fmt.Println(z)

	var xy uint64 = math.MaxUint64
	fmt.Println(xy)

	// 2. Demonstrate integer overflow by assigning uint8 variable with 255 then adding 1. Print before and after values.

	var i uint8 = 255
	fmt.Println("before adding 1", i) //255
	i += 1
	fmt.Println("After adding 1", i) //0

	// 3. Show precision differences between float32 and float64 by computing 1.0/3.0 with both types and printing with 10 decimal places.

	var f32 float32 = 1.0 / 3.0
	var f64 float64 = 1.0 / 3.0
	fmt.Printf("float32 (10 decimal places): %.10f\n", f32) //0.3333333433
	fmt.Printf("float64 (10 decimal places): %.10f\n", f64) //0.3333333333

	// 4. Create two bool variables and print.

	var b1 bool = true
	var b2 bool = false
	fmt.Println(b1, b2)

	// 5. Write two functions - one returning a value, one returning a pointer to local variable. Build with -gcflags="-m" to analyze escape.

	val:=returnValue()
	val1:=returnPointer()
	fmt.Println(val)
	fmt.Println(*val1)

	// 6. Use unsafe.Sizeof() to print sizes of all integer types, float types, bool, and string.

	var i1 int
	var i8 int8 
	var i16 int16
	var i322 int32
	var i64 int64

	var u uint
	var u8 uint8
	var u16 uint16
	var u322 uint32
	var u64 uint64

	fmt.Println("int size :=" , unsafe.Sizeof(i1))
	fmt.Println("int8 size :=" , unsafe.Sizeof(i8))
	fmt.Println("int16 size :=" , unsafe.Sizeof(i16))
	fmt.Println("int32 size :=" , unsafe.Sizeof(i322))
	fmt.Println("int64 size :=" , unsafe.Sizeof(i64))
	fmt.Println("uint size :=" , unsafe.Sizeof(u))
	fmt.Println("uint8 size :=" , unsafe.Sizeof(u8))
	fmt.Println("uint16 size :=" , unsafe.Sizeof(u16))
	fmt.Println("uint32 size :=" , unsafe.Sizeof(u322))
	fmt.Println("uint64 size :=" , unsafe.Sizeof(u64))

	var f322 float32
	var f644 float64
	fmt.Println("float32 size:=" , unsafe.Sizeof(f322))
	fmt.Println("float64 size:=" , unsafe.Sizeof(f644))

	var boo bool
	var str string
	fmt.Println("bool size:=", unsafe.Sizeof(boo))
	fmt.Println("String size:=", unsafe.Sizeof(str))

	// 7. Create exported and unexported variables in a separate package. Try accessing both from main package and handle the error.

	fmt.Println(test.ExportedVar)
	//fmt.Println(test.unExportedVar) //undefined: test.unExportedVar
	
	// 8. Declare x := 10 in main. Inside a block, redeclare x := 20. Print both values to demonstrate shadowing.

	x1:=10
	{
		x1:=20 //creates new varible , not modifying the old x1
		fmt.Println(x1)
	}
	fmt.Println(x1)


	// 9. Use _ to ignore multiple return values from a function. Try to declare _ as a variable and assign to it.

	_, remainder:=divide(5,2)
	fmt.Println(remainder)
	
	// 10. Create package-level variables using grouped declaration at global scope. Initialize some, leave others with zero values.

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)

	// 11. Declare variable inside for loop: for i := 0; i < 3; i++ { x := i*2 }. Try to access x outside loop and handle error.

	for i:=0;i<3;i++ {
		fx:=i*2
		fmt.Println(fx)
	}
	//fmt.Println(fx) // undefined: fx

	// 12. Compare performance concept: Create slice with pre-allocated capacity vs dynamic appends (just write both versions).

	//Slice without pre-allocation
	var s1 []int
	for i:=0;i<5;i++ {
		s1=append(s1,i)
	}
	fmt.Println(s1)

	//Slice with pre-allocation
	s:= make([]int, 0, 5)  
	for i:=0;i<5;i++ {
		s=append(s,i)
	}
	fmt.Println(s)

	// 13. Create large array [1000000]int locally in function. Check if it escapes with -gcflags="-m".

	createArray()
	
	// 14. Declare variables without initialization: var a int; var b float64; var c string; var d bool; var e *int; var f []int; var g map[string]int. Print them to show zero values.

	var a1 int
	var b3 float64
	var c1 string
	var d1 bool
	var e *int
	var f []int
	var g map[string]int

	fmt.Println(a1) //0
	fmt.Println(b3) //0
	fmt.Println(c1) //empty
	fmt.Println(d1) //false
	fmt.Println(e) //<nil>
	fmt.Println(f) //[]
	fmt.Println(g) //map[]

	// 15. Declare multiple variables in single line: var a, b, c int = 1, 2, 3. Then update them in single assignment: a, b, c = 4, 5, 6.

	var aa,bb,cc=1,2,3
	fmt.Println(aa,bb,cc)
	aa,bb,cc=4,5,6
	fmt.Println(aa,bb,cc)

	// 16. Swap two variables without temporary: x, y := 10, 20; x, y = y, x. Print before and after.
	xx,yy:=10,20
	fmt.Println("before swap",xx,yy)
	xx,yy=yy,xx
	fmt.Println("after swap",xx,yy)

	// 17. Try redeclaring variables with := when no new variables: x, y := 10, 20; x, y := 30, 40. Handle error.

	x2,y2 :=10,20
	// x2,y2:=20,30  //no new variables on left side of :=
	fmt.Println(x2,y2)

	// 18. Redeclare with one new variable: x, z := 50, 60. Print both.

	x2,z2:=50,60
	fmt.Println(x2,z2)

	// 19. Use math.MaxInt32, math.MinInt32, math.MaxUint32 constants. Assign them to variables and print.

	var i32 int32=math.MaxInt32
	var i3 int32=math.MinInt32
	var u32 uint32=math.MaxUint32
	fmt.Println(i32)
	fmt.Println(i3)
	fmt.Println(u32)

	// 20. Declare variables with zero values and check if they equal nil for pointer, slice, and map types.

	var ptr *int
	var slice []int
	var m map[string] int
	fmt.Println(ptr==nil)   //true
	fmt.Println(slice==nil) //true
	fmt.Println(m==nil)     //true
}

// 5. Write two functions - one returning a value, one returning a pointer to local variable. Build with -gcflags="-m" to analyze escape.
//function returning value
func returnValue() int {
	x:=20
	return x
}

//function returning pointer
func returnPointer() *int {
	y:=20 // moved to heap: y
	return &y
}

// 9. Use _ to ignore multiple return values from a function. Try to declare _ as a variable and assign to it.
func divide(a,b int) (int,int){
	return a/b, a%b
}

//Q13
func createArray(){
	var arr [1000000]int
	arr[0]=100 //moved to heap: arr
}