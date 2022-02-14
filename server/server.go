package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/geyslan/grpc-go/proto"
)

type server struct {
	pb.UnimplementedCalculationServer
}

func (s *server) CalculateValues(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationReply, error) {
	log.Printf("%v, %v, %v", in.GetOperator(), in.GetValue1(), in.GetValue2())

	intValue1, err := strconv.Atoi(in.GetValue1())
	if err != nil {
		log.Fatalf("Error converting a int: %v", err)
	}
	intValue2, err := strconv.Atoi(in.GetValue2())
	if err != nil {
		log.Fatalf("Error converting a int: %v", err)
	}

	var result int
	switch in.GetOperator() {
	case "add":
		result = intValue1 + intValue2
	case "sub":
		result = intValue1 - intValue2
	case "mul":
		result = intValue1 * intValue2
	case "div":
		result = intValue1 / intValue2
	}

	return &pb.CalculationReply{
		Result: strconv.Itoa(result),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculationServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
