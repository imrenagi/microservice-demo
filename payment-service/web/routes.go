package web

func (s Server) routes() {
	s.Router.HandleFunc("/payments/", s.MakePayment()).Methods("POST")
	s.Router.HandleFunc("/payments/", s.GetPayments()).Methods("GET")
}
