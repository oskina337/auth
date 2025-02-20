package service

import (
	pb "auth/pkg/userAPI_v1"
	"context"
	"fmt"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Server отвечает за обработку пользовательских запросов.
type Server struct {
	pb.UnimplementedUserAPIV1Server
}

// Create создаёт новый объект на основе запроса
func (s *Server) Create(_ context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	fmt.Printf("Create: %+v\n", req)
	return &pb.CreateResponse{Id: 1}, nil

}

// Get получает объект по идентификатору
func (s *Server) Get(_ context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Printf("Get: %+v\n", in)
	return &pb.GetResponse{
		Id:    in.Id,
		Name:  "Test User",
		Email: "test@example.com",
		Role:  pb.UserRole_USER,
	}, nil
}

// Update обновляет объект на основе запроса.
func (s *Server) Update(_ context.Context, in *pb.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Printf("Update: %+v\n", in)
	return &emptypb.Empty{}, nil
}

// Delete удаляет объект по идентификатору.
func (s *Server) Delete(_ context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Printf("Delete: %+v\n", in)
	return &emptypb.Empty{}, nil
}
