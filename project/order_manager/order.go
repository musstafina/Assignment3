package order_manager

type Order struct {
	orderID int
	products []string
}

func NewOrder(orderID int, products []string) *Order {
	return &Order{
		orderID: orderID,
		products: products,
	}
}