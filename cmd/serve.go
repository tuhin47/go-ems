package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/tuhin47/go-ems/config"
	"github.com/tuhin47/go-ems/conn"
	"github.com/tuhin47/go-ems/controllers"
	dbRepo "github.com/tuhin47/go-ems/repositories/db"
	"github.com/tuhin47/go-ems/routes"
	"github.com/tuhin47/go-ems/server"
	"github.com/tuhin47/go-ems/services"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: Serve,
}

func Serve(cmd *cobra.Command, args []string) {
	dbClient := conn.Db()

	// repository
	eventRepo := dbRepo.NewEventRepositoryImpl(dbClient)

	// service
	eventSvc := services.NewEventServiceImpl(eventRepo)

	// controller
	eventCtrl := controllers.NewEventController(eventSvc)

	// Initialize the server
	echoServer := echo.New()
	server := server.New(echoServer)

	// Initialize the routes
	routes := routes.New(echoServer, eventCtrl)
	routes.Init()

	// Start the server
	server.Start(config.App().Port)
}
