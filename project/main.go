package main

import (
	"project/commands"
	"project/order_manager"
)

func main() {
	orderMgr := order_manager.GetOrderManager()

	createOrderCommand := commands.NewCreateOrderCommand(1, []string{"Product A", "Product B"}, orderMgr)
	createOrderCommand.Execute()

	modifyOrderCommand := commands.NewModifyOrderCommand(1, []string{"Modified Product A", "Modified Product B"}, orderMgr)
	modifyOrderCommand.Execute()

	cancelOrderCommand := commands.NewCancelOrderCommand(1, orderMgr)
	cancelOrderCommand.Execute()

}