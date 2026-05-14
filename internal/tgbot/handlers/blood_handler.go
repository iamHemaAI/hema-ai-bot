package handlers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/iamHemaAI/hema-ai-bot/internal/domain"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain/mappers"
	"github.com/iamHemaAI/hema-ai-bot/internal/tgbot/tgdto"
	"github.com/iamHemaAI/hema-ai-bot/pkg/apperrors"
	tele "gopkg.in/telebot.v4"
)

type bloodService interface {
	CreateBloodAnalyze(ctx context.Context, blood *domain.Blood) ([]byte, error)
}

type BloodHandler struct {
	s     bloodService
	mu    sync.Mutex
	state map[int64]*domain.QuizState
}

func NewBloodHandler(s bloodService) *BloodHandler {
	return &BloodHandler{
		s:     s,
		state: make(map[int64]*domain.QuizState),
	}
}

// CreateBloodAnalyze — reply кнопка analyzeButtonText, начинает квиз
func (h *BloodHandler) CreateBloodAnalyze(c tele.Context) error {
	userID := c.Sender().ID

	h.mu.Lock()
	h.state[userID] = &domain.QuizState{Step: 0, DTO: tgdto.BloodDTO{UserID: userID}}
	h.mu.Unlock()

	return c.Send(domain.QuizSteps[0].Prompt)
}

// BloodQuiz — обработчик OnText, ведёт квиз по шагам
// тут мы проводим квиз для выявления каждого параметра крови (особо не валидируем)
// квиз по отдельности сообщения
// просто говорим ему что вводить
func (h *BloodHandler) BloodQuiz(c tele.Context) error {
	userID := c.Sender().ID

	h.mu.Lock()
	st, active := h.state[userID]
	h.mu.Unlock()

	if !active {
		return nil
	}

	if err := domain.QuizSteps[st.Step].Parse(st, c.Text()); err != nil {
		return c.Send(fmt.Sprintf("%v, попробуйте ещё раз.\n%s", err, domain.QuizSteps[st.Step].Prompt))
	}

	st.Step++

	if st.Step < len(domain.QuizSteps) {
		return c.Send(domain.QuizSteps[st.Step].Prompt)
	}

	// квиз завершён — создаём анализ
	h.mu.Lock()
	delete(h.state, userID)
	h.mu.Unlock()

	blood, err := mappers.BloodDTOtoBlood(&st.DTO)
	if err != nil {
		if bErr, ok := errors.AsType[*apperrors.DomainValidationError](err); ok {
			return c.Send(fmt.Sprintf("Твои данные нереалистичны: %s", bErr.Error()))
		}

		return c.Send("Не могу обработать твои данные")
	}

	pdf, err := h.s.CreateBloodAnalyze(context.Background(), blood)
	if err != nil {
		return c.Send("Ошибка при создании анализа крови")
	}

	pdfName := fmt.Sprintf(
		"report_%s_%s.pdf",
		c.Sender().Username,
		time.Now().Format("2006-01-02_15-04"),
	)

	return c.Send(&tele.Document{
		File:     tele.FromReader(bytes.NewReader(pdf)),
		FileName: pdfName,
		MIME:     "application/pdf",
	})
}
