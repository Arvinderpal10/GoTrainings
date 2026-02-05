// Q1. Declare an integer variable x with value 200 and print it.

// var x int
// x = 200
// fmt.print(x)

// Q2. Declare a variable y using type inference with value 4000 and print it.

// func main() {
// 	x := 4000
// 	fmt.Println(x)
// }

// Q3. Declare a variable a of type int, initialize it later with value 300, and print it.

// func main() {
// 	var a int
// 	a = 300
// 	fmt.Println(a)
// }

// Q4. Use short declaration to create a string variable slogan with value
//     "Variables in Go - Day 3" and print it.

// func main(){
// 	slogan:="Variables in Go - Day 3"
// 	fmt.Println(slogan)
// }

// Q5. Declare a variable x as int and assign it a number.
//     Then try assigning a string to x and observe the compiler error.

// func main() {
// 	var x int = 10
// 	fmt.Print("hello"+x)
// }

// Q6. Declare a variable using := with a string value.
//     Then try assigning an integer to the same variable.

// func main(){
// 	s:="this is string"
// 	fmt.Println(s+"5")
// }

// Q7. Declare a variable x inside main() and print it.

// func main(){
// 	x:=10
// 	fmt.Println(x)
// }

// Q8. Create a block {} inside main().
//     Declare a variable y inside the block and print it inside the block.
//     Try printing y outside the block.

// func main(){
// 	{
// 		y:=10
// 		fmt.Println(y)
// 	}
// 	fmt.Println(y)
// }

// Q9. Declare a global variable salary = 10000.
//     Access and print it inside main().

// var salary int = 10000

// func man(){
// 	fmt.Println(salary)
// }

// Q10. Create a function varUse() that declares and prints a local variable.
//      Try accessing that variable from main().

// func varUse() {
// 	x := 10
// 	fmt.Println(x)
// }

// func main() {
// 	varUse()
// }

// Q11. Declare global variables age, salary, and id using individual declarations.

// var age int = 22
// var salary int = 30000
// var id int = 1

// func main(){

// }

// Q12. Rewrite Q11 using grouped variable declaration.
// var(
// 	age int = 22
// 	salary int = 30000
// 	id int = 1
// )
// func main(){}

// Q13. Declare grouped variables inside main() and print them.

// func main(){
// 	var(
// 		age int = 22
// 		name string = "ram"
// 	)

// 	fmt.Println(age,name)
// }

// Q14. Write code to demonstrate that grouped declarations work
//      both at global scope and function scope.

// var (
// 	a int     = 10
// 	b string  = "Hello"
// 	c float64 = 3.14
// )

// func main() {
// 	var (
// 		x = 100
// 		y = "Go"
// 		z = true
// 	)

// 	fmt.Println("Global variables:")
// 	fmt.Println(a, b, c)

// 	fmt.Println("Local variables:")
// 	fmt.Println(x, y, z)
// }

// Q15. Declare two variables b and c in a single line with values 10 and 20.
//      Print both variables.

// func main(){
// 	b,c:=10,20
// 	fmt.Println(b,c)
// }

// Q16. Update variables b and c using a single assignment statement.

// func main() {
// 	b, c := 10, 20

// 	b, c = 30, 40

// 	fmt.Println(b, c)
// }

// Q17. Swap two variables using multiple assignment
//      without using a temporary variable.

// func main(){
// 	var(
// 		a = 10
// 		b = 20
// 	)
// 	a = a+b
// 	b = a-b
// 	a = a-b

// 	fmt.Println(a,b)
// }

// Q18. Declare variables x and y using :=
//      Then try redeclaring x and y again using := and observe the error.

// func main() {
// 	x, y := 5, 10

// 	x, y := 20, 30

// 	fmt.Println(x,y)
// }

// Q19. Redeclare x using := by introducing one new variable.
//      Print both variables.

// func main() {
// 	x := 10

// 	x, y := 20, 30

// 	fmt.Println(x, y)
// }

