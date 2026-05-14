package records

type User struct {
	TgID       int64  `db:"tg_id"`
	TgUsername string `db:"tg_username"`
	Name       string `db:"name"`
}
