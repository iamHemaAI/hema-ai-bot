package services

import (
	"context"

	"github.com/iamHemaAI/hema-ai-bot/internal/domain"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain/dto"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain/mappers"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain/records"
)

type bloodRepository interface {
	CreateBlood(ctx context.Context, blood *records.Blood) error
	GetBloodByID(ctx context.Context, id string) (*records.Blood, error)
	GetBloodByUserID(ctx context.Context, userID string) ([]*records.Blood, error)
	UpdateBlood(ctx context.Context, blood *records.Blood) error
	DeleteBlood(ctx context.Context, id string) error
}

type bloodReportClient interface {
	Analyze(ctx context.Context, req *dto.BloodRequest) ([]byte, error)
}

type BloodService struct {
	r bloodRepository
	c bloodReportClient
}

func NewBloodService(r bloodRepository, c bloodReportClient) *BloodService {
	return &BloodService{r: r, c: c}
}

func (s *BloodService) CreateBloodAnalyze(ctx context.Context, blood *domain.Blood) ([]byte, error) {
	if err := s.r.CreateBlood(ctx, mappers.BloodToRecord(blood)); err != nil {
		return nil, err
	}

	return s.c.Analyze(ctx, &dto.BloodRequest{
		Hb:      blood.Hb.Value(),
		Rbc:     blood.Rbc.Value(),
		Wbc:     blood.Wbc.Value(),
		Plt:     blood.Plt.Value(),
		Hct:     blood.Hct.Value(),
		Mcv:     blood.Mcv.Value(),
		Mch:     blood.Mch.Value(),
		Esr:     blood.Esr.Value(),
		Glucose: blood.Glucose.Value(),
		Protein: blood.Protein.Value(),
	})
}
