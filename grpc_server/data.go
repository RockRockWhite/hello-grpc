package main

import "grpc_server/pb"

var employees = []pb.Employee{
	{
		Id:        1,
		No:        1994,
		FirstName: "Hal",
		LastName:  "Bin",
		Salary:    2333,
	}, {
		Id:        2,
		No:        1995,
		FirstName: "R",
		LastName:  "Ral",
		Salary:    0,
	},
}
