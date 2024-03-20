package common

import (
	"github.com/shiron-dev/mpocket/common/data"
	"github.com/shiron-dev/mpocket/exec"
)

func GetUserData() (string, string) {
	config, _ := data.GetConfig()
	userData := config.UserData
	if userData.FromGit {
		return exec.GetGitUserData()
	} else {
		return userData.Name, userData.Email
	}
}
