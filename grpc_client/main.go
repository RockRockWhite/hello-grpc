package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_client/pb"
	"log"
)

const port = ":5001"

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	client := pb.NewEmployeeServiceClient(conn)

	fmt.Println(GetByNo(client, 1994))
	fmt.Println("hello")
}

func GetByNo(client pb.EmployeeServiceClient, no int32) *pb.EmployeeResponse {
	res, err := client.GetByNo(context.Background(), &pb.GetByNoRequest{No: no})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res
}
