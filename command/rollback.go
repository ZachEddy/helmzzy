package command

import (
	"os"
	"os/exec"
)

// Rollback is a command that will roll a Helm release back to a targeted
// version
type rollback struct{}

// Run will rollback a Helm release via the command line
func (rollback) Run() error {
	cmdFuzzyReleaseList, err := fuzzyReleaseList()
	if err != nil {
		return err
	}
	outFuzzyReleaseList, err := cmdFuzzyReleaseList.Output()
	if err != nil {
		return err
	}
	release := getReleaseName(string(outFuzzyReleaseList))
	cmdFuzzyReleaseHistory, err := fuzzyRevisionHistory(release)
	if err != nil {
		return err
	}
	outFuzzyReleaseHistory, err := cmdFuzzyReleaseHistory.Output()
	if err != nil {
		return err
	}
	rollbackVersion := getRevision(string(outFuzzyReleaseHistory))
	cmdReleaseDelete := exec.Command("helm", "rollback", release, rollbackVersion)
	cmdReleaseDelete.Stdout, cmdReleaseDelete.Stderr = os.Stdout, os.Stderr
	if err := cmdReleaseDelete.Run(); err != nil {
		return err
	}
	return nil
}

// Describe provides a usage message for the rollback command
func (rollback) Describe() string {
	return "rollback a Helm release by searching for a target revision"
}
