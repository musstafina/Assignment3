package commands

import (
	"fmt"
	"project/order_manager"
)

type ModifyOrderCommand struct {
	orderID int
	newProducts []string
	orderMgr *order_manager.OrderManager
}

func NewModifyOrderCommand(orderID int, newProducts []string, orderMgr *order_manager.OrderManager) *ModifyOrderCommand {
	return &ModifyOrderCommand{
		orderID: orderID,
		newProducts: newProducts,
		orderMgr: orderMgr,
	}
}

func (c *ModifyOrderCommand) Execute() {
	if order, found := c.orderMgr.GetOrder(c.orderID); found {
		order.SetProducts(c.newProducts)
		fmt.Printf("Order #%d modified\n", c.orderID) 
	} else {
		fmt.Printf("Order #%d not found\n", c.orderID)
	}
}