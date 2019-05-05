package web

func (s Server) routes() {
	s.Router.HandleFunc("/orders", s.CreateOrder()).Methods("POST")
	s.Router.HandleFunc("/orders", s.GetOrders()).Methods("GET")
}
