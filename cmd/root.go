package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tuhin47/go-ems/config"
	"github.com/tuhin47/go-ems/conn"
	//"github.com/tuhin47/golang-course-utils/logger"
)

var rootCmd = &cobra.Command{
	Use: "root",
}

func Execute() {
	// Load the configuration
	config.LoadConfig()

	// Initialize the logger
	initLogger()

	// Initialize the database connection
	conn.ConnectDB()

	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func initLogger() {
	//logger.SetFileLogger(config.Logger().FilePath)
}
