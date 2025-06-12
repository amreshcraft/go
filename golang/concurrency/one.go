package main

import (
	"fmt"
	"time"
)

type Order struct{
	ID int
	Status string
}
func main() {
	orders :=  generateOrders(20)
	go processOrders(orders)
	fmt.Println("all operation completed. exited")
}

func updateOrderStatuses(orders []*Order){
	for _,order:= range orders{
		status := []string{
			"Processing","Shipped","Delivered",
		}
		order.Status = status[1]
		fmt.Printf("Update Order %d status: %s \n", order.ID,order.Status,)
	}
}
func processOrders(order []*Order){
	for _, order := range order{
		time.Sleep(time.Duration(500 ) * time.Millisecond)
		fmt.Printf("Processing Orders : %d\n",order.ID)
	}
}

func generateOrders(count int) []*Order{
	orders := make([]*Order,count)
for i := 0; i < count; i++ {
 orders[i] = &Order{ID: i, Status: "Pending"}	
}
	return orders
}