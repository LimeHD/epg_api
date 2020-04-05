package main

import (
	pb "github.com/LimeHD/epg_api/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterEpgServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	n := 0
	rune := make([]rune, len(request.Message))

	for _, r := range request.Message {
		rune[n] = r
		n++
	}

	rune = rune[0:n]

	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	output := string(rune)
	response = &pb.Response{
		Message: output,
	}

	return response, nil
}
