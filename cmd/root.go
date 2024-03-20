/*
Copyright © 2024 shiron-dev
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mp",
	Short: "MPocket is a command line tool that provides useful tools as if it were that person's pocket!",
	Long:  "MPocket is a command line tool that provides useful tools as if it were that person's pocket!",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
