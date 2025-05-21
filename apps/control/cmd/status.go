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

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Version: caches.Version,
	Short:   "status group/program/process",
	Long:    `Use this subcommand to status group/program/process`,
	Run: func(cmd *cobra.Command, args []string) {
		fileconfig, _ := cmd.Flags().GetString("config")
		config.LoadMainConfig(fileconfig)

		client.QueuedStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	statusCmd.Flags().StringP("config", "c", "", "Select your config file")
}
