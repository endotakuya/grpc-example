package main

import (
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "github.com/endotakuya/grpc-example/article"

	_ "github.com/go-sql-driver/mysql"
)

const (
	port   = ":50051"
	db     = "grpc_example"
	dbUser = "root"
	dbPass = ""
	dbHost = "db"
	dbPort = 3306
)

type server struct{}

func (s *server) First(ctx context.Context, in *pb.Empty) (*pb.Article, error) {
	db, _ := initDb()
	defer db.Close()

	article := pb.Article{}
	dist := []interface{}{&article.Id, &article.Title, &article.Content, &article.Status}
	err := db.QueryRow("SELECT * FROM articles ORDER BY id DESC LIMIT 1").Scan(dist...)
	return &article, err
}

func (s *server) Post(ctx context.Context, in *pb.Article) (*pb.Empty, error) {
	db, _ := initDb()
	defer db.Close()

	stmtIns, err := db.Prepare(fmt.Sprintf("INSERT INTO %s (title, content, status) VALUES (?, ?, ?)", "articles"))
	defer stmtIns.Close()

	_, err = stmtIns.Exec(in.Title, in.Content, in.Status)
	return &pb.Empty{}, err
}

func initDb() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, db)
	return sql.Open("mysql", dataSourceName)
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
