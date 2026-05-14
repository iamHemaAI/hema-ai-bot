package app

import (
	"context"

	"github.com/iamHemaAI/hema-ai-bot/internal/config"
	bloodrepo "github.com/iamHemaAI/hema-ai-bot/internal/repository/db/sqlite/blood"
	userrepo "github.com/iamHemaAI/hema-ai-bot/internal/repository/db/sqlite/user"
	"github.com/iamHemaAI/hema-ai-bot/internal/services"
	botapp "github.com/iamHemaAI/hema-ai-bot/internal/tgbot/app"
	"github.com/iamHemaAI/hema-ai-bot/internal/tgbot/handlers"
	bloodreportclient "github.com/iamHemaAI/hema-ai-bot/pkg/blood_report"
	sqliteclient "github.com/iamHemaAI/hema-ai-bot/pkg/sqlite"
	"go.uber.org/zap"
)

type diContainer struct {
	ctx    context.Context
	logger *zap.Logger
	cfg    config.Config

	// infrastructure
	sqliteClient      *sqliteclient.Sqlite
	bloodReportClient *bloodreportclient.BloodReportClient

	// repositories
	userRepo  *userrepo.UserRepository
	bloodRepo *bloodrepo.BloodRepository

	// services
	userService  *services.UserService
	bloodService *services.BloodService

	// handlers
	userHandler  *handlers.UserHandler
	bloodHandler *handlers.BloodHandler
}

func newDIContainer(
	ctx context.Context,
	logger *zap.Logger,
	cfg config.Config,
) *diContainer {
	return &diContainer{
		ctx:    ctx,
		logger: logger,
		cfg:    cfg,
	}
}

func (d *diContainer) SqliteClient() *sqliteclient.Sqlite {
	if d.sqliteClient == nil {
		client, err := sqliteclient.New(d.logger, d.cfg.Sqlite)
		if err != nil {
			panic("failed to create sqlite client: " + err.Error())
		}
		d.sqliteClient = client
	}
	return d.sqliteClient
}

func (d *diContainer) BloodReportClient() *bloodreportclient.BloodReportClient {
	if d.bloodReportClient == nil {
		d.bloodReportClient = bloodreportclient.NewBloodReportClient(d.cfg.BloodReport.BaseURL)
	}
	return d.bloodReportClient
}

func (d *diContainer) UserRepo() *userrepo.UserRepository {
	if d.userRepo == nil {
		d.userRepo = userrepo.New(d.SqliteClient().DB())
	}
	return d.userRepo
}

func (d *diContainer) BloodRepo() *bloodrepo.BloodRepository {
	if d.bloodRepo == nil {
		d.bloodRepo = bloodrepo.New(d.SqliteClient().DB())
	}
	return d.bloodRepo
}

func (d *diContainer) UserService() *services.UserService {
	if d.userService == nil {
		d.userService = services.NewUserService(d.UserRepo())
	}
	return d.userService
}

func (d *diContainer) BloodService() *services.BloodService {
	if d.bloodService == nil {
		d.bloodService = services.NewBloodService(d.BloodRepo(), d.BloodReportClient())
	}
	return d.bloodService
}

func (d *diContainer) UserHandler() *handlers.UserHandler {
	if d.userHandler == nil {
		d.userHandler = handlers.NewUserHandler(d.UserService())
	}
	return d.userHandler
}

func (d *diContainer) BloodHandler() *handlers.BloodHandler {
	if d.bloodHandler == nil {
		d.bloodHandler = handlers.NewBloodHandler(d.BloodService())
	}
	return d.bloodHandler
}

func (d *diContainer) Handlers() *botapp.Handlers {
	return &botapp.Handlers{
		UserHandler:  d.UserHandler(),
		BloodHandler: d.BloodHandler(),
	}
}
