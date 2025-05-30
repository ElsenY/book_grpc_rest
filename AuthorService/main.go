package main

import (
	_ "github.com/lib/pq"

	authorPb "AuthorService/pb/AuthorService"
	userPb "AuthorService/pb/UserService"
	"AuthorService/services"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	dbConn := InitDb()

	userServiceConn, err := ConnectUserService()

	if err != nil {
		log.Println("failed to connect to user server")
	}

	defer userServiceConn.Close()

	userClient := userPb.NewUserClient(userServiceConn)

	authorPb.RegisterAuthorServer(grpcServer, &services.Server{Db: dbConn, UserClient: userClient})

	fmt.Println("🚀 Author Service is running on port 50052...")
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

func ConnectUserService() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(os.Getenv("USER_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}
