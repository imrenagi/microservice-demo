package main

import (
	"log"
	"net/http"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/imrenagi/microservice-demo/payment-service/internal/payment"
	orderProto "github.com/imrenagi/microservice-demo/payment-service/pkg/proto/order"
	"github.com/imrenagi/microservice-demo/payment-service/web"
	nats "github.com/nats-io/go-nats"
)

func main() {

	natsConn, err := nats.Connect(os.Getenv("NATS_HOST"), nats.UserInfo(os.Getenv("NATS_USERNAME"), os.Getenv("NATS_PASSWORD")))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer natsConn.Close()

	paymentService := payment.NewPaymentService(natsConn)

	if _, err := natsConn.Subscribe("orderCreated", func(m *nats.Msg) {
		// if _, err := natsConn.QueueSubscribe("orderCreated", "worker", func(m *nats.Msg) {
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
	if err := http.ListenAndServe(":8080", s.Router); err != nil {
		log.Fatalf("Server can't run. Got: `%v`", err)
	}
}

func subscribeToOrderCreated(natsConn *nats.Conn) {

}
