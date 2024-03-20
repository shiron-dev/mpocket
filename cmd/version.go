/*
Copyright © 2024 shiron-dev
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/shiron-dev/mpocket/vars"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print MPocket version",
	Long:  "print MPocket version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("MPocket version %s - %s\n", vars.Tag, vars.CommitHash[:7])
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
