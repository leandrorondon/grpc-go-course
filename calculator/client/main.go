package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//doSum(c)
	//doPrimes(c, 120)
	//doAvg(c, []int32{1, 2, 3, 4})
	//doMax(c, []int32{1, 5, 3, 6, 2, 20})
	doSqrt(c, 10)
	doSqrt(c, 36)
	doSqrt(c, -10)
}
