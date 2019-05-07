package main

import (
	"log"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/imrenagi/microservice-demo/order-service/internal/order"
	"github.com/imrenagi/microservice-demo/order-service/web"
	paymentProto "github.com/imrenagi/microservice-demo/payment-service/pkg/proto/payment"
	nats "github.com/nats-io/go-nats"
)

func main() {

	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer natsConn.Close()

	orderService := order.NewOrderService(natsConn)

	if _, err := natsConn.Subscribe("paymentCreated", func(m *nats.Msg) {
		var payment paymentProto.PaymentCreated
		err := proto.Unmarshal(m.Data, &payment)
		if err != nil {
			log.Fatalf(err.Error())
		}

		order, err := orderService.GetOrder(payment.OrderID)
		if err != nil {
			log.Fatalf(err.Error())
		}

		if order != nil {
			log.Println("New Payment Information is accepted. Updating order status")
			orderService.UpdateStatus(order.ID, "PAID")
			log.Println("Order status is updated")
		}

	}); err != nil {
		log.Fatal(err.Error())
	}

	s := web.NewServer(orderService)
	if err := http.ListenAndServe(":8080", s.Router); err != nil {
		log.Fatalf("Server can't run. Got: `%v`", err)
	}
}
