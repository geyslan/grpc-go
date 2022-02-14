package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/geyslan/grpc-go/proto"
)

func main() {
	conn, err := grpc.Dial(":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CalculateValues(ctx, &pb.CalculationRequest{
		Operator: "add",
		Value1:   "111",
		Value2:   "555",
	})
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}
	log.Printf("Calculation: %s", r.GetResult())

	r, err = c.CalculateValues(ctx, &pb.CalculationRequest{
		Operator: "sub",
		Value1:   "555",
		Value2:   "111",
	})
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}
	log.Printf("Calculation: %s", r.GetResult())
}
