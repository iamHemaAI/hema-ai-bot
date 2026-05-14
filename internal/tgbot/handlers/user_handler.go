package handlers

import (
	"context"
	"errors"

	"github.com/iamHemaAI/hema-ai-bot/internal/domain"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain/mappers"
	"github.com/iamHemaAI/hema-ai-bot/internal/tgbot/tgdto"
	"github.com/iamHemaAI/hema-ai-bot/pkg/apperrors"
	tele "gopkg.in/telebot.v4"
)

type userService interface {
	Register(ctx context.Context, user *domain.User) error
}

type UserHandler struct {
	s userService
}

func NewUserHandler(s userService) *UserHandler {
	return &UserHandler{
		s: s,
	}
}

func (h *UserHandler) Register(c tele.Context) error {
	req := tgdto.RegisterRequestDTO{
		TgID:       c.Sender().ID,
		TgUsername: c.Sender().Username,
		Name:       c.Sender().FirstName,
	}

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	menu.Reply(menu.Row(menu.Text("Анализировать")))

	if err := h.s.Register(
		context.Background(),
		mappers.RegisterRequestDTOToUser(&req),
	); err != nil {
		if errors.Is(err, apperrors.ErrUserAlreadyExists) {
			return c.Send("Ты уже зарегистрирован. Нажми \"Анализировать\", чтобы начать", menu)
		}

		return c.Send("Ошибка регистрации")
	}

	_ = c.Send(&tele.Sticker{File: tele.File{FileID: registerSuccessSticker}})

	return c.Send("Приветствую в нашей AI лаборатории! Нажми \"Анализировать\", чтобы начать 👽", menu)
}
