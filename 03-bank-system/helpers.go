package main

func newCustomer(id int, name string, email string) Customer {
	return Customer {
		ID: id,
		Name : name, 
		Email: email,
	}
}

func newAccount(accountNumber int, owner Customer, balance float64) BankAccount {
     return BankAccount{
		AccountNumber: accountNumber,
		Owner: owner,
		Balance: balance,
	 }
}

func newBank(name string) Bank {
	return Bank{
	  Name: name,
	}
}