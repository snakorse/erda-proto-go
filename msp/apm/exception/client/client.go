// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: exception.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/exception/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// ExceptionService exception.proto
	ExceptionService() pb.ExceptionServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		exceptionService: pb.NewExceptionServiceClient(cc),
	}
}

type serviceClients struct {
	exceptionService pb.ExceptionServiceClient
}

func (c *serviceClients) ExceptionService() pb.ExceptionServiceClient {
	return c.exceptionService
}

type exceptionServiceWrapper struct {
	client pb.ExceptionServiceClient
	opts   []grpc1.CallOption
}

func (s *exceptionServiceWrapper) GetExceptions(ctx context.Context, req *pb.GetExceptionsRequest) (*pb.GetExceptionsResponse, error) {
	return s.client.GetExceptions(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *exceptionServiceWrapper) GetExceptionEventIds(ctx context.Context, req *pb.GetExceptionEventIdsRequest) (*pb.GetExceptionEventIdsResponse, error) {
	return s.client.GetExceptionEventIds(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *exceptionServiceWrapper) GetExceptionEvent(ctx context.Context, req *pb.GetExceptionEventRequest) (*pb.GetExceptionEventResponse, error) {
	return s.client.GetExceptionEvent(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}