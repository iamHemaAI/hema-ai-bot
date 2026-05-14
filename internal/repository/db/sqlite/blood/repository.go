package bloodrepo

import (
	"context"
	"database/sql"

	"github.com/iamHemaAI/hema-ai-bot/internal/domain/records"
)

type BloodRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BloodRepository {
	return &BloodRepository{db: db}
}

func (r *BloodRepository) CreateBlood(ctx context.Context, blood *records.Blood) error {
	_, err := r.db.ExecContext(ctx, queryCreateBlood,
		blood.ID, blood.UserID, blood.Hb, blood.Rbc, blood.Wbc, blood.Plt, blood.Hct,
		blood.Mcv, blood.Mch, blood.Esr, blood.Glucose, blood.Protein, blood.CreatedAt,
	)
	return err
}

func (r *BloodRepository) GetBloodByID(ctx context.Context, id string) (*records.Blood, error) {
	row := r.db.QueryRowContext(ctx, queryGetBloodByID, id)

	var blood records.Blood
	err := row.Scan(
		&blood.ID, &blood.UserID, &blood.Hb, &blood.Rbc, &blood.Wbc, &blood.Plt,
		&blood.Hct, &blood.Mcv, &blood.Mch, &blood.Esr, &blood.Glucose, &blood.Protein, &blood.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &blood, nil
}

func (r *BloodRepository) GetBloodByUserID(ctx context.Context, userID string) ([]*records.Blood, error) {
	rows, err := r.db.QueryContext(ctx, queryGetBloodByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*records.Blood
	for rows.Next() {
		var blood records.Blood
		err := rows.Scan(
			&blood.ID, &blood.UserID, &blood.Hb, &blood.Rbc, &blood.Wbc, &blood.Plt,
			&blood.Hct, &blood.Mcv, &blood.Mch, &blood.Esr, &blood.Glucose, &blood.Protein, &blood.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, &blood)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *BloodRepository) UpdateBlood(ctx context.Context, blood *records.Blood) error {
	_, err := r.db.ExecContext(ctx, queryUpdateBlood,
		blood.UserID, blood.Hb, blood.Rbc, blood.Wbc, blood.Plt, blood.Hct,
		blood.Mcv, blood.Mch, blood.Esr, blood.Glucose, blood.Protein, blood.CreatedAt, blood.ID,
	)
	return err
}

func (r *BloodRepository) DeleteBlood(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, queryDeleteBlood, id)
	return err
}
