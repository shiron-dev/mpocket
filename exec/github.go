package exec

import (
	"strings"

	"golang.org/x/sys/execabs"
)

func ExistsRepository(repoName string) error {
	cmd := execabs.Command(GetCommandName(GitHub), "repo", "view", repoName)
	return cmd.Run()
}

func CreateRepository(repoName string, public bool, description string, sourcePath string) (string, error) {
	visibility := "--private"
	if public {
		visibility = "--public"
	}

	cmd := execabs.Command(
		GetCommandName(GitHub),
		"repo", "create", repoName,
		"-d", description,
		visibility,
		"--source="+sourcePath,
		"--remote=upstream",
	)
	out, err := cmd.Output()
	return strings.TrimSpace(string(out)), err
}
