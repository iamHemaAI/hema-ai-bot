package services

import (
	"context"

	"github.com/iamHemaAI/hema-ai-bot/internal/domain"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain/mappers"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain/records"
	"github.com/iamHemaAI/hema-ai-bot/pkg/apperrors"
)

type userRepository interface {
	CreateUser(ctx context.Context, user *records.User) error
	GetUserByTgID(ctx context.Context, tgID int64) (*records.User, error)
	UpdateUser(ctx context.Context, user *records.User) error
	DeleteUser(ctx context.Context, tgID string) error
}

type UserService struct {
	r userRepository
}

func NewUserService(r userRepository) *UserService {
	return &UserService{r: r}
}

func (s *UserService) Register(ctx context.Context, user *domain.User) error {
	u, err := s.r.GetUserByTgID(ctx, user.TgID)
	if err != nil {
		return err
	}
	if u != nil {
		return apperrors.ErrUserAlreadyExists
	}

	err = s.r.CreateUser(ctx, mappers.UserToRecord(user))
	if err != nil {
		return err
	}

	return nil
}