// 1. Declare variables of types int8, int16, int32, int64, uint8, uint16, uint32, and uint64. Assign them their maximum values and print.

// func main() {
// 	var (
// 		a int8   = math.MaxInt8
// 		b int16  = math.MaxInt16
// 		c int32  = math.MaxInt32
// 		d int64  = math.MaxInt64
// 		e uint8  = math.MaxUint8
// 		f uint16 = math.MaxUint16
// 		g uint32 = math.MaxUint32
// 		h uint64 = math.MaxUint64
// 	)
// 	fmt.Println("a :", a)
// 	fmt.Println("b :", b)
// 	fmt.Println("c :", c)
// 	fmt.Println("d :", d)
// 	fmt.Println("e :", e)
// 	fmt.Println("f :", f)
// 	fmt.Println("g :", g)
// 	fmt.Println("h :", h)

// }

// 2. Demonstrate integer overflow by assigning uint8 variable with 255 then adding 1. Print before and after values.

// func main() {
// 	var a uint8 = 255
// 	fmt.Println(a)

// 	a = a + 1
// 	fmt.Println(a)
// }

// 3. Show precision differences between float32 and float64 by computing 1.0/3.0 with both types and printing with 10 decimal places.

// func main() {
// 	var f32 float32 = 1.0 / 3.0
// 	var f64 float64 = 1.0 / 3.0

// 	fmt.Printf("float32: %.10f\n", f32)
// 	fmt.Printf("float64: %.10f\n", f64)
// }

// 4. Create two bool variables and print.

// func main(){
// 	var(
// 		a bool = true
// 		b bool = false
// 	)
// 	fmt.Println(a,b)
// }

// 5. Write two functions - one returning a value, one returning a pointer to local variable. Build with -gcflags="-m" to analyze escape.

// func ValueReturn()int{
// 	x:=10
// 	return x
// }

// func PointerReturn()*int{
// 	y:=12
// 	return &y
// }

// func main(){
// 	a:=ValueReturn()
// 	b:=PointerReturn()
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// 6. Use unsafe.Sizeof() to print sizes of all integer types, float types, bool, and string.

// func main() {
// 	fmt.Println("Integer types:")
// 	fmt.Println("int8   :", unsafe.Sizeof(int8(0)))
// 	fmt.Println("int16  :", unsafe.Sizeof(int16(0)))
// 	fmt.Println("int32  :", unsafe.Sizeof(int32(0)))
// 	fmt.Println("int64  :", unsafe.Sizeof(int64(0)))
// 	fmt.Println("uint8  :", unsafe.Sizeof(uint8(0)))
// 	fmt.Println("uint16 :", unsafe.Sizeof(uint16(0)))
// 	fmt.Println("uint32 :", unsafe.Sizeof(uint32(0)))
// 	fmt.Println("uint64 :", unsafe.Sizeof(uint64(0)))
// 	fmt.Println("int    :", unsafe.Sizeof(int(0)))
// 	fmt.Println("uint   :", unsafe.Sizeof(uint(0)))

// 	fmt.Println("\nFloat types:")
// 	fmt.Println("float32:", unsafe.Sizeof(float32(0)))
// 	fmt.Println("float64:", unsafe.Sizeof(float64(0)))

// 	fmt.Println("\nOther types:")
// 	fmt.Println("bool   :", unsafe.Sizeof(true))
// 	fmt.Println("string :", unsafe.Sizeof(""))
// }

// 8. Declare x := 10 in main. Inside a block, redeclare x := 20. Print both values to demonstrate shadowing.

// func main() {
// 	x := 10
// 	{
// 		x := 20
// 		fmt.Println(x)
// 	}
// 	fmt.Println(x)
// }

// 9. Use _ to ignore multiple return values from a function. Try to declare _ as a variable and assign to it.

// func getValues() (int, string, bool) {
// 	return 10, "Go", true
// }

