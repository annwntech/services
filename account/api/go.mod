module github.com/micro/services/account/api

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro/v2 v2.6.1-0.20200506095814-5387f73b5d8e
	github.com/micro/services/account/invite v0.0.0-20200421094732-38d776e22810
	github.com/micro/services/payments/provider v0.0.0-20200421094732-38d776e22810
	github.com/micro/services/projects/invite v0.0.0-00010101000000-000000000000
	github.com/micro/services/users/service v0.0.0-20200421094732-38d776e22810
)

replace github.com/micro/services/payments/provider => ../../payments/provider

replace github.com/micro/services/projects/service => ../../projects/service

replace github.com/micro/services/projects/invite => ../../projects/invite
