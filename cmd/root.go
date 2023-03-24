package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var host, port, apiKey string

var rootCmd = &cobra.Command{
	Use:     "sconnector-cli",
	Version: version,
	Short:   "ShadeConnector command line utility",
	Long: `A command line utility to operate ShadeConnector blinds and other products
				  https://github.com/alexbacchin/ShadeConnectorCLI/`,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&host, "host", os.Getenv("CONNECTOR_BRIDGE_HOST"), "The hostname of IP address of the Connector Bridge. Default 238.0.0.18")
	rootCmd.PersistentFlags().StringVar(&port, "port", os.Getenv("CONNECTOR_BRIDGE_PORT"), "The port for the Connector Bridge connection. Default 32100")
	rootCmd.PersistentFlags().StringVar(&apiKey, "apikey", os.Getenv("CONNECTOR_BRIDGE_APIKEY"), "The ApiKey from Connector Bridge. On the mobile app: Go to Settings (gear), About. Tap 5 times on the screen")
}

func Execute() {
	if port == "" {
		port = "32100"
	}
	if host == "" {
		host = "238.0.0.18"
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
