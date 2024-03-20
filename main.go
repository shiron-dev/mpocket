/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/shiron-dev/mpocket/cmd"
	"github.com/shiron-dev/mpocket/exec"
	"github.com/shiron-dev/mpocket/vars"
)

const (
	defaultTag  = "v0.0.0 (Testing)"
	defaultHash = "hserror-xxxxxxx"
)

var (
	CommitHash = "-"
	Tag        = "-"
)

func main() {
	exec.CheckAllCommands()

	setVars()

	cmd.Execute()
}

func setVars() {
	if CommitHash == "-" {
		if h := exec.GetCommitHash(); h != "" {
			CommitHash = h
		} else {
			CommitHash = defaultHash
		}
	}
	if Tag == "-" {
		if t := exec.GetTag(); t != "" {
			Tag = t
		} else {
			Tag = defaultTag
		}
	}

	vars.CommitHash = CommitHash
	vars.Tag = Tag
}
