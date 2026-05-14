package records

type Blood struct {
	ID     string `db:"id"`
	UserID int64  `db:"user_id"`

	Hb      int     `db:"hb"`
	Rbc     float64 `db:"rbc"`
	Wbc     float64 `db:"wbc"`
	Plt     int     `db:"plt"`
	Hct     int     `db:"hct"`
	Mcv     int     `db:"mcv"`
	Mch     int     `db:"mch"`
	Esr     int     `db:"esr"`
	Glucose float64 `db:"glucose"`
	Protein int     `db:"protein"`

	CreatedAt string `db:"created_at"`
}
