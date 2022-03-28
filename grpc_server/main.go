package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc_server/pb"
	"log"
	"net"
)

const port = ":5001"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("hello")

	server := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(server, employeeService{})

	log.Println("gRCP Server started..." + port)
	server.Serve(listen)
}

type employeeService struct {
}

func (e employeeService) GetByNo(ctx context.Context, request *pb.GetByNoRequest) (*pb.EmployeeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e employeeService) GetAll(request *pb.GetAllRequest, server pb.EmployeeService_GetAllServer) error {
	//TODO implement me
	panic("implement me")
}

func (e employeeService) AddPhoto(server pb.EmployeeService_AddPhotoServer) error {
	//TODO implement me
	panic("implement me")
}

func (e employeeService) Save(ctx context.Context, request *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e employeeService) SaveAll(server pb.EmployeeService_SaveAllServer) error {
	//TODO implement me
	panic("implement me")
}