// func main() {
// 	a, _, c := getValues()
// 	fmt.Println(a, c)

// 	_ = 100
// 	_ = "ignored"
// }

// 10. Create package-level variables using grouped declaration at global scope. Initialize some, leave others with zero values.

// var (
// 	appName   string = "MyApp"
// 	port      int    = 8080
// 	debugMode bool
// 	timeout   int
// 	ratio     float64
// )

// func main() {
// 	fmt.Println("appName   :", appName)
// 	fmt.Println("port      :", port)
// 	fmt.Println("debugMode :", debugMode)
// 	fmt.Println("timeout   :", timeout)
// 	fmt.Println("ratio     :", ratio)
// }

// 12. Compare performance concept: Create slice with pre-allocated capacity vs dynamic appends (just write both versions).

// func main() {
// 	n := 1_000_000

// 	data := make([]int, 0, n)

// 	for i := 0; i < n; i++ {
// 		data = append(data, i)
// 	}

// 	fmt.Println("Length:", len(data), "Capacity:", cap(data))
// }

// 13. Create large array [1000000]int locally in function. Check if it escapes with -gcflags="-m".

// func createArray() {
// 	var arr [1_000_000]int
// 	arr[0] = 42
// 	fmt.Println(arr[0])
// }

// func main() {
// 	createArray()
// }

// 14. Declare variables without initialization: var a int; var b float64; var c string; var d bool; var e *int; var f []int; var g map[string]int. Print them to show zero values.package main

// func main() {
// 	var a int
// 	var b float64
// 	var c string
// 	var d bool
// 	var e *int
// 	var f []int
// 	var g map[string]int

// 	fmt.Println("a (int)        :", a)
// 	fmt.Println("b (float64)    :", b)
// 	fmt.Println("c (string)     :", c)
// 	fmt.Println("d (bool)       :", d)
// 	fmt.Println("e (*int)       :", e)
// 	fmt.Println("f ([]int)      :", f)
// 	fmt.Println("g (map)        :", g)
// }

// 15. Declare multiple variables in single line: var a, b, c int = 1, 2, 3. Then update them in single assignment: a, b, c = 4, 5, 6.

// func main(){
// 	var a, b, c int = 1, 2, 3

// 	a, b, c = 4, 5, 6

// 	fmt.Println(a,b,c)

// }

// 16. Swap two variables without temporary: x, y := 10, 20; x, y = y, x. Print before and after.

// func main() {
// 	x, y := 10, 20
// 	fmt.Println("before :", x, y)
// 	x, y = y, x
// 	fmt.Println("after :", x, y)
// }

// 17. Try redeclaring variables with := when no new variables: x, y := 10, 20; x, y := 30, 40. Handle error.

// func main() {
// 	x, y := 10, 20
// 	x, y := 30, 40
// }

// 	18. Redeclare with one new variable: x, z := 50, 60. Print both.

// func main() {
// 	x, y := 10, 20
// 	x, z := 50, 60

// 	fmt.Println("x:", x)
// 	fmt.Println("y:", y)
// 	fmt.Println("z:", z)
// }

// 19. Use math.MaxInt32, math.MinInt32, math.MaxUint32 constants. Assign them to variables and print.

// func main() {
// 	var maxInt32 int32 = math.MaxInt32
// 	var minInt32 int32 = math.MinInt32
// 	var maxUint32 uint32 = math.MaxUint32

// 	fmt.Println("MaxInt32 :", maxInt32)
// 	fmt.Println("MinInt32 :", minInt32)
// 	fmt.Println("MaxUint32:", maxUint32)
// }

// 	20. Declare variables with zero values and check if they equal nil for pointer, slice, and map types.

// func main() {

// 	var p *int
// 	var s []int
// 	var m map[string]int

// 	// nil checks
// 	fmt.Println("Pointer is nil:", p == nil)
// 	fmt.Println("Slice is nil  :", s == nil)
// 	fmt.Println("Map is nil    :", m == nil)
// }
