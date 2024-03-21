/*
Copyright © 2024 shiron-dev
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/shiron-dev/mpocket/common"
	"github.com/shiron-dev/mpocket/exec"
)

const (
	initCommitMessage = "Initial commit"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <repository name> [<exts>]",
	Short: "Create a new repository",
	Long: `Create a new repository.
	
This command is a simple repository creation tool that configures settings based on git commands, github commands and gitignore.io.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "error: missing repository name")
			cmd.Usage()
			os.Exit(1)
		}

		repoName := args[0]
		exts := args[1:]
		public, _ := cmd.Flags().GetBool("public")
		description, _ := cmd.Flags().GetString("description")
		license, _ := cmd.Flags().GetString("license")

		licenseEnum, err := common.AssertLicense(license)
		if license != "" && err != nil {
			fmt.Fprintln(os.Stderr, `error: invalid value for license: "`+license+`"
The following licenses are currently available.
Please contact "https://github.com/shiron-dev/mpocket/issues" if you have any problems.
`)
			for _, l := range common.OkLicenseList {
				fmt.Fprintln(os.Stderr, "- "+l)
			}
			os.Exit(1)
		}

		createRepository(repoName, exts, public, description, licenseEnum)
	},
}

func init() {

	createCmd.Flags().BoolP("public", "p", false, "Is public")
	createCmd.Flags().StringP("description", "d", "", "Description")
	createCmd.Flags().StringP("license", "l", "", "License")

	repoCmd.AddCommand(createCmd)
}

func createRepository(repoName string, exts []string, public bool, description string, license int) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	repoPath := filepath.Join(dir, repoName)

	// Check if the repository already exists
	err = exec.ExistsRepository(repoName)
	if err == nil {
		fmt.Println("error: (remote) repository", repoName, "already exists")
		os.Exit(1)
	}

	// Check if the repository directory already exists
	_, err = os.Stat(repoPath)
	if err == nil {
		fmt.Fprintln(os.Stderr, "error: (local) repository directory", repoPath, "already exists")
		os.Exit(1)
	}

	// Make a new repository directory
	err = os.MkdirAll(repoPath, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Create git repository
	exec.GitInit(repoPath)

	out, err := exec.CreateRepository(repoName, public, description, repoPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		fmt.Println(out)
	}

	exec.Commit(initCommitMessage, repoPath, "--allow-empty")

	gitignore := ""
	for _, ext := range exts {
		gitignore += common.GetGitignoreIo(ext)
		gitignore += "\n"
	}
	common.CreateFile(filepath.Join(repoPath, ".gitignore"), gitignore)

	if license != -1 {
		licenseMgs := common.GenLicense(license)
		common.CreateFile(filepath.Join(repoPath, "LICENSE"), licenseMgs)
	}
}
