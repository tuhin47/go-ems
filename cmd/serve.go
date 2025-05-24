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

	// repositories
	dbRepo := db_repo.NewRepository(dbClient)

	// services
	redisSvc := services.NewRedisService(redisClient)
	eventSvc := services.NewEventServiceImpl(dbRepo)
	userSvc := services.NewUserServiceImpl(redisSvc, dbRepo)
	tokenSvc := services.NewTokenServiceImpl(redisSvc)
	authSvc := services.NewAuthServiceImpl(userSvc, tokenSvc)

	// controllers
	eventCtrl := controllers.NewEventController(eventSvc)
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
	Server.Start()
}
