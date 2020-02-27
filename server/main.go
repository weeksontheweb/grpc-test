package main

import (
	"context"
	"grpc-test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Want server to implement the interface (AddServiceServer) that was 
//declared in the compiled proto file.
type server struct{}

func (s server) Add(ctx context.Context, request proto.Request) (proto.response, error)
{
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}	

func (s server) Multiply(ctx context.Context, request proto.Request) (proto.response, error)
{
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}	

func main() {

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		panic(err)
	)

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})

	//So we can serialise and deserialise.
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
