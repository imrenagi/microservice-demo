package payment

import (
	"fmt"

	"github.com/google/uuid"
)

type NewPaymentCommand struct {
	OrderID       string  `json:"order_id"`
	Value         float64 `json:"value"`
	CreditCardNum string  `json:"credit_card_number"`
	CreditCardCVC string  `json:"credit_card_cvc"`
}

func NewPayment(orderID string, value float64) Payment {
	return Payment{
		ID:      uuid.New().String(),
		OrderID: orderID,
		Value:   value,
	}
}

type Payment struct {
	ID      string
	OrderID string
	Value   float64
}

type Order struct {
	OrderID string
	Price   float64
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		orders: make(map[string]Order, 0),
		data:   make(map[string]Payment, 0),
	}
}

type PaymentService struct {
	orders map[string]Order
	data   map[string]Payment
}

func (p *PaymentService) MakePayment(command NewPaymentCommand) (string, error) {
	if order, ok := p.orders[command.OrderID]; ok {
		if command.Value < order.Price {
			return "", fmt.Errorf("balance is not enough")
		}
		payment := NewPayment(order.OrderID, command.Value)
		p.data[payment.ID] = payment
		//TODO publish payment created

		return payment.ID, nil
	}
	return "", fmt.Errorf("order not found. Payment is not created")
}

func (p *PaymentService) AddOrder(order Order) error {
	p.orders[order.OrderID] = order
	return nil
}

func (p PaymentService) GetPayment() ([]Payment, error) {
	payments := make([]Payment, 0)
	for _, v := range p.data {
		payments = append(payments, v)
	}
	return payments, nil
}
