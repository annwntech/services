module github.com/micro/services/apps/service

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro/v2 v2.6.1-0.20200506095814-5387f73b5d8e
	github.com/micro/services/payments/provider v0.0.0-20200318105532-9c3078c484d5
)

replace github.com/micro/services/payments/provider => ../../payments/provider
