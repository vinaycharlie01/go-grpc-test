package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vinaycharlie07/ai/aitech/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataServiceClient(conn)

	stream, err := client.SendData(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling SendData: %v", err)
	}

	for {
		data, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error while receiving data: %v", err)
		}
		fmt.Println(data.GetMessage())
	}
}
