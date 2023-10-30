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

func (om *OrderManager) AddOrder(orderID int, order *Order) {
	om.mu.Lock()
	om.orders[orderID] = order
	om.mu.Unlock()
}

func (om *OrderManager) GetOrder(orderID int) (*Order, bool) {
	om.mu.Lock()
	order, found := om.orders[orderID]
	om.mu.Unlock()
	return order, found
}

func (om *OrderManager) RemoveOrder(orderID int) {
	om.mu.Lock()
	delete(om.orders, orderID)
	om.mu.Unlock()
}