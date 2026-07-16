package main 
import "fmt"

func getBundlePrice(bundle string) float64 {
	switch bundle {
	case "Daily":
		return 300.0
	case "Weekly":
		return 600.0
	case "Monthly":
		return 1000.0
	}
	return 0.0
}

func getDiscountPrice(price float64, isPremium bool ) float64 {
	if  isPremium {
		return price * 0.9
	}
	return price

}

func purchaseBundle(balance float64, cost float64) (float64, bool) {
	if (balance >= cost){
		return balance - cost, true
	}
	return balance, false
}

func main(){
	name := "Gizaw"
	bundle := "Monthly"
	balance := 300.0
	isPremium := false

	fmt.Println("Welcome", name)
	price := getBundlePrice(bundle)
	discountPrice := getDiscountPrice(price, isPremium)

	fmt.Println("Bundle", bundle)
	fmt.Println("Original price", price)
	fmt.Println("Final price", discountPrice)

	newBalance, success := purchaseBundle(balance, discountPrice)

	if (success){
		fmt.Println("purchage successfull !")
		fmt.Println("New balance", newBalance)
	}else {
		fmt.Println("Insufficient balance")
	}

}