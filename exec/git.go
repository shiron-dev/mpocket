package exec

import (
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
	if out, err := cmd.Output(); err != nil {
		return ""
	} else {
		return string(out)
	}
}
