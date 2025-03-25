package services

import (
	authorPb "BookService/pb/AuthorService"
	bookPb "BookService/pb/BookService"
	userPb "BookService/pb/UserService"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Server struct {
	bookPb.UnimplementedBookServer
	AuthorClient authorPb.AuthorClient
	UserClient   userPb.UserClient
	Db           *sql.DB
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

	var returnDate sql.NullTime
	row = s.Db.QueryRow(CHECK_BOOK_RETURNED_QUERY, userResp.Id, bookId)

	err = row.Scan(&returnDate)

	if err != nil {
		log.Println(err)
		return
	}

	if returnDate.Valid {
		return &bookPb.ReturnBookResponse{Message: "failed to return book because Book has been returned"}, err
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

func (s *Server) GetBookIdByTitle(ctx context.Context, req *bookPb.GetBookIdByTitleRequest) (resp *bookPb.GetBookIdByTitleResponse, err error) {

	var bookId string
	var stock int
	row := s.Db.QueryRow(GET_BOOK_DATA_BY_TITLE_QUERY, req.Title)

	err = row.Scan(&bookId, &stock)

	if err != nil {
		log.Println(err)
		return
	}

	return &bookPb.GetBookIdByTitleResponse{BookId: bookId}, err
}

// recommend book that has been borrowed the most
func (s *Server) RecommendBook(ctx context.Context, req *bookPb.RecommendBookRequest) (resp *bookPb.RecommendBookResponse, err error) {

	var bookTitle string
	var borrowCount int
	row := s.Db.QueryRow(GET_MOST_BORROWED_BOOK_QUERY)

	err = row.Scan(&bookTitle, &borrowCount)

	if err != nil {
		log.Println(err)
		return
	}

	return &bookPb.RecommendBookResponse{BookTitle: fmt.Sprintf("Recommend %s, has been borrowed the most %d", bookTitle, borrowCount)}, err
}

func (s *Server) SearchBook(ctx context.Context, req *bookPb.SearchBookRequest) (resp *bookPb.SearchBookResponse, err error) {

	var id string
	var stock int32
	row := s.Db.QueryRow(GET_BOOK_DATA_BY_TITLE_QUERY, req.BookTitle)

	err = row.Scan(&id, &stock)

	if err != nil {
		if err == sql.ErrNoRows {
			return &bookPb.SearchBookResponse{Message: "No book with the provided title found!"}, nil
		}

		log.Println(err)
		return
	}

	return &bookPb.SearchBookResponse{BookId: id, BookTitle: req.BookTitle, Stock: stock}, err
}

func (s *Server) EditBookStock(ctx context.Context, req *bookPb.EditBookStockRequest) (resp *bookPb.EditBookStockResponse, err error) {
	var bookId string
	var stock int
	row := s.Db.QueryRow(GET_BOOK_DATA_BY_TITLE_QUERY, req.Title)

	err = row.Scan(&bookId, &stock)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = s.Db.Exec(UPDATE_BOOK_STOCK_QUERY, req.Stock, bookId)

	if err != nil {
		log.Println(err)
		return
	}

	return &bookPb.EditBookStockResponse{Message: "Success update book stock"}, err
}
