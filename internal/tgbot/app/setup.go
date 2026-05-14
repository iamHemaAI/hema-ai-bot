package botapp

import tele "gopkg.in/telebot.v4"

type Handlers struct {
	BloodHandler
	UserHandler
}

type BloodHandler interface {
	CreateBloodAnalyze(c tele.Context) error
	BloodQuiz(c tele.Context) error
}

type UserHandler interface {
	Register(c tele.Context) error
}

func (a *App) Setup(h *Handlers) {
	a.bot.Handle("/start", h.UserHandler.Register)
	a.bot.Handle(analyzeButtonText, h.BloodHandler.CreateBloodAnalyze)
	a.bot.Handle(tele.OnText, h.BloodHandler.BloodQuiz)
}

const analyzeButtonText = "Анализировать"
