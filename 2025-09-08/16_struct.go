package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        int
	amount    float64
	status    string
	createdAt time.Time
}

func (order *Order) changeStatus(status string) {
	order.status = status
}

func (order Order) getAmount() float64 {
	return order.amount
}

func main() {
	myOrder := Order{
		id:     1,
		amount: 1.2,
		status: "received",
	}

	myOrder.createdAt = time.Now()

	fmt.Println(myOrder.status)
	fmt.Println(myOrder)

	myOrder.changeStatus("delivered")
	fmt.Println(myOrder.status)

	fmt.Println(myOrder.getAmount())

	language := struct {
		name   string
		isGood bool
	}{
		name:   "japanese",
		isGood: true,
	}
	fmt.Println(language)
}
