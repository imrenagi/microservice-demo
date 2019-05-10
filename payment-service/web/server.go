package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imrenagi/microservice-demo/payment-service/internal/payment"
)

//NewServer returns new http server
func NewServer(paymentService *payment.PaymentService) *Server {
	s := &Server{
		Router:         mux.NewRouter(),
		PaymentService: paymentService,
	}
	s.routes()

	return s
}

//Server is just a server
type Server struct {
	Router         *mux.Router
	PaymentService *payment.PaymentService
}

func (s Server) MakePayment() http.HandlerFunc {

	type Response struct {
		PaymentID string `json:"payment_id"`
		OrderID   string `json:"order_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var command payment.NewPaymentCommand
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&command)
		if err != nil {
			WriteFailResponse(w, http.StatusBadRequest, err)
			return
		}

		paymentID, err := s.PaymentService.MakePayment(command)
		if err != nil {
			log.Println(err)
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, Response{PaymentID: paymentID, OrderID: command.OrderID}, nil)
		return
	}
}

func (s Server) HC() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		WriteSuccessResponse(w, http.StatusOK, "Ok", nil)
	}
}

func (s Server) GetPayments() http.HandlerFunc {

	type Payment struct {
		PaymentID string  `json:"payment_id"`
		OrderID   string  `json:"order_id"`
		Value     float64 `json:"value"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		payments, err := s.PaymentService.GetPayment()
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		resp := make([]Payment, 0)
		for _, p := range payments {
			resp = append(resp, Payment{
				PaymentID: p.ID,
				OrderID:   p.OrderID,
				Value:     p.Value,
			})
		}

		WriteSuccessResponse(w, http.StatusOK, resp, nil)
		return
	}

}
