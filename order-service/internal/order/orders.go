package order

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	nats "github.com/nats-io/go-nats"

	"github.com/imrenagi/microservice-demo/order-service/pkg/proto/order"
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

func NewOrderService(natsConn *nats.Conn) *OrderService {
	return &OrderService{
		data:     make(map[string]Order, 0),
		natsConn: natsConn,
	}
}

type OrderService struct {
	data     map[string]Order //map
	natsConn *nats.Conn
}

func (o *OrderService) GetOrder(orderID string) (*Order, error) {
	order := o.data[orderID]
	return &order, nil
}

func (o *OrderService) NewOrder(command NewOrderCommand) (string, error) {
	order := NewOrder(command.CustomerID, command.LineItems)
	o.data[order.ID] = order
	fmt.Println(order)

	err := o.publishNewOrder(order)
	if err != nil {
		fmt.Errorf(err.Error())
		return "", nil
	}

	return order.ID, nil
}

func (o OrderService) publishNewOrder(newOrder Order) error {

	event := order.OrderCreated{
		ID:            newOrder.ID,
		CustomerID:    newOrder.CustomerID,
		LineItems:     newOrder.LineItems,
		Price:         float32(newOrder.Price),
		PaymentStatus: newOrder.PaymentStatus,
	}

	bytes, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	if err := o.natsConn.Publish("orderCreated", bytes); err != nil {
		return err
	}
	o.natsConn.Flush()

	log.Printf("New Order is published: %v", newOrder)
	return nil
}

func (o *OrderService) UpdateStatus(orderID, status string) error {
	if order, ok := o.data[orderID]; ok {
		order.MakePayment()
		o.data[orderID] = order
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
