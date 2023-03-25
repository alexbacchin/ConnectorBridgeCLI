/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/alexbacchin/ConnectorBridgeCLI/pkg/shadeconnector"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var statusCmd = &cobra.Command{
	Use:   "status [device id]",
	Short: "Returns device status",
	Long:  `Returns device status, including position, batery, etc.`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		shadeconnector.Init(host, port, apiKey)
		device_id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("device ID must me a number: %s", err)
			return
		}
		message, err := shadeconnector.QueryStatus(device_id)
		if err != nil {
			fmt.Printf("Cannot execute the operation to device:%d, error: %s", device_id, err.Error())
			return
		}
		PrintStatus(message)

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func PrintStatus(message *shadeconnector.DeviceStatus) {
	if simple_status {
		position := message.CurrentPosition
		if invert {
			fmt.Println(100 - position)
			return
		} else {
			fmt.Println(position)
			return
		}
	}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Cannot covert message: %s", err.Error())
		return
	}
	fmt.Println(string(output))
}
