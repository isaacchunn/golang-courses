package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/iz/packages/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())
	var choice int
	var accountBalance = fileops.GetFloatFromFile(accountBalanceFile, 1000)

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		//Handle choices
		case 1:
			fmt.Println("Your account balance is", accountBalance)
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
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing our bank.")
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}
