package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "github.com/endotakuya/grpc-exsample/article"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) First(ctx context.Context, in *pb.Empty) (*pb.Article, error) {
	return &pb.Article{Id: 1, Title: "タイトル", Content: "本文", Status: pb.Article_PUBLISH}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on port", port)
	s := grpc.NewServer()
	pb.RegisterArticleServiceServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
