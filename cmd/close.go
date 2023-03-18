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
	Use:   "close [device_id]",
	Short: "Send the close operation to a Device",
	Long:  `Send the close operation to the device via the UDP.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		device_id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("device ID must me a number: %s", err)
			return
		}
		shadeconnector.Operation(device_id, int(shadeconnector.Close))
		fmt.Printf("Close device %s sucessfully", args[0])

	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
