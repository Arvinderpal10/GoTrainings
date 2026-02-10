// Switch Statement & Conditionals
package main

import "fmt"

func getAmount(transaction string) float64 { // Related to Q5
	var amount float64

	if transaction == "deposit" || transaction == "withdraw" || transaction == "transfer" {
		fmt.Print("Enter amount: ")
		fmt.Scanln(&amount)
	}

	return amount
}

func main() {
	fmt.Println("Q1. Vowel Checker")
	//   Write a program that takes a single character as input and determines if it's a vowel or consonant.
	// - Handle both uppercase and lowercase letters
	// - Use multiple values in cases (a, e, i, o, u, A, E, I, O, U)
	// - For non-alphabetic characters → "Not a letter"
	// var ch int

	// fmt.Print("Enter a single character: ")
	// fmt.Scanf("%c", &ch)

	// if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') { // Check if character is an alphabet

	// 	switch ch {
	// 	case 'a', 'e', 'i', 'o', 'u',
	// 		'A', 'E', 'I', 'O', 'U':
	// 		fmt.Println("Vowel")
	// 	default:
	// 		fmt.Println("Consonant")
	// 	}

	// } else {
	// 	fmt.Println("Not a letter")
	// }
	// fmt.Println("Q2. BMI Categorization")
	// Calculate BMI (Body Mass Index) using formula: BMI = weight(kg) / (height(m))²
	// Then categorize using tagless switch:
	// - BMI < 18.5 → "Underweight"
	// - 18.5 ≤ BMI < 25 → "Normal"
	// - 25 ≤ BMI < 30 → "Overweight"
	// - BMI ≥ 30 → "Obese"
	// Print the BMI value and category.
	// var weight float64
	// var height float64
	// fmt.Print("Enter weight in kg: ")
	// fmt.Scan(&weight)
	// fmt.Print("Enter height in meters: ")
	// fmt.Scan(&height)
	// fmt.Println("Enter weight")
	weight := 70.0
	height := 1.75
	bmi := weight / (height * height)
	switch {
	case bmi < 18.5:
		fmt.Println("Underweight")
	case bmi >= 18.5 && bmi < 25:
		fmt.Println("Normal")
	case bmi >= 25 && bmi < 30:
		fmt.Println("Overweight")
	case bmi >= 30:
		fmt.Println("Obese")
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
	// var age int
	// var dayType string
	// var isStudent string

	// basePrice := 10.0
	// finalPrice := basePrice

	// fmt.Print("Enter age: ")
	// fmt.Scanln(&age)

	// fmt.Print("Enter day type (weekday/weekend): ")
	// fmt.Scanln(&dayType)

	// fmt.Print("Are you a student? (yes/no): ")
	// fmt.Scanln(&isStudent)

	// // Age-based discounts
	// if age < 12 {
	// 	finalPrice = finalPrice * 0.5 // 50% discount
	// } else if age >= 65 {
	// 	finalPrice = finalPrice * 0.7 // 30% discount
	// } else {
	// 	// Adult student discount (only on weekdays)
	// 	if isStudent == "yes" && dayType == "weekday" {
	// 		finalPrice = finalPrice * 0.8 // 20% discount
	// 	}
	// }

	// // Weekend extra charge
	// if dayType == "weekend" {
	// 	finalPrice += 2
	// }

	// fmt.Printf("Final Ticket Price: $%.2f\n", finalPrice)

	// fmt.Println("Q4. Roman Numeral to Integer")

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
	fmt.Println("Enter a Roman Number: ")
	fmt.Scan(&roman)
	switch roman {
	case "I":
		fmt.Println("1")
	case "II":
		fmt.Println("2")
	case "III":
		fmt.Println("3")
	case "IV":
		fmt.Println("4")
	case "V":
		fmt.Println("5")
	case "VI":
		fmt.Println("6")
	case "VII":
		fmt.Println("7")
	case "VIII":
		fmt.Println("8")
	case "IX":
		fmt.Println("9")
	case "X":
		fmt.Println("10")
	default:
		fmt.Println("Invalid Roman numeral")
	}

	fmt.Println("Q5. Banking Transaction System")
	// Simulate banking operations:
	// - Input: transaction type ("deposit", "withdraw", "balance", "transfer")
	// - For deposit: add amount to balance
	// - For withdraw: check if sufficient balance
	// - For transfer: deduct from one account, add to another
	// - For invalid transaction → "Invalid operation"
	// Use switch with short statement to initialize variables.
	// Initial balances
	accountA := 5000
	accountB := 3000

	var transactionType string
	var amount int

	// Menu
	fmt.Println("Choose transaction:")
	fmt.Println("deposit | withdraw | balance | transfer")
	fmt.Print("Enter transaction type: ")
	fmt.Scan(&transactionType)

	switch t := transactionType; t {

	case "deposit":
		fmt.Print("Enter amount: ")
		fmt.Scan(&amount)

		accountA += amount
		fmt.Println("Deposit successful")
		fmt.Println("Account A balance:", accountA)

	case "withdraw":
		fmt.Print("Enter amount: ")
		fmt.Scan(&amount)

		if amount <= accountA {
			accountA -= amount
			fmt.Println("Withdrawal successful")
			fmt.Println("Account A balance:", accountA)
		} else {
			fmt.Println("Insufficient balance")
		}

	case "balance":
		fmt.Println("Account A balance:", accountA)

	case "transfer":
		fmt.Print("Enter amount: ")
		fmt.Scan(&amount)

		if amount <= accountA {
			accountA -= amount
			accountB += amount
			fmt.Println("Transfer successful")
			fmt.Println("Account A balance:", accountA)
			fmt.Println("Account B balance:", accountB)
		} else {
			fmt.Println("Insufficient balance for transfer")
		}

	default:
		fmt.Println("Invalid operation")
	}
	fmt.Println(" - - - - - - - - - - Theory Questions - - - - - - - - - - - ")

	// Section 1: Multiple Choice Questions

	// 1. What happens when no case matches in a switch statement?
	//    a) Compilation error
	//    b) Runtime error
	//    c) Executes default case (if present) **correct
	//    d) Executes first case

	// 2. How does Go handle fallthrough in switch?
	//    a) Automatic fallthrough by default
	//    b) Manual using `continue`
	//    c) Manual using `fallthrough` **correct
	//    d) No fallthrough allowed

	// 3. Which is NOT allowed in Go's if condition?
	//    a) if true { ... }
	//    b) if 1 { ... } **correct
	//    c) if x > 5 { ... }
	//    d) if condition() { ... }

	// 4. What is the scope of a variable declared in switch short statement?
	//    a) Entire program
	//    b) Only the case where declared
	//    c) Entire switch block **correct
	//    d) Only the function

	// 5. Which is better for checking same variable against many values?
	//    a) Multiple if statements
	//    b) Switch statement **correct
	//    c) Both are equal
	//    d) Ternary operator

	// Section 2: True/False Questions

	// 1. T/F: Go requires braces {} for if blocks: True
	// 2. T/F: Switch cases in Go are case-sensitive for strings : True
	// 3. T/F: You can use float values in switch cases : False
	// 4. T/F: Default case is mandatory in switch : False
	// 5. T/F: Tagless switch can have conditions in cases : True

}
