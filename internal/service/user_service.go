package service

import (
	pb "auth/pkg/userAPI_v1"
	"context"
	"fmt"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedUserAPIV1Server
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	fmt.Printf("Create: %+v\n", req)
	return &pb.CreateResponse{Id: 1}, nil

}

func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Printf("Get: %+v\n", in)
	return &pb.GetResponse{
		Id:    in.Id,
		Name:  "Test User",
		Email: "test@example.com",
		Role:  pb.UserRole_USER,
	}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Printf("Update: %+v\n", in)
	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Printf("Delete: %+v\n", in)
	return &emptypb.Empty{}, nil
}
