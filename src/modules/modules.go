package modules

import (
	"context"

	"github.com/gin-gonic/gin"

	userhandler "simple-api.com/m/src/modules/users/handlers"
	userrepository "simple-api.com/m/src/modules/users/repositories"
	userusecase "simple-api.com/m/src/modules/users/usecases"
	"simple-api.com/m/src/pkg/databases/mysql"
	"simple-api.com/m/src/pkg/logger"
)

type Modules struct {
	ctx    context.Context
	router *gin.Engine
	log    logger.Logger
	db     *mysql.Mysql
}

func NewModules(ctx context.Context, router *gin.Engine, log logger.Logger, db *mysql.Mysql) *Modules {

	return &Modules{
		ctx:    ctx,
		router: router,
		log:    log,
		db:     db,
	}
}

func (m *Modules) Init() error {
	m.InitUsers()
	return nil
}

func (m *Modules) InitUsers() error {
	repository := userrepository.NewUserRepository(m.log, m.db)
	usecase := userusecase.NewUserUsecase(m.log, repository)
	handlers := userhandler.NewUserHandlers(m.log, usecase)

	group := m.router.Group("/api/users")
	handlers.UserRoutes(group)
	return nil
}
