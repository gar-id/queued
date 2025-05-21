/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gar-id/queued/internal/client"
	"github.com/gar-id/queued/internal/general/config"
	"github.com/gar-id/queued/internal/general/config/caches"
	"github.com/spf13/cobra"
)

// logCmd represents the logs command
var logCmd = &cobra.Command{
	Use:     "log",
	Version: caches.Version,
	Short:   "log process",
	Long:    `Use this subcommand to log process`,
	Run: func(cmd *cobra.Command, args []string) {
		fileconfig, _ := cmd.Flags().GetString("config")
		config.LoadMainConfig(fileconfig)
		processName, _ := cmd.Flags().GetString("process-name")
		client.QueuedLogs(processName)

	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	logCmd.Flags().StringP("config", "c", "", "Select your config file")
	logCmd.Flags().StringP("process-name", "n", "", "Insert process name")
}
