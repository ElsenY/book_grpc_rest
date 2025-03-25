package services

import (
	authorPb "BookService/pb/AuthorService"
	bookPb "BookService/pb/BookService"
	userPb "BookService/pb/UserService"
	"context"
	"database/sql"
	"log"

	"github.com/golang-jwt/jwt"
)

type Server struct {
	bookPb.UnimplementedBookServer
	AuthorClient authorPb.AuthorClient
	UserClient   userPb.UserClient
	Db           *sql.DB
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (s *Server) InsertBook(ctx context.Context, req *bookPb.InsertBookRequest) (resp *bookPb.InsertBookResponse, err error) {
	userResp, err := s.UserClient.GetUserIdByEmail(ctx, &userPb.GetUserIdByEmailRequest{Email: req.AuthorEmail})

	if err != nil {
		log.Println(err)
		return
	}

	authorResp, err := s.AuthorClient.GetAuthorIdByUserId(ctx, &authorPb.GetAuthorIdByUserIdRequest{UserId: userResp.Id})

	if err != nil {
		log.Println(err)
		return
	}

	_, err = s.Db.Exec(INSERT_BOOK_QUERY, req.Title, req.Stock, authorResp.AuthorId)

	if err != nil {
		log.Println(err)
		return
	}

	return &bookPb.InsertBookResponse{Message: "Success insert Book"}, err
}

func (s *Server) BorrowBook(ctx context.Context, req *bookPb.BorrowBookRequest) (resp *bookPb.BorrowBookResponse, err error) {
	userResp, err := s.UserClient.GetUserIdByEmail(ctx, &userPb.GetUserIdByEmailRequest{Email: req.BorrowerEmail})

	if err != nil {
		log.Println(err)
		return
	}

	var bookId string
	var stock int
	row := s.Db.QueryRow(GET_BOOK_DATA_BY_TITLE_QUERY, req.Title)

	err = row.Scan(&bookId, &stock)

	if err != nil {
		log.Println(err)
		return
	}

	if stock <= 0 {
		log.Println("book stock not enough for borrowing")
		return
	}

	tx, err := s.Db.Begin()

	if err != nil {
		log.Println(err)
		return
	}

	_, err = tx.Exec(UPDATE_BOOK_STOCK_QUERY, stock-1, bookId)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}

	_, err = tx.Exec(INSERT_BORROW_BOOK_QUERY, userResp.Id, bookId)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}

	err = tx.Commit()

	if err != nil {
		log.Println(err)
		return
	}

	return &bookPb.BorrowBookResponse{Message: "Success borrow book"}, err
}

func (s *Server) ReturnBook(ctx context.Context, req *bookPb.ReturnBookRequest) (resp *bookPb.ReturnBookResponse, err error) {
	userResp, err := s.UserClient.GetUserIdByEmail(ctx, &userPb.GetUserIdByEmailRequest{Email: req.ReturnerEmail})

	if err != nil {
		log.Println(err)
		return
	}

	var bookId string
	var stock int
	row := s.Db.QueryRow(GET_BOOK_DATA_BY_TITLE_QUERY, req.Title)

	err = row.Scan(&bookId, &stock)

	if err != nil {
		log.Println(err)
		return
	}

	tx, err := s.Db.Begin()

	if err != nil {
		log.Println(err)
		return
	}

	_, err = tx.Exec(UPDATE_BOOK_STOCK_QUERY, stock+1, bookId)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}

	result, err := tx.Exec(RETURN_BOOK_QUERY, userResp.Id, bookId)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}

	// Check if stock was updated (no rows affected = user is not borrowing the book)
	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		tx.Rollback()
		log.Println("user has not borrowed this book")
		return
	}

	err = tx.Commit()

	if err != nil {
		log.Println(err)
		return
	}

	return &bookPb.ReturnBookResponse{Message: "Success return book"}, err
}
