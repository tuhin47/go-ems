package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/tuhin47/go-ems/conn"
	"github.com/tuhin47/go-ems/controllers"
	"github.com/tuhin47/go-ems/middlewares"
	db_repo "github.com/tuhin47/go-ems/repositories/db"
	"github.com/tuhin47/go-ems/routes"
	"github.com/tuhin47/go-ems/server"
	"github.com/tuhin47/go-ems/services"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// clients
	dbClient := conn.Db()
	redisClient := conn.Redis()
	emailClient := conn.EmailClient()
	asynqClient := conn.Asynq()
	asynqInspector := conn.AsynqInspector()

	// repositories
	dbRepo := db_repo.NewRepository(dbClient)
	asynqRepo := asynq_repo.NewRepository(config.Asynq(), asynqClient, asynqInspector)
	mailRepo := mail_repo.NewRepository(emailClient, config.Email())

	// services
	redisSvc := services.NewRedisService(redisClient)
	eventSvc := services.NewEventServiceImpl(dbRepo, dbRepo)
	userSvc := services.NewUserServiceImpl(redisSvc, dbRepo)
	tokenSvc := services.NewTokenServiceImpl(redisSvc)
	authSvc := services.NewAuthServiceImpl(userSvc, tokenSvc)
	mailSvc := services.NewMailService(dbRepo, dbRepo, mailRepo)
	asynqSvc := services.NewAsynqService(config.Asynq(), asynqRepo, dbRepo, dbRepo)

	// controllers
	eventCtrl := controllers.NewEventController(eventSvc, mailSvc, asynqSvc)
	userCtrl := controllers.NewUserController(userSvc)
	authCtrl := controllers.NewAuthController(authSvc)

	// middlewares
	authMiddleware := middlewares.NewAuthMiddleware(authSvc, userSvc)

	// Server
	var echo_ = echo.New()
	var Routes = routes.New(echo_, eventCtrl, userCtrl, authCtrl, authMiddleware)
	var Server = server.New(echo_)

	// Spooling
	Routes.Init()

	// Stopping running workers
	Server.Start()
}
