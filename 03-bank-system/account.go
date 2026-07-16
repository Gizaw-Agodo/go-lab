package main

import "fmt"

type BankAccount struct {
	AccountNumber int
	Owner         Customer
	Balance       float64
}

func (account *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Amount should be greater than zero")
		return 
	}

	account.Balance += amount
	fmt.Printf("$%.2f deposited successfylly!",amount)
}


func (account *BankAccount) Withdraw(amount float64){
	if amount <= 0 {
		fmt.Println("invalid amount")
		return 
	}

	if amount > account.Balance {
		fmt.Println("Insuficcent balance")
	}

	account.Balance -= amount 
	fmt.Printf("$%.2f Withdrawn successfully", amount)
}

// transfer from one account to another 
func (account *BankAccount) Transfer(reciver *BankAccount, amount float64){
	if amount <=0 {
		fmt.Println("invalid amount")
		return 
	}
	if amount > account.Balance {
		fmt.Println("Insuficcient balance")
		return 
	}

	account.Withdraw(amount)
	reciver.Deposit(amount)

	fmt.Printf("$%.2f Transfered from %s, to %s successfully!", amount, account.Owner.Name, reciver.Owner.Name)

}

func (account BankAccount) Display(){
	fmt.Println("====Bank account====")
	fmt.Printf("Account Number %d ", account.AccountNumber)
	fmt.Printf("Owner : %s  ", account.Owner.Name)
	fmt.Printf("Email: %s  ", account.Owner.Email)
	fmt.Printf("balance: $%.2f\n", account.Balance)
}