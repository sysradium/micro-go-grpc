package micro

import (
	"github.com/micro/go-micro"
	broker "github.com/micro/go-micro/broker"
	client "github.com/micro/go-micro/client/grpc"
	server "github.com/micro/go-micro/server/grpc"
)

func NewService(opts ...micro.Option) micro.Service {
	// our grpc client
	c := client.NewClient()
	// our grpc server
	s := server.NewServer()
	// our grpc broker
	b := broker.NewBroker()

	// create options with priority for our opts
	options := []micro.Option{
		micro.Client(c),
		micro.Server(s),
		micro.Broker(b),
	}

	// append passed in opts
	options = append(options, opts...)

	// generate and return a service
	return micro.NewService(options...)
}
