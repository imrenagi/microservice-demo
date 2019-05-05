package order

import (
	"fmt"

	"github.com/google/uuid"
)

type NewOrderCommand struct {
	CustomerID string   `json:"customer_id"`
	LineItems  []string `json:"line_items"`
}

func NewOrder(customerID string, lineItems []string) Order {
	return Order{
		ID:            uuid.New().String(),
		CustomerID:    customerID,
		LineItems:     lineItems,
		Price:         20.0, //hardcoded
		PaymentStatus: "UNPAID",
	}
}

type Order struct {
	ID            string
	CustomerID    string
	LineItems     []string
	Price         float64
	PaymentStatus string
}

func (o *Order) MakePayment() {
	o.PaymentStatus = "PAID"
}

func NewOrderService() *OrderService {
	return &OrderService{
		data: make(map[string]Order, 0),
	}
}

type OrderService struct {
	data map[string]Order //map
}

func (o *OrderService) NewOrder(command NewOrderCommand) (string, error) {
	order := NewOrder(command.CustomerID, command.LineItems)
	o.data[order.ID] = order
	fmt.Println(order)

	//TODO publish event

	return order.ID, nil
}

func (o OrderService) UpdateStatus(orderID, status string) error {
	if order, ok := o.data[orderID]; !ok {
		order.MakePayment()
		return nil
	}
	return fmt.Errorf("order doesn't exist")
}

func (o OrderService) GetOrders() ([]Order, error) {
	orders := make([]Order, 0)
	for _, v := range o.data {
		orders = append(orders, v)
	}
	return orders, nil
}
