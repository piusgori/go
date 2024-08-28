package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

 	return value, nil
}

func main() {

	accountBalance, err := getFloatFromFile(accountBalanceFile);

	
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Can't continue, sorry.")
	}
	
	fmt.Println("Welcome to Go Bank!")
	
	for {
		presentOptions()
		
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is", accountBalance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
	
			accountBalance += depositAmount
			writeFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount:", accountBalance)

		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
	
			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
	
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
	
			accountBalance -= withdrawalAmount
			writeFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount:", accountBalance)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}