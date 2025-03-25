package main

import (
	_ "github.com/lib/pq"

	bookPb "CategoryService/pb/BookService"
	categoryPb "CategoryService/pb/CategoryService"
	"CategoryService/services"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	dbConn := InitDb()

	bookServiceConn, err := ConnectBookService()

	if err != nil {
		log.Println("failed to connect to user server")
	}

	defer bookServiceConn.Close()

	bookClient := bookPb.NewBookClient(bookServiceConn)

	categoryPb.RegisterCategoryServer(grpcServer, &services.Server{Db: dbConn, BookClient: bookClient})

	fmt.Println("🚀 Category Service is running on port 50054...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func InitDb() *sql.DB {
	os.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	dbDsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbDsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

func ConnectBookService() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}
