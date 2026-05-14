package bloodrepo

const (
	queryCreateBlood = `INSERT INTO blood (id, user_id, hb, rbc, wbc, plt, hct, mcv, mch, esr, glucose, protein, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	queryGetBloodByID = `SELECT id, user_id, hb, rbc, wbc, plt, hct, mcv, mch, esr, glucose, protein, created_at FROM blood WHERE id = ?`

	queryGetBloodByUserID = `SELECT id, user_id, hb, rbc, wbc, plt, hct, mcv, mch, esr, glucose, protein, created_at FROM blood WHERE user_id = ?`

	queryUpdateBlood = `UPDATE blood SET user_id = ?, hb = ?, rbc = ?, wbc = ?, plt = ?, hct = ?, mcv = ?, mch = ?, esr = ?, glucose = ?, protein = ?, created_at = ? WHERE id = ?`

	queryDeleteBlood = `DELETE FROM blood WHERE id = ?`
)
