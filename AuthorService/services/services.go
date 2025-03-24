package services

import (
	authorPb "AuthorService/pb/AuthorService"
	userPb "AuthorService/pb/UserService"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	authorPb.UnimplementedAuthorServer
	UserClient userPb.UserClient
	Db         *sql.DB
}

type Claims struct {
	email string `json:"email"`
	jwt.StandardClaims
}

func (s *Server) RegisterUserAsAuthor(ctx context.Context, req *authorPb.RegisterUserAsAuthorRequest) (resp *authorPb.RegisterUserAsAuthorResponse, err error) {

	jwtContent, err := extractJWT(ctx)

	if err != nil {
		log.Println(err)
		return
	}

	userId, err := s.UserClient.GetUserIdByEmail(ctx, &userPb.GetUserIdByEmailRequest{Email: jwtContent["email"].(string)})

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(userId)
	_, err = s.Db.Exec(INSERT_AUTHOR_QUERY, userId.Id)

	if err != nil {
		log.Println(err)
		return
	}

	return &authorPb.RegisterUserAsAuthorResponse{Message: "Success insert Author"}, err
}

func extractJWT(ctx context.Context) (map[string]interface{}, error) {
	// Get metadata from gRPC request
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	// Extract the token from "authorization" header
	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return nil, errors.New("authorization token not provided")
	}

	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")

	// Parse and validate JWT
	claims, err := validateJWT(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func validateJWT(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
