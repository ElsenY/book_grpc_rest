package services

const INSERT_CATEGORY_QUERY = `INSERT INTO category (name) VALUES ($1);`
const GET_CATEGORY_ID_BY_NAME = `SELECT id FROM category WHERE name = $1;`
const INSERT_BOOK_CATEGORY_QUERY = `INSERT INTO book_category (category_id,book_id) VALUES ($1,$2);`
