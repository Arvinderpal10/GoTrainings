// Switch Statement & Conditionals
package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Q1. Vowel Checker")
	//   Write a program that takes a single character as input and determines if it's a vowel or consonant.
	/*var ch string

	fmt.Print("Enter a single character: ")
	fmt.Scan(&ch)

	switch ch {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}*/

	// - Handle both uppercase and lowercase letters
	/*var ch string
	fmt.Println("Enter a single character:")
	fmt.Scan(&ch)
	switch ch {
	case "a", "e", "i", "o", "u",
		"A", "E", "I", "O", "U":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}*/

	// - Use multiple values in cases (a, e, i, o, u, A, E, I, O, U)
	// import unicode package
	/*var ch rune
	fmt.Print("Enter a single character: ")
	fmt.Scanf("%c", &ch)
	if !unicode.IsLetter(ch) {
		fmt.Println("Not a letter")
		return
	}
	switch ch {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}*/

	// - For non-alphabetic characters → "Not a letter"
	// import unicode package
	/*var ch rune
	fmt.Print("Enter a single character: ")
	fmt.Scanf("%c", &ch)
	if !unicode.IsLetter(ch) {
		fmt.Println("Not a letter")
		return
	}
	switch ch {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}*/

	//fmt.Println("Q2. BMI Categorization")
	/*var weight, height float64
	fmt.Print("Enter weight in kg: ")
	fmt.Scan(&weight)
	fmt.Print("Enter height in meters: ")
	fmt.Scan(&height)
	// Calculate BMI (Body Mass Index) using formula: BMI = weight(kg) / (height(m))²
	BMI := weight / (height * height)
	fmt.Printf("Your BMI is: %.2f\n", BMI)
	// Then categorize using tagless switch:
	// - BMI < 18.5 → "Underweight"
	// - 18.5 ≤ BMI < 25 → "Normal"
	// - 25 ≤ BMI < 30 → "Overweight"
	// - BMI ≥ 30 → "Obese"
	switch {
	case BMI < 18.5:
		fmt.Println("Category: Underweight")
	case BMI >= 18.5 && BMI < 25:
		fmt.Println("Category: Normal")
	case BMI >= 25 && BMI < 30:
		fmt.Println("Category: Overweight")
	default:
		fmt.Println("Category: Obese")
	}*/
	// Print the BMI value and category.

	//fmt.Println("Q3. Ticket Price Calculator")
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
	/*var age int
	var dayType string
	var student string
	price := 10.0
	fmt.Println("Enter age:")
	fmt.Scan(&age)
	fmt.Print("Enter day type (weekday/weekend): ")
	fmt.Scan(&dayType)
	fmt.Print("Are you a student? (yes/no): ")
	fmt.Scan(&student)
	dayType = strings.ToLower(dayType)
	student = strings.ToLower(student)
	switch {
	case age < 12:
		price *= 0.5
	case age >= 65:
		price *= 0.7
	case age >= 12 && age <= 64:
		if student == "yes" && dayType == "weekday" {
			price *= 0.8
		}
	}
	if dayType == "weekend" {
		price += 2
	}
	fmt.Printf("Final Ticket Price: $%.2f\n", price)*/

	//fmt.Println("Q4. Roman Numeral to Integer")
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
	/*var roman string
	fmt.Print("Enter a Roman numeral (I, II, III, IV, V, VI, VII, VIII, IX, X): ")
	fmt.Scan(&roman)
	switch roman {
	case "I":
		fmt.Println("Integer value: 1")
	case "II":
		fmt.Println("Integer value: 2")
	case "III":
		fmt.Println("Integer value: 3")
	case "IV":
		fmt.Println("Integer value: 4")
	case "V":
		fmt.Println("Integer value: 5")
	case "VI":
		fmt.Println("Integer value: 6")
	case "VII":
		fmt.Println("Integer value: 7")
	case "VIII":
		fmt.Println("Integer value: 8")
	case "IX":
		fmt.Println("Integer value: 9")
	case "X":
		fmt.Println("Integer value: 10")
	default:
		fmt.Println("Invalid Roman numeral")
	}*/

	fmt.Println("Q5. Banking Transaction System")
	// Simulate banking operations:
	// - Input: transaction type ("deposit", "withdraw", "balance", "transfer")
	// - For deposit: add amount to balance
	// - For withdraw: check if sufficient balance
	// - For transfer: deduct from one account, add to another
	// - For invalid transaction → "Invalid operation"
	// Use switch with short statement to initialize variables.
	var balance float64 = 1000
	var secondAccount float64 = 500
	var input string
	fmt.Println("Initial Balance:", balance)
	fmt.Println("Second Account Balance:", secondAccount)
	fmt.Print("Enter transaction type (deposit/withdraw/balance/transfer): ")
	fmt.Scan(&input)
	switch input {
	case "deposit":
		var amount float64
		fmt.Print("Enter amount to deposit: ")
		fmt.Scan(&amount)
		balance += amount
		fmt.Printf("New Balance: $%.2f\n", balance)
	case "withdraw":
		var amount float64
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scan(&amount)
		if amount > balance {
			fmt.Println("Insufficient balance")
		} else {
			balance -= amount
			fmt.Printf("New Balance: $%.2f\n", balance)
		}
	case "balance":
		fmt.Printf("Current Balance: $%.2f\n", balance)
	case "transfer":
		var amount float64
		fmt.Print("Enter amount to transfer: ")
		fmt.Scan(&amount)
		if amount > balance {
			fmt.Println("Insufficient balance for transfer")
		} else {
			balance -= amount
			secondAccount += amount
			fmt.Printf("New Balance: $%.2f\n", balance)
			fmt.Printf("Second Account New Balance: $%.2f\n", secondAccount)
		}
	default:
		fmt.Println("Invalid operation")
	}

	//fmt.Println(" - - - - - - - - - - Theory Questions - - - - - - - - - - - ")

	// Section 1: Multiple Choice Questions

	// 1. What happens when no case matches in a switch statement?
	//    a) Compilation error
	//    b) Runtime error
	//    c) Executes default case (if present)
	//    d) Executes first case
	//ANS: (c)

	// 2. How does Go handle fallthrough in switch?
	//    a) Automatic fallthrough by default
	//    b) Manual using `continue`
	//    c) Manual using `fallthrough`
	//    d) No fallthrough allowed
	//ANS: (c)

	// 3. Which is NOT allowed in Go's if condition?
	//    a) if true { ... }
	//    b) if 1 { ... }
	//    c) if x > 5 { ... }
	//    d) if condition() { ... }
	//ANS: (b)

	// 4. What is the scope of a variable declared in switch short statement?
	//    a) Entire program
	//    b) Only the case where declared
	//    c) Entire switch block
	//    d) Only the function
	//ANS: (c)

	// 5. Which is better for checking same variable against many values?
	//    a) Multiple if statements
	//    b) Switch statement
	//    c) Both are equal
	//    d) Ternary operator
	//ANS:(b)

	// Section 2: True/False Questions

	// 1. T/F: Go requires braces {} for if blocks TRUE
	// 2. T/F: Switch cases in Go are case-sensitive for strings TRUE
	// 3. T/F: You can use float values in switch cases TRUE
	// 4. T/F: Default case is mandatory in switch FALSE
	// 5. T/F: Tagless switch can have conditions in cases FALSE

}
