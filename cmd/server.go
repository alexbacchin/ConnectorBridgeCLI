/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/alexbacchin/ConnectorBridgeCLI/pkg/shadeconnector"
	"github.com/alexbacchin/ConnectorBridgeCLI/web"
	"github.com/spf13/cobra"
)

var serverApiKey, serverListenPort string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "server",
	Short: " API",
	Long:  `Run the REST API server to manage devices via HTTP GET `,
	Run: func(cmd *cobra.Command, args []string) {
		shadeconnector.Init(host, port, apiKey)
		if serverListenPort == "" {
			serverListenPort = "8080"
		}
		web.Init(serverApiKey, serverListenPort)
		web.Serve()
	},
}

func init() {
	serveCmd.Flags().StringVar(&serverApiKey, "server-apikey", os.Getenv("CONNECTOR_API_SERVER_APIKEY"), "The API Key to be used when calling the web server (Header X-API-Key")
	serveCmd.Flags().StringVar(&serverListenPort, "server-port", os.Getenv("CONNECTOR_API_SERVER_PORT"), "The port the server will listen. Default 8080 ")
	rootCmd.AddCommand(serveCmd)

}
