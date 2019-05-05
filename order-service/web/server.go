package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imrenagi/microservice-demo/order-service/internal/order"
)

//NewServer returns new http server
func NewServer() *Server {
	s := &Server{
		Router:       mux.NewRouter(),
		OrderService: order.NewOrderService(),
	}
	s.routes()

	return s
}

//Server is just a server
type Server struct {
	Router       *mux.Router
	OrderService *order.OrderService
}

func (s Server) CreateOrder() http.HandlerFunc {

	type Response struct {
		ID string `json:"order_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var command order.NewOrderCommand
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&command)
		if err != nil {
			WriteFailResponse(w, http.StatusBadRequest, err)
			return
		}

		id, err := s.OrderService.NewOrder(command)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err)
			return
		}

		WriteSuccessResponse(w, http.StatusOK, &Response{ID: id}, nil)
		return
	}
}

func (s Server) GetOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		orders, err := s.OrderService.GetOrders()
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, nil)
			return
		}

		resp := make([]Order, 0)
		for _, o := range orders {
			order := Order{
				CustomerID:    o.CustomerID,
				ID:            o.ID,
				LineItems:     o.LineItems,
				Price:         o.Price,
				PaymentStatus: o.PaymentStatus,
			}
			resp = append(resp, order)
		}

		WriteSuccessResponse(w, http.StatusOK, resp, nil)
		return
	}
}
