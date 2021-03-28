package main

import (
	"github.com/xylong/grpc-server/pb"
	"github.com/xylong/grpc-server/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":9000"
)

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

	server := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(server, new(service.EmployeeService))
	log.Printf("gRPC server starting in port:%d", 9000)

	if err := server.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
