package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/imrenagi/microservice-demo/order-service/internal/order"
	paymentProto "github.com/imrenagi/microservice-demo/order-service/pkg/proto/payment"
	"github.com/imrenagi/microservice-demo/order-service/web"
	nats "github.com/nats-io/go-nats"
)

func main() {

	fmt.Println(fmt.Sprintf("%s - %s - %s", os.Getenv("NATS_HOST"), os.Getenv("NATS_USERNAME"), os.Getenv("NATS_PASSWORD")))
	natsConn, err := nats.Connect(os.Getenv("NATS_HOST"), nats.UserInfo(os.Getenv("NATS_USERNAME"), os.Getenv("NATS_PASSWORD")))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer natsConn.Close()

	orderService := order.NewOrderService(natsConn)

	if _, err := natsConn.QueueSubscribe("paymentCreated", "worker", func(m *nats.Msg) {
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
