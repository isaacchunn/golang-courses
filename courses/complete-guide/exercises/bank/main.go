package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() float64 {
	bs, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error:", err)
		//We can just return a default value of 1000
		return 1000.0
	}
	balanceText := string(bs)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	var choice int
	var accountBalance = readBalanceFromFile()

	for {
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		//Handle choices
		case 1:
			fmt.Println("Your account balance is", accountBalance)
			break
		case 2:
			fmt.Print("How much do you want to deposit? ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount > 0 {
				accountBalance += depositAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
			} else {
				fmt.Println("Invalid deposit amount!")
			}
			break
		case 3:
			fmt.Print("How much do you want to withdraw? ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount < accountBalance {
				accountBalance -= withdrawAmount
				fmt.Printf("%f was withdrawn from the account. New balance: %f\n", withdrawAmount, accountBalance)
			} else if withdrawAmount < 0 {
				fmt.Println("Invalid withdraw amount!")
			} else {
				fmt.Println("Insufficient balance in bank account!")
			}
			break
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing our bank.")
			writeBalanceToFile(accountBalance)
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}
