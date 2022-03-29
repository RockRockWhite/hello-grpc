package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"grpc_client/pb"
	"io"
	"log"
	"os"
)

const port = ":5001"

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	client := pb.NewEmployeeServiceClient(conn)

	//fmt.Println(GetByNo(client, 1994))
	//fmt.Println(GetAll(client))
	//fmt.Println(addPhoto(client))
	saveAll(client)
}

func GetByNo(client pb.EmployeeServiceClient, no int32) *pb.EmployeeResponse {
	res, err := client.GetByNo(context.Background(), &pb.GetByNoRequest{No: no})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res
}

func GetAll(client pb.EmployeeServiceClient) []*pb.Employee {
	var employees []*pb.Employee
	stream, err := client.GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err.Error())
		}

		employees = append(employees, res.Employee)
	}

	return employees
}

func addPhoto(client pb.EmployeeServiceClient) *pb.AddPhotoResponse {
	imgFile, err := os.Open("fire.jpg")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer imgFile.Close()

	md := metadata.New(map[string]string{"no": "2004"})
	context := metadata.NewOutgoingContext(context.Background(), md)

	stream, _ := client.AddPhoto(context)

	for {
		chunk := make([]byte, 128*1024)
		chunkSize, err := imgFile.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err.Error())
		}
		if chunkSize < len(chunk) {
			chunk = chunk[:chunkSize]
		}
		stream.Send(&pb.AddPhotoRequest{Data: chunk})
	}
	res, _ := stream.CloseAndRecv()
	return res
}

func saveAll(client pb.EmployeeServiceClient) {
	stream, err := client.SaveAll(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, each := range employees {
		stream.Send(&pb.EmployeeRequest{Employee: &each})
	}

	finishChannel := make(chan struct{})
	// 异步接收
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err.Error())
			}
			fmt.Printf("received %v\n", res)
		}
		finishChannel <- struct{}{}
	}()

	stream.CloseSend()

	<-finishChannel
}
