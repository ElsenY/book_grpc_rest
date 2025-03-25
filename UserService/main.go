package main

import (
	userPb "UserService/pb/UserService"
	"UserService/services"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	dbConn := InitDb()

	userPb.RegisterUserServer(grpcServer, &services.Server{Db: dbConn})

	fmt.Println("🚀 User Service is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func InitDb() *sql.DB {
	dbDsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbDsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
