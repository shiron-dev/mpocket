/*
Copyright © 2024 shiron-dev
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "commands for git repositories",
	Long:  "commands for git repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
