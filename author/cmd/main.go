package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/dnox7/drako/contracts/gen/go/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var authorsData = []*pb.Author{
	{
		Id:   1,
		Name: "Tom",
		Age:  12,
	},
	{
		Id:   2,
		Name: "Jerry",
		Age:  12,
	},
}

type authorServer struct {
	pb.UnimplementedAuthorServiceServer
	data []*pb.Author
}

func (s *authorServer) GetAuthor(ctx context.Context, req *pb.GetAuthorRequest) (*pb.GetAuthorResponse, error) {
	for _, author := range s.data {
		if author.Id == req.Id {
			return &pb.GetAuthorResponse{
				Author: author,
			}, nil
		}
	}
	return nil, errors.New("record not found")
}

func (s *authorServer) ListAuthors(ctx context.Context, _ *emptypb.Empty) (*pb.ListAuthorsResponse, error) {
	return &pb.ListAuthorsResponse{
		Authors: s.data,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthorServiceServer(grpcServer, &authorServer{
		data: authorsData,
	})
	grpcServer.Serve(lis)
}
