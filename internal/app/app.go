package app

import (
	"context"

	"github.com/iamHemaAI/hema-ai-bot/internal/config"
	botapp "github.com/iamHemaAI/hema-ai-bot/internal/tgbot/app"
	"go.uber.org/zap"
)

type App struct {
	diContainer *diContainer
	BOTserver   *botapp.App
}

func New(logger *zap.Logger, cfg config.Config) *App {
	a := &App{
		diContainer: newDIContainer(
			context.Background(),
			logger,
			cfg,
		),
	}

	a.initDeps()

	return a
}

func (a *App) initDeps() {
	inits := []func(){
		a.initBOTServer,
	}

	for _, fn := range inits {
		fn()
	}
}

func (a *App) initBOTServer() {
	botServer, err := botapp.New(
		a.diContainer.logger,
		&a.diContainer.cfg.Bot,
	)
	if err != nil {
		a.diContainer.logger.Fatal("init bot failed", zap.Error(err))
		return
	}

	a.BOTserver = botServer
	a.BOTserver.Setup(a.diContainer.Handlers())
}
