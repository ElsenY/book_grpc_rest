package services

import (
	"UserService/pb"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	pb.UnimplementedUserServer
	Db *sql.DB
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (resp *pb.RegisterUserResponse, err error) {

	bcryptPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error hashing password:", err)
		return
	}

	_, err = s.Db.Exec(INSERT_USER_QUERY, req.Name, bcryptPass, req.Email)

	if err != nil {
		log.Printf("Error hashing password:", err)
		return
	}

	return &pb.RegisterUserResponse{Message: "User created!"}, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{Message: "John Doe"}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (resp *pb.LoginResponse, err error) {

	row := s.Db.QueryRow(GET_USER_PASSWORD_BY_EMAIL_QUERY, req.Email)

	var bcryptPass string

	err = row.Scan(&bcryptPass)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with that ID")
		} else {
			fmt.Println("There is an error while scanning row", err)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(bcryptPass), []byte(req.Password))

	if err != nil {
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: req.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// TODO : remove the secret from code!
	jwtSecretKey := "test123"

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return
	}

	return &pb.LoginResponse{Token: tokenString}, nil
}
