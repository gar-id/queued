/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/gar-id/queued/internal/client"
	"github.com/gar-id/queued/internal/general/config"
	"github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/tools"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Version: caches.Version,
	Short:   "start group/program/process",
	Long:    `Use this subcommand to start group/program/process`,
	Run: func(cmd *cobra.Command, args []string) {
		fileconfig, _ := cmd.Flags().GetString("config")
		config.LoadMainConfig(fileconfig)
		var groupName, programName, processName string
		var groupNameArray, programNameArray, processNameArray []string
		groupName, _ = cmd.Flags().GetString("group-name")
		if groupName == "" {
			groupNameArray = nil
		} else {
			groupNameArray = strings.Split(groupName, ",")
		}
		programName, _ = cmd.Flags().GetString("program-name")
		if programName == "" {
			programNameArray = nil
		} else {
			programNameArray = strings.Split(programName, ",")
		}
		processName, _ = cmd.Flags().GetString("process-name")
		if processName == "" {
			processNameArray = nil
		} else {
			processNameArray = strings.Split(processName, ",")
		}
		if len(groupNameArray) == 0 && len(programNameArray) == 0 && len(processNameArray) == 0 {
			tools.ZapLogger("console").Info("Please insert group-name, program-name or / and process-name")
			return
		}
		client.QueuedAction(groupNameArray, programNameArray, processNameArray, "start")

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().StringP("config", "c", "", "Select your config file")
	startCmd.Flags().StringP("group-name", "g", "", "Insert group name")
	startCmd.Flags().StringP("program-name", "p", "", "Insert program name")
	startCmd.Flags().StringP("process-name", "n", "", "Insert process name")
}
