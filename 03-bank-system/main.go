package main 

// import "fmt"

func main(){
	bank := newBank("CBE")
	bank.Welcome()

	// create customers 
	gizaw := newCustomer(0, "Gizaw", "g@gmail.com")
	emuyu := newCustomer(1, "Emuyu", "e@gmail.com")
	dad := newCustomer(2, "Dad", "dad@gmail.com")

	gizaw.display()
	emuyu.display()
	dad.display()

	// create account 
	gizawAccount := newAccount(10001, gizaw, 0.0)
	emuyuAccount := newAccount(1002, emuyu, 0.0)
	dadAccount := newAccount(1003, dad, 0.0)

	gizawAccount.Display()
	emuyuAccount.Display()
	dadAccount.Display()
}