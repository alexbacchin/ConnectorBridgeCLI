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
var openCmd = &cobra.Command{
	Use:   "open [device id]",
	Short: "Send the Open operation to a Device",
	Long:  `Send the open operation to the device via the UDP.`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		shadeconnector.Init(host, port, apiKey)
		device_id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("device ID must me a number: %s", err)
			return
		}
		shadeconnector.Operation(device_id, int(shadeconnector.Open))
		fmt.Printf("Open device %s sucessfully", args[0])
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

}
