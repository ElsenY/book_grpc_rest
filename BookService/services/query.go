package services

const INSERT_BOOK_QUERY = `INSERT INTO books(title,stock,author_id) VALUES ($1,$2,$3);`
const GET_BOOK_DATA_BY_TITLE_QUERY = `SELECT id,stock FROM books WHERE title = $1;`
const UPDATE_BOOK_STOCK_QUERY = `UPDATE books SET stock = $1 WHERE id = $2;`

const INSERT_BORROW_BOOK_QUERY = `INSERT INTO borrow_records(user_id,book_id) VALUES ($1,$2);`
const RETURN_BOOK_QUERY = `UPDATE borrow_records SET return_date = CURRENT_TIMESTAMP WHERE user_id = $1 AND book_id = $2;`
