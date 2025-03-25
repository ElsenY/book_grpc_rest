package services

import (
	authorPb "AuthorService/pb/AuthorService"
	userPb "AuthorService/pb/UserService"
	"context"
	"database/sql"
	"log"
)

type Server struct {
	authorPb.UnimplementedAuthorServer
	UserClient userPb.UserClient
	Db         *sql.DB
}

func (s *Server) RegisterUserAsAuthor(ctx context.Context, req *authorPb.RegisterUserAsAuthorRequest) (resp *authorPb.RegisterUserAsAuthorResponse, err error) {
	userId, err := s.UserClient.GetUserIdByEmail(ctx, &userPb.GetUserIdByEmailRequest{Email: req.UserEmail})

	if err != nil {
		log.Println(err)
		return
	}

	_, err = s.Db.Exec(INSERT_AUTHOR_QUERY, userId.Id)

	if err != nil {
		log.Println(err)
		return
	}

	return &authorPb.RegisterUserAsAuthorResponse{Message: "Success insert Author"}, err
}

func (s *Server) GetAuthorIdByUserId(ctx context.Context, req *authorPb.GetAuthorIdByUserIdRequest) (resp *authorPb.GetAuthorIdByUserIdResponse, err error) {

	var authorId string
	row := s.Db.QueryRow(GET_AUTHOR_ID_BY_USER_ID_QUERY, req.UserId)

	err = row.Scan(&authorId)

	if err != nil {
		log.Println(err)
		return
	}

	return &authorPb.GetAuthorIdByUserIdResponse{AuthorId: authorId}, err
}
