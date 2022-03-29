package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_server/pb"
	"io"
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

	for _, each := range employees {
		if each.No == request.No {
			return &pb.EmployeeResponse{Employee: &each}, nil
		}
	}
	return nil, errors.New("employee not found")
}

func (e employeeService) GetAll(request *pb.GetAllRequest, server pb.EmployeeService_GetAllServer) error {
	for _, each := range employees {
		server.Send(&pb.EmployeeResponse{
			Employee: &each,
		})
	}

	return nil
}

func (e employeeService) AddPhoto(server pb.EmployeeService_AddPhotoServer) error {
	md, ok := metadata.FromIncomingContext(server.Context())
	if ok {
		fmt.Printf("Employee: %s\n", md["no"][0])
	}

	img := []byte{}
	for {
		data, err := server.Recv()
		if err == io.EOF {
			fmt.Printf("File Size:%d \n", len(img))
			break
		}
		if err != nil {
			return err
		}

		fmt.Printf("File received %d", len(data.Data))
		img = append(img, data.Data...)
		fmt.Println(data.Data)
		fmt.Println("========================")
	}

	return nil
}

func (e employeeService) Save(ctx context.Context, request *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e employeeService) SaveAll(server pb.EmployeeService_SaveAllServer) error {
	for {
		req, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		employees = append(employees, *req.Employee)
		server.Send(&pb.EmployeeResponse{Employee: req.Employee})
	}

	for _, each := range employees {
		fmt.Println(each)
	}

	return nil
}
