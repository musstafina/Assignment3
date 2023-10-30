package order_manager

import (
	"sync"
)

type OrderManager struct {
	orders map[int]*Order
	mu sync.Mutex
}

var instance *OrderManager

func GetOrderManager() *OrderManager {
	if instance == nil {
		instance = &OrderManager{
			orders: make(map[int]*Order),
		}
	}
	return instance
}