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

// closeCmd represents the close command
var closeCmd = &cobra.Command{
	Use:   "close [device id]",
	Short: "Send the close operation to a Device",
	Long:  `Send the close operation to the device via the UDP.`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		shadeconnector.Init(host, port, apiKey)
		device_id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("device ID must me a number: %s", err)
			return
		}
		shadeconnector.Operation(device_id, int(shadeconnector.Close))
		fmt.Printf("Close device %s sucessfully", args[0])
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
