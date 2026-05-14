package botapp

import (
	"context"

	"github.com/iamHemaAI/hema-ai-bot/internal/config"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v4"
)

type App struct {
	bot *tele.Bot
	cfg *config.BotConfig
}

func New(logger *zap.Logger, cfg *config.BotConfig) (*App, error) {
	settings := tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: cfg.Poller},
	}

	b, err := tele.NewBot(settings)
	if err != nil {
		return nil, err
	}

	return &App{
		bot: b,
		cfg: cfg,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		a.bot.Stop()
	}()

	a.bot.Start()
	return nil
}
