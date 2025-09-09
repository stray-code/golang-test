package main

import "fmt"

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}
func main() {
	changeOrderStatus(Delivered)
}
