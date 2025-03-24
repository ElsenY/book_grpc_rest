package services

import (
	userPb "UserService/pb/UserService"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	userPb.UnimplementedUserServer
	Db *sql.DB
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (s *Server) RegisterUser(ctx context.Context, req *userPb.RegisterUserRequest) (resp *userPb.RegisterUserResponse, err error) {

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

	return &userPb.RegisterUserResponse{Message: "User created!"}, nil
}

func (s *Server) GetUserIdByEmail(ctx context.Context, req *userPb.GetUserIdByEmailRequest) (resp *userPb.GetUserIdByEmailResponse, err error) {

	row := s.Db.QueryRow(GET_USER_ID_BY_EMAIL_QUERY, req.Email)

	var id string

	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with that ID")
		} else {
			fmt.Println("There is an error while scanning row", err)
		}
		return
	}

	return &userPb.GetUserIdByEmailResponse{Id: id}, nil
}

func (s *Server) Login(ctx context.Context, req *userPb.LoginRequest) (resp *userPb.LoginResponse, err error) {

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
		Email: req.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return
	}

	return &userPb.LoginResponse{Token: tokenString}, nil
}
