package services

const INSERT_USER_QUERY = `INSERT INTO users (name,password,email) VALUES ($1,$2,$3);`
const GET_USER_PASSWORD_BY_EMAIL_QUERY = `SELECT password FROM users WHERE email = $1;`
const GET_USER_ID_BY_EMAIL_QUERY = `SELECT id FROM users WHERE email = $1;`
