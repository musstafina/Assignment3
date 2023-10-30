package commands

import (
	"fmt"
	"project/order_manager"
)

type CreateOrderCommand struct {
	orderID int
	products []string
	orderMgr *order_manager.OrderManager
}

func NewCreateOrderCommand(orderID int, products []string, orderMgr *order_manager.OrderManager) *CreateOrderCommand {
	return &CreateOrderCommand{
		orderID: orderID,
		products: products,
		orderMgr: orderMgr,
	}
} 

func (c *CreateOrderCommand) Execute() {
	order := order_manager.NewOrder(c.orderID, c.products)
	c.orderMgr.AddOrder(c.orderID, order)
	fmt.Printf("Order #%d created\n", c.orderID)

}