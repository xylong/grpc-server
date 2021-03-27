package main

import (
	"github.com/xylong/grpc-server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	port = ":9000"
)

type employeeService struct {
}

func (s *employeeService) GetByNo(context.Context, *pb.GetByNoRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}

func (s *employeeService) GetAll(*pb.GetAllRequest, pb.EmployeeService_GetAllServer) error {
	return nil
}
func (s *employeeService) AddPhoto(pb.EmployeeService_AddPhotoServer) error {
	return nil
}

func (s *employeeService) Save(context.Context, *pb.EmpolyeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}
func (s *employeeService) SaveAll(pb.EmployeeService_SaveAllServer) error {
	return nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	creds, err := credentials.NewServerTLSFromFile("cret.pem", "key.pem")
	if err != nil {
		log.Fatal(err.Error())
	}
	options := []grpc.ServerOption{grpc.Creds(creds)}
	server := grpc.NewServer(options...)

	pb.RegisterEmployeeServiceServer(server, new(employeeService))

	log.Printf("gRPC server starting in port:%d", 9000)
	server.Serve(listener)
}
