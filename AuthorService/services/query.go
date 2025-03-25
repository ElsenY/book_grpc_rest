package services

const INSERT_AUTHOR_QUERY = `INSERT INTO authors (user_id) VALUES ($1);`
const GET_AUTHOR_ID_BY_USER_ID_QUERY = `SELECT id FROM authors WHERE user_id = $1;`
