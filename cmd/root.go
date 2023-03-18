package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/alexbacchin/ConnectorBridgeAPI/pkg/shadeconnector"
)

var version = "0.0.1"
var host, port, apiKey string

var rootCmd = &cobra.Command{
	Use:     "sconnector-cli",
	Version: version, //-ldflags="-X 'github.com/ThorstenHans/stringer/cmd/stringer.version=0.0.2'"
	Short:   "ShadeConnector command line utility",
	Long: `A command line utility to operate Dooyaa ShadeConnector blinds and other products
				  https://github.com/alexbacchin/ShadeConnectorAPI/`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SharedConnector CLI")
	},
}

func init() {
	port = "31200"
	rootCmd.PersistentFlags().StringVar(&host, "host", os.Getenv("CONNECTOR_BRIDGE_HOST"), "The hostname of IP address of the Connector Bridge")
	rootCmd.PersistentFlags().StringVar(&port, "port", os.Getenv("CONNECTOR_BRIDGE_PORT"), "The port for the Connector Bridge connection")
	rootCmd.PersistentFlags().StringVar(&apiKey, "apikey", os.Getenv("CONNECTOR_BRIDGE_APIKEY"), "The ApiKey from Connector Bridge. On the mobile app: Go to Settings (gear), About. Tap 5 times on the screen")
	shadeconnector.Init(host, port, apiKey)
	fmt.Println(host)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
