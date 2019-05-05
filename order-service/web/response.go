package web

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	ID            string   `json:"id"`
	CustomerID    string   `json:"customer_id"`
	LineItems     []string `json:"line_iterms"`
	Price         float64  `json:"price"`
	PaymentStatus string   `json:"payment_status"`
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}, headMap map[string]string) {
	w.Header().Add("Content-Type", "application/json")
	if headMap != nil && len(headMap) > 0 {
		for key, val := range headMap {
			w.Header().Add(key, val)
		}
	}
	w.WriteHeader(statusCode)
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
	//TODO handle error in marshalling and writing bytes
}

func WriteFailResponse(w http.ResponseWriter, statusCode int, error interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonData, _ := json.Marshal(error)
	w.Write(jsonData)
	//TODO handle error in marshalling and writing bytes
}
