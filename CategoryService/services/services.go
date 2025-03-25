package services

import (
	bookPb "CategoryService/pb/BookService"
	categoryPb "CategoryService/pb/CategoryService"
	"context"
	"database/sql"
	"log"
)

type Server struct {
	categoryPb.UnimplementedCategoryServer
	BookClient bookPb.BookClient
	Db         *sql.DB
}

func (cs *Server) InsertCategory(ctx context.Context, req *categoryPb.InsertCategoryRequest) (resp *categoryPb.InsertCategoryResponse, err error) {

	_, err = cs.Db.Exec(INSERT_CATEGORY_QUERY, req.Name)

	if err != nil {
		log.Println(err)
		return
	}

	return &categoryPb.InsertCategoryResponse{Message: "Success insert category"}, err
}

func (cs *Server) LinkBookWithCategory(ctx context.Context, req *categoryPb.LinkBookWithCategoryRequest) (resp *categoryPb.LinkBookWithCategoryResponse, err error) {

	bookResp, err := cs.BookClient.GetBookIdByTitle(ctx, &bookPb.GetBookIdByTitleRequest{Title: req.BookTitle})

	if err != nil {
		log.Println(err)
		return
	}

	row := cs.Db.QueryRow(GET_CATEGORY_ID_BY_NAME, req.Name)

	var categoryId string

	err = row.Scan(&categoryId)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = cs.Db.Exec(INSERT_BOOK_CATEGORY_QUERY, categoryId, bookResp.BookId)

	if err != nil {
		log.Println(err)
		return
	}

	return &categoryPb.LinkBookWithCategoryResponse{Message: "Success link category with book"}, err
}
