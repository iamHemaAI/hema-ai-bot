package userrepo

const (
	queryCreateUser = `INSERT INTO users (tg_id, tg_username, name) VALUES (?, ?, ?)`

	queryUpdateUser = `UPDATE users SET tg_username = ?, name = ? WHERE tg_id = ?`

	queryDeleteUser = `DELETE FROM users WHERE tg_id = ?`

	queryGetUserByTgID = `SELECT tg_id, tg_username, name FROM users WHERE tg_id = ?`
)
