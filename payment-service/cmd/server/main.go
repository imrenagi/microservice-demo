package main

import (
	"log"
	"net/http"

	"github.com/gogo/protobuf/proto"
	orderProto "github.com/imrenagi/microservice-demo/order-service/pkg/proto/order"
	"github.com/imrenagi/microservice-demo/payment-service/internal/payment"
	"github.com/imrenagi/microservice-demo/payment-service/web"
	nats "github.com/nats-io/go-nats"
)

func main() {

	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer natsConn.Close()

	paymentService := payment.NewPaymentService(natsConn)

	if _, err := natsConn.Subscribe("orderCreated", func(m *nats.Msg) {
		var order orderProto.OrderCreated
		err := proto.Unmarshal(m.Data, &order)
		if err != nil {
			log.Fatalf(err.Error())
		}

		//store the order to the payment service
		paymentService.AddOrder(payment.Order{
			OrderID: order.ID,
			Price:   float64(order.Price),
		})
		log.Println("New Order Accepted. Stored to local storage.")

	}); err != nil {
		log.Fatal(err.Error())
	}

	s := web.NewServer(paymentService)
	if err := http.ListenAndServe(":8081", s.Router); err != nil {
		log.Fatalf("Server can't run. Got: `%v`", err)
	}
}

func subscribeToOrderCreated(natsConn *nats.Conn) {

}
