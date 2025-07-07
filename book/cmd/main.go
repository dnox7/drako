package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/dnox7/drako/contracts/gen/go/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authorService struct {
	client pb.AuthorServiceClient
	logger *log.Logger
}

func (s *authorService) GetAuthorByID(ctx context.Context, id int32) {
	author, err := s.client.GetAuthor(ctx, &pb.GetAuthorRequest{
		Id: id,
	})
	if err != nil {
		s.logger.Printf("failed to get author by id: %v", err)
	}
	fmt.Printf("%+v\n", author)
}

func (s *authorService) GetAllAuthors(ctx context.Context) {
	authors, _ := s.client.ListAuthors(ctx, nil)
	fmt.Printf("%+v\n", authors)
}

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthorServiceClient(conn)

	authorService := &authorService{
		client: client,
		logger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	authorService.GetAuthorByID(context.Background(), 1)
	authorService.GetAllAuthors(context.Background())
}
