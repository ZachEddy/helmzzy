package command

import (
	"os"
	"os/exec"
)

// Delete is a command that will delete a Helm release
type delete struct{}

// Run will delete a Helm release via the command line
func (delete) Run() error {
	cmdFuzzy, err := fuzzyReleaseList()
	if err != nil {
		return err
	}
	outFuzzy, err := cmdFuzzy.Output()
	if err != nil {
		return err
	}
	release := getReleaseName(string(outFuzzy))
	cmdReleaseDelete := exec.Command("helm", "delete", release)
	cmdReleaseDelete.Stdout = os.Stdout
	if err := cmdReleaseDelete.Run(); err != nil {
		return err
	}
	return nil
}

// Describe provides a usage message for the delete command
func (delete) Describe() string {
	return "delete a specific Helm release"
}
