// Switch Statement & Conditionals
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Q1. Vowel Checker")
	//   Write a program that takes a single character as input and determines if it's a vowel or consonant.
	// - Handle both uppercase and lowercase letters
	// - Use multiple values in cases (a, e, i, o, u, A, E, I, O, U)
	// - For non-alphabetic characters → "Not a letter"

	var ch rune
	fmt.Print("Enter a character:")
	fmt.Scanf("%c",&ch)

	switch ch{
		case 'a','e','i','o','u',
		      'A','E','I','O','U':
			  fmt.Println("vowel")

		default :
			if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			fmt.Println("Consonant")
		} else {
			fmt.Println("Not a letter")
		}

	}


	fmt.Println("Q2. BMI Categorization")
	// Calculate BMI (Body Mass Index) using formula: BMI = weight(kg) / (height(m))²
	// Then categorize using tagless switch:
	// - BMI < 18.5 → "Underweight"
	// - 18.5 ≤ BMI < 25 → "Normal"
	// - 25 ≤ BMI < 30 → "Overweight"
	// - BMI ≥ 30 → "Obese"
	// Print the BMI value and category.

	var weight,height float32

	fmt.Print("Enter weight and height: ")
	fmt.Scan(&weight, &height)
	height = height / 100     // Convert cm to meters

	bmi:= (weight)/(height*height)
	fmt.Printf("BMI %.2f\n",bmi)

	switch{
	case bmi<18.5:
		fmt.Println("underweight")

	case bmi>=18.5&& bmi<25:
		fmt.Println("normal")

	case bmi>=23&& bmi<30:
		fmt.Println("overweight")	

	default:
		fmt.Println("obese")
	}


	fmt.Println("Q3. Ticket Price Calculator")
	// Calculate ticket price based on:
	// - Age (child: <12, adult: 12-64, senior: 65+)
	// - Day type (weekday/weekend)
	// - Student status (yes/no)

	// Rules:
	// - Base price: $10
	// - Children: 50% discount
	// - Seniors: 30% discount
	// - Students (adult age): 20% discount (only on weekdays)
	// - Weekend: +$2 extra

	var age int
	var dayType string
	var isStudent string

	baseprice:=10.0
	price:=baseprice

	fmt.Println("Enter age")
	fmt.Scan(&age)

	fmt.Println("Enter day type(weekend/weekday)")
	fmt.Scan(&dayType)

	fmt.Println("Are you a student(yes/no)")
	fmt.Scan(&isStudent)

	dayType=strings.ToLower(dayType)
	isStudent=strings.ToLower(isStudent)

	if age<12{
		price=price*0.5
	}else if age>=65{
		price=price*0.7
	}else if age>=12 && age<=64 && dayType=="weekday" && isStudent=="yes"{
		price=price*0.8
	}

	if dayType=="weekend"{
		price=price+2
	}
	fmt.Printf("Final Ticket Price: $%.2f\n", price)


	fmt.Println("Q4. Roman Numeral to Integer")

	//  Convert Roman numerals to integers for values 1-10:
	// - I → 1
	// - II → 2
	// - III → 3
	// - IV → 4
	// - V → 5
	// - VI → 6
	// - VII → 7
	// - VIII → 8
	// - IX → 9
	// - X → 10
	// For invalid input → "Invalid Roman numeral"

	var roman string
	fmt.Print("Enter a roman numeral:=")
	fmt.Scan(&roman)

	switch roman{
		case "I":
			fmt.Println(1)
		case "II":
			fmt.Println(2)
		case "III":
			fmt.Println(3)
		case "IV":
			fmt.Println(4)
		case "V":
			fmt.Println(5)
		case "VI":
			fmt.Println(6)
		case "VII":
			fmt.Println(7)
		case "VIII":
			fmt.Println(8)
		case "IX":
			fmt.Println(9)
		case "X":
			fmt.Println(10)
		default:
			fmt.Println("Invalid numeral")
		
	}

	fmt.Println("Q5. Banking Transaction System")
	// Simulate banking operations:
	// - Input: transaction type ("deposit", "withdraw", "balance", "transfer")
	// - For deposit: add amount to balance
	// - For withdraw: check if sufficient balance
	// - For transfer: deduct from one account, add to another
	// - For invalid transaction → "Invalid operation"
	// Use switch with short statement to initialize variables.

	balance := 1000
	var op string

	fmt.Print("Enter operation: ")
	fmt.Scan(&op)

	switch op {

	case "deposit":
		var amount int
		fmt.Scan(&amount)
		balance = balance + amount
		fmt.Println("Balance:", balance)

	case "withdraw":
		var amount int
		fmt.Scan(&amount)
		if amount <= balance {
			balance = balance - amount
			fmt.Println("Balance:", balance)
		} else {
			fmt.Println("Insufficient balance")
		}

	case "balance":
		fmt.Println("Balance:", balance)

	case "transfer":
		var amount int
		fmt.Scan(&amount)
		if amount <= balance {
			balance = balance - amount
			fmt.Println("Transfer done")
			fmt.Println("Balance:", balance)
		} else {
			fmt.Println("Insufficient balance")
		}

	default:
		fmt.Println("Invalid operation")
	}



	fmt.Println(" - - - - - - - - - - Theory Questions - - - - - - - - - - - ")

	// Section 1: Multiple Choice Questions

	// 1. What happens when no case matches in a switch statement?-----------C
	//    a) Compilation error
	//    b) Runtime error
	//    c) Executes default case (if present)
	//    d) Executes first case

	// 2. How does Go handle fallthrough in switch?--------------------------C
	//    a) Automatic fallthrough by default
	//    b) Manual using `continue`
	//    c) Manual using `fallthrough`
	//    d) No fallthrough allowed

	// 3. Which is NOT allowed in Go's if condition?-------------------------B
	//    a) if true { ... }
	//    b) if 1 { ... }
	//    c) if x > 5 { ... }
	//    d) if condition() { ... }

	// 4. What is the scope of a variable declared in switch short statement?----------C
	//    a) Entire program
	//    b) Only the case where declared
	//    c) Entire switch block
	//    d) Only the function

	// 5. Which is better for checking same variable against many values?---------------B
	//    a) Multiple if statements
	//    b) Switch statement
	//    c) Both are equal
	//    d) Ternary operator

	// Section 2: True/False Questions

	// 1. T/F: Go requires braces {} for if blocks    -------------------------True
	// 2. T/F: Switch cases in Go are case-sensitive for strings --------------True
	// 3. T/F: You can use float values in switch cases -----------------------True
	// 4. T/F: Default case is mandatory in switch ----------------------------False
	// 5. T/F: Tagless switch can have conditions in cases --------------------True

}
