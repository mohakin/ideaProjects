package main

import "fmt"
import "os"
import "strconv"
import "errors"

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error){
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil{
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil{
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}


func balanceWrite(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------------------------")
		//panic("Can't continue, sorry.")
	}

	var depositAmount float64
	var withdrawAmount float64

	fmt.Println("Welcome to the Go bank!")
	for {
		fmt.Println("Please select and option: ")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")


		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)


		if choice == 1{
			fmt.Printf("Your balance is $%.2f\n", accountBalance)
		} else if choice == 2{
			fmt.Println("How much would you like to deposit?")
			fmt.Scan(&depositAmount)

			if depositAmount<=0{
				panic(" Invalid amount detected. Must be greater than zero.")
				//return
				//continue
			}

			accountBalance += depositAmount
			fmt.Printf("Your new balance is: $%.2f\n", accountBalance)
			balanceWrite(accountBalance)
			continue
		} else if choice == 3{
			fmt.Println("How much would you like to withdraw?")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount<=0{
				panic(" Invalid amount detected. Must be greater than zero.")
				//return
				//continue
			}

			if withdrawAmount > accountBalance{
				panic("You can not withdraw more than what is in your account already.")
				//return
				//continue
			}

			accountBalance -= withdrawAmount
			fmt.Printf("Your new balance is: $%.2f\n", accountBalance)
			balanceWrite(accountBalance)
			continue
		} else{
			fmt.Println("Goodbye!")
			//return
			break
		}
	}

	fmt.Println("Thank you for using the Go bank App!")

}
