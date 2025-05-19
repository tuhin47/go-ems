package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/vivasoft-ltd/go-ems/config"
	"github.com/vivasoft-ltd/go-ems/conn"
	"github.com/vivasoft-ltd/go-ems/controllers"
	dbRepo "github.com/vivasoft-ltd/go-ems/repositories/db"
	"github.com/vivasoft-ltd/go-ems/routes"
	"github.com/vivasoft-ltd/go-ems/server"
	"github.com/vivasoft-ltd/go-ems/services"
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
