package main

import (
	"fmt"
	"time"

	"net"

	pb "github.com/vinaycharlie07/ai/aitech/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDataServiceServer
}

func GeneEle(ch chan []int64) {
	defer close(ch)
	elements := make([]int64, 1000)
	for i := int64(0); i < 1000; i++ {
		elements[i] = i + 1
	}

	// Send 5 ele every 5 min
	b := 5
	for i := 0; i < len(elements); i += b {

		end := i + b
		if end > len(elements) {
			end = len(elements)
		}
		ch <- elements[i:end]
		// batch := elements[i:end]
	}
}

func (s *server) SendData(req *pb.Empty, stream pb.DataService_SendDataServer) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	ch := make(chan []int64)
	go GeneEle(ch)
	for {

		select {
		case <-ticker.C:
			response := &pb.DataResponse{
				Message:   <-ch,
				Timestamp: time.Now().Unix(),
			}
			if err := stream.Send(response); err != nil {
				return err
			}
		case <-stream.Context().Done():

			return nil

		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataServiceServer(grpcServer, &server{})

	fmt.Println("running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
