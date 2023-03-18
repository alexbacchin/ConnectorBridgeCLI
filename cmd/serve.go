/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/alexbacchin/ConnectorBridgeCLI/web"
	"github.com/spf13/cobra"
)

var serverApiKey, serverListenPort string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve API",
	Long:  `Run the REST API server to manage devices`,
	Run: func(cmd *cobra.Command, args []string) {
		web.Init(serverApiKey, serverListenPort)
		web.Serve()
	},
}

func init() {
	serverListenPort = "8080"
	serveCmd.Flags().StringVar(&serverApiKey, "server-apikey", os.Getenv("API_SERVER_APIKEY"), "The API Key to be used when calling the web server (Header X-API-Key")
	serveCmd.Flags().StringVar(&serverListenPort, "server-port", os.Getenv("API_SERVER_PORT"), "The port the server will listen. Default 8080 ")
	rootCmd.AddCommand(serveCmd)

}
