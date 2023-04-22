package main // needed at start of every executable Go file

import "fmt" // allows for formatted I/O

type Customer struct {
	Name    string
	Balance float64
}

type BobaOrder struct {
	User     Customer // Example of using struct type in other struct
	Order    []string
	Subtotal float64
}

/* Map Syntax: map[key-type]value-type */
var menu = map[string]float64{
	"Wintermelon":  5.99,
	"Passionfruit": 6.49,
	"Mango":        8.95,
	"Strawberry":   10.99,
	"Coffee":       9.95,
	"Honeydew":     7.49,
	"Avocado":      11.29,
	"Water":        1.99,
}

/*
Function Syntax: func name (params) return
Note: multiple return values allowed!
*/
func calculateSubtotal(order []string) (float64, error) {
	var subtotal float64
	// use _ to ignore index in (index, value) pair of order
	for _, item := range order {
		price, ok := menu[item]
		if ok {
			subtotal += price
		} else {
			return 0, fmt.Errorf("Invalid menu item: %s", item)
		}
	}
	return subtotal, nil
}

/* The main function is the entry point to a Go program. */
func main() {
	var orders []BobaOrder

	// Prompt the customer to input name
	var name string
	fmt.Print("Please enter your name: ")
	fmt.Scanln(&name)

	// Prompt the customer to input balance
	var balance float64
	fmt.Print("Please enter your spending balance: ")
	fmt.Scanln(&balance)

	// Create a Customer object from the struct
	customer := Customer{Name: name, Balance: balance}

	// Prompt the customer to input boba order item(s)
	fmt.Println("Enter your boba order, or 'done' to finish:")
	var orderItem string
	var order []string
	// This loop is infinite until 'done' is entered
	for {
		fmt.Scanln(&orderItem)
		if orderItem == "done" {
			break
		}
		order = append(order, orderItem) // append to slice object
	}

	// Calculate subtotal of boba order
	subtotal, err := calculateSubtotal(order)
	if err != nil {
		fmt.Println(err)
		return
	}

	if customer.Balance < subtotal {
		fmt.Println("Sorry, you have insufficient funds!")
		return
	}

	// Create a boba order object from the struct
	bobaOrder := BobaOrder{User: customer, Order: order, Subtotal: subtotal}

	// Add the boba order to the pending orders
	orders = append(orders, bobaOrder) // append to slice object

	// Print message confirming order
	fmt.Println(orders)
}
