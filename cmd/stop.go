/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/alexbacchin/ConnectorBridgeCLI/pkg/shadeconnector"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var stopCmd = &cobra.Command{
	Use:   "stop [device id]",
	Short: "Send the Stop operation to a Device",
	Long:  `Send the Stop operation to the device via the UDP.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		shadeconnector.Init(host, port, apiKey)
		device_id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("device ID must me a number: %s", err)
			return
		}
		shadeconnector.Operation(device_id, int(shadeconnector.Stop))
		fmt.Printf("Stop device %s sucessfully", args[0])
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

}
