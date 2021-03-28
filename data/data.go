package data

import (
	"github.com/xylong/grpc-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var Employees = []pb.Employee{
	{
		Id:        1,
		No:        1,
		FirstName: "li",
		LastName:  "lin",
		MonthSalary: &pb.MonthSalary{
			Basic: 20000,
			Bouns: 9800.5,
		},
		Status: pb.EmployeeStatus_NORMAL,
		LastModfied: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	},
	{
		Id:        2,
		No:        2,
		FirstName: "静",
		LastName:  "静",
		MonthSalary: &pb.MonthSalary{
			Basic: 20000,
			Bouns: 9800.5,
		},
		Status: pb.EmployeeStatus_NORMAL,
		LastModfied: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	},
}
