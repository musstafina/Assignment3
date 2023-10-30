package commands

import (
	"fmt"
	"project/order_manager"
)

type CancelOrderCommand struct {
	orderID int
	orderMgr *order_manager.OrderManager
}

func NewCancelOrderCommand(orderID int, orderMgr *order_manager.OrderManager) *CancelOrderCommand {
	return &CancelOrderCommand{
		orderID: orderID,
		orderMgr: orderMgr,
	}
}

func (c *CancelOrderCommand) Execute() {
	if _, found := c.orderMgr.GetOrder(c.orderID); found {
		c.orderMgr.RemoveOrder(c.orderID)
		fmt.Printf("Order #%d cancelled\n", c.orderID)
	} else {
		fmt.Printf("Order #%d not found\n", c.orderID)
	}
}