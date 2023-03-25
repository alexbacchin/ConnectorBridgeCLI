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
var setpositionCmd = &cobra.Command{
	Use:   "set-position [device id] [position]",
	Short: "Send the postion to a Device",
	Long:  `Send the postion to the device via the UDP.`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		shadeconnector.Init(host, port, apiKey)
		device_id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("device ID must me a number: %s", err)
			return
		}
		position, err := strconv.Atoi(args[1])
		if err != nil && position >= 0 && position <= 100 {
			fmt.Printf("postion must me a number between 0 and 100: %s", err)
			return
		}
		if invert {
			position = 100 - position
		}
		message, err := shadeconnector.SetPosition(device_id, position)
		if err != nil {
			fmt.Printf("Cannot execute the operation to device:%d, error: %s", device_id, err.Error())
			return
		}
		PrintStatus(message)
	},
}

func init() {
	rootCmd.AddCommand(setpositionCmd)

}
