package main

import "fmt"

func main() {
	 input:=34
	 if (input>='a'&& input<='z') || (input>='A'&& input<='Z'){
		switch input {
	 case 'A','E','I','O','U','a','e','i','o','u':
		fmt.Println("vowel")
	 default:
		fmt.Println("consonant")
	 }
    }else{
		fmt.Println("Not a letter")
	}


	}
func main(){
	var weight float32=62.00
	var height float32=1.79
	BMI:=weight / (height * height)
	switch {
	case BMI < 18.5 :
		fmt.Println("Underweight")
	case BMI < 25:
		fmt.Println("Normal")
	case BMI < 30 :
		fmt.Println("overweight")
	case BMI >=30 :
		fmt.Println("obese")
	}



}
func main(){
	age := 20
	dayType := "weekend"   
	student := "yes"       

	price := 10.0

	if age < 12 {
		price = price * 0.5
	} else if age >= 65 {
		price = price * 0.7
	} else {
		if student == "yes" && dayType == "weekday" {
			price = price * 0.8
		}
	}

	if dayType == "weekend" {
		price = price + 2
	}

	fmt.Printf("Final Ticket Price: $%.2f\n", price)
}



func main(){
	roman:="I"
	switch roman{
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
}



func main() {
	balance1 := 1000.0
	balance2 := 500.0

	transactionType := "transfer" 
	amount := 300.0

	switch txn := transactionType; txn {

	case "deposit":
		balance1 += amount
		fmt.Println("Deposit successful.")
		fmt.Printf("Updated Balance: $%.2f\n", balance1)

	case "withdraw":
		if amount <= balance1 {
			balance1 -= amount
			fmt.Println("Withdrawal successful.")
			fmt.Printf("Updated Balance: $%.2f\n", balance1)
		} else {
			fmt.Println("Insufficient balance.")
		}

	case "transfer":
		if amount <= balance1 {
			balance1 -= amount
			balance2 += amount
			fmt.Println("Transfer successful.")
			fmt.Printf("Account1 Balance: $%.2f\n", balance1)
			fmt.Printf("Account2 Balance: $%.2f\n", balance2)
		} else {
			fmt.Println("Insufficient balance for transfer.")
		}

	case "balance":
		fmt.Printf("Current Balance: $%.2f\n", balance1)

	default:
		fmt.Println("Invalid operation")
	}
}

	
		