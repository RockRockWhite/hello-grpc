package main

import "grpc_client/pb"

var employees = []pb.Employee{
	{
		Id:        3,
		No:        1996,
		FirstName: "Hal",
		LastName:  "Bin",
		Salary:    2333,
	}, {
		Id:        4,
		No:        1997,
		FirstName: "R",
		LastName:  "Ral",
		Salary:    0,
	},
}
