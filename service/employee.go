package service

import (
	"errors"
	"fmt"
	"github.com/xylong/grpc-server/data"
	"github.com/xylong/grpc-server/helper"
	"github.com/xylong/grpc-server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"io"
	"time"
)

// EmployeeService 员工
type EmployeeService struct{}

// GetByNo 根据编号查找
func (s *EmployeeService) GetByNo(ctx context.Context, request *pb.GetByNoRequest) (*pb.EmployeeResponse, error) {
	for _, e := range data.Employees {
		if request.No == e.No {
			return &pb.EmployeeResponse{
				Employee: &e,
			}, nil
		}
	}

	return nil, errors.New("not found")
}

// GetAll 获取全部
func (s *EmployeeService) GetAll(request *pb.GetAllRequest, server pb.EmployeeService_GetAllServer) error {
	for _, e := range data.Employees {

		server.Send(&pb.EmployeeResponse{
			Employee: &e,
		})

		time.Sleep(time.Second * 2)
	}

	return nil
}

// AddPhoto 添加头像
func (s *EmployeeService) AddPhoto(server pb.EmployeeService_AddPhotoServer) error {
	md, ok := metadata.FromIncomingContext(server.Context())
	if ok {
		fmt.Printf("Employee: %s\n", md["no"][0])

		img := []byte{}
		for {
			data, err := server.Recv()

			if err == io.EOF {
				if err := helper.Save("avatar", img); err != nil {
					return err
				} else {
					return server.SendAndClose(&pb.AddPhotoResponse{
						Ok: ok,
					})
				}
			}

			if err != nil {
				return err
			}

			fmt.Printf("received %d\n", len(data.Data))
			img = append(img, data.Data...)
		}
	}

	return nil
}

func (s *EmployeeService) Save(ctx context.Context, request *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}

func (s *EmployeeService) SaveAll(server pb.EmployeeService_SaveAllServer) error {
	return nil
}
