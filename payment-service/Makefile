.PHONY: proto test docker

proto:
	protoc --go_out=plugins=grpc:.  --micro_out=. pkg/proto/order/order.proto
	protoc --go_out=plugins=grpc:.  --micro_out=. pkg/proto/payment/payment.proto

