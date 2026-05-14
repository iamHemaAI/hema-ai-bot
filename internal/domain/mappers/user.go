package mappers

import (
	"github.com/iamHemaAI/hema-ai-bot/internal/domain"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain/records"
	"github.com/iamHemaAI/hema-ai-bot/internal/tgbot/tgdto"
)

func UserToRecord(user *domain.User) *records.User {
	return &records.User{
		TgID:       user.TgID,
		TgUsername: user.TgUsername,
		Name:       user.Name,
	}
}

func RecordToUser(record *records.User) *domain.User {
	return &domain.User{
		TgID:       record.TgID,
		TgUsername: record.TgUsername,
		Name:       record.Name,
	}
}

func RegisterRequestDTOToUser(req *tgdto.RegisterRequestDTO) *domain.User {
	return &domain.User{
		TgID:       req.TgID,
		TgUsername: req.TgUsername,
		Name:       req.Name,
	}
}
