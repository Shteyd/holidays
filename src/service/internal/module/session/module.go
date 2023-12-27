package session

import (
	"github.com/Shteyd/holidays/src/service/internal/module/session/adapter/token"
	getListHandler "github.com/Shteyd/holidays/src/service/internal/module/session/controller/getList"
	"github.com/Shteyd/holidays/src/service/internal/module/session/storage/postgres"
	"github.com/Shteyd/holidays/src/service/internal/module/session/usecase"
	"github.com/indigo-web/indigo/router/inbuilt"
)

type Module struct {
	Usecase  usecase.SessionUsecase
	Handlers Handlers
}

type Handlers struct {
	GetList inbuilt.Handler
}

func New(opts ...Option) Module {
	var options options
	for _, opt := range opts {
		opt(&options)
	}

	sessionUsecase := usecase.New(
		options.logger,
		token.New(options.tokenLength),
		postgres.New(options.logger, options.database),
	)

	return Module{
		Usecase: sessionUsecase,
		Handlers: Handlers{
			GetList: getListHandler.New(
				sessionUsecase,
				options.timeout,
			),
		},
	}
}
