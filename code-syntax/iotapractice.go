package main

import "fmt"

func main() {
	const (
		Pending = iota
		Processing
		Shipped
		Deliverd
		Cancelled
	)

	var order = 10
	switch order {
	case Pending:
		{
			fmt.Printf("Your Order is Pending.. %d", Pending)
		}
	case Processing:
		{
			fmt.Printf("Your Order is Proccessing %d", Processing)
		}
	case Shipped:
		{
			fmt.Printf("your Order is shipping right away %d", Shipped)
		}
	case Deliverd:
		{
			fmt.Printf("Your Order is deliverd. %d", Deliverd)
		}
	case Cancelled:
		{
			fmt.Printf("Order Cancelled. %d", Cancelled)
		}
	default:
		{
			fmt.Printf("Some Error Occured ,  Contact Our Support.")
		}

	}
}
