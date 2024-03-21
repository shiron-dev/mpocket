package exec

import (
	"fmt"
	"os"

	"golang.org/x/sys/execabs"
)

func GetCommitHash() string {
	cmd := execabs.Command(GetCommandName(Git), "rev-parse", "HEAD")
	if out, err := cmd.Output(); err != nil {
		return ""
	} else {
		return string(out)
	}
}

func GetTag() string {
	cmd := execabs.Command(GetCommandName(Git), "describe", "--tags", "--abbrev=0")
	return RunOutDef(cmd, "")
}

func Commit(message string, path string, args ...string) {
	cmdArgs := append([]string{"-C", path, "commit", "-m", message}, args...)
	cmd := execabs.Command(GetCommandName(Git), cmdArgs...)
	RunErr(cmd)
}

func GetGitUserData() (string, string) {
	cmd := execabs.Command(GetCommandName(Git), "config", "user.name")
	userName := RunOutFunc(cmd, func(_ error) string {
		fmt.Fprintln(os.Stderr, "error: git user.name is not set")
		os.Exit(1)
		return ""
	})

	cmd = execabs.Command(GetCommandName(Git), "config", "user.email")
	userEmail := RunOutDef(cmd, "")

	return userName, userEmail
}

func GitInit(path string) {
	cmd := execabs.Command(
		GetCommandName(Git),
		"-C", path,
		"init",
	)
	RunErr(cmd)
}
