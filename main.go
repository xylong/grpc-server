package main

import (
	"errors"
	"github.com/xylong/grpc-server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":9000"
)

type employeeService struct {
}

func (s *employeeService) GetByNo(ctx context.Context, req *pb.GetByNoRequest) (*pb.EmployeeResponse, error) {
	for _, e := range employees {
		if req.No == e.No {
			return &pb.EmployeeResponse{
				Employee: &e,
			}, nil
		}
	}
	return nil, errors.New("employee not found")
}

func (s *employeeService) GetAll(req *pb.GetAllRequest, stream pb.EmployeeService_GetAllServer) error {
	for _, e := range employees {
		stream.Send(&pb.EmployeeResponse{
			Employee: &e,
		})

		time.Sleep(time.Second * 2) // 模拟streaming
	}
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

/*
func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	creds, err := credentials.NewServerTLSFromFile("server.pem", "server.key")
	if err != nil {
		log.Fatal(err.Error())
	}
	options := []grpc.ServerOption{grpc.Creds(creds)}
	server := grpc.NewServer(options...)

	pb.RegisterEmployeeServiceServer(server, new(employeeService))

	log.Printf("gRPC server starting in port:%d", 9000)
	server.Serve(listener)
}
*/
func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	rpcServer := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(rpcServer, new(employeeService))
	log.Printf("gRPC server starting in port:%d", 9000)
	rpcServer.Serve(listener)
}
