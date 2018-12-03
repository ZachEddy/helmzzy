package command

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	// CommandMap ties a command name to a specific implemenation of the Command
	// inteface
	CommandMap = map[string]Command{
		"delete":   delete{},
		"rollback": rollback{},
	}

	baseFuzzyFlags = []string{}
)

// Command is the basic set of behavior that all commands need to implement
type Command interface {
	Run() error
	Describe() string
}

// fuzzyReleaseList takes all the current helm releases from stdout and pipes
// it to fzf
func fuzzyReleaseList() (*exec.Cmd, error) {
	cmdReleaseList := exec.Command("helm", "list")
	outReleaseList, err := cmdReleaseList.Output()
	if err != nil {
		return nil, err
	}
	if string(outReleaseList) == "" {
		fmt.Println("No Helm releases were found ⛵")
		os.Exit(0)
	}
	lines := strings.Split(string(outReleaseList), "\n")
	cmdFuzzy := exec.Command("fzf",
		"--reverse",
		"--border",
		"--height", "50%",
		"--header", lines[0],
	)
	outReleaseList = []byte(strings.Join(lines[1:], "\n"))
	cmdFuzzy.Stdin, cmdFuzzy.Stderr = bytes.NewBuffer(outReleaseList), os.Stderr
	return cmdFuzzy, nil
}

// fuzzyRevisionHistory takes the release history for a given helm release and
// pipes it to fzf
func fuzzyRevisionHistory(release string) (*exec.Cmd, error) {
	cmdReleaseHistory := exec.Command("helm", "history", release)
	outReleaseHistory, err := cmdReleaseHistory.Output()
	if err != nil {
		return nil, err
	}
	if string(outReleaseHistory) == "" {
		fmt.Println("No Hem releases were found ⛵")
		os.Exit(0)
	}
	lines := strings.Split(string(outReleaseHistory), "\n")
	cmdFuzzy := exec.Command("fzf",
		"--reverse",
		"--border",
		"--height", "50%",
		"--header", lines[0],
	)
	outReleaseHistory = []byte(strings.Join(lines[1:], "\n"))
	cmdFuzzy.Stdin, cmdFuzzy.Stderr = bytes.NewBuffer(outReleaseHistory), os.Stderr
	return cmdFuzzy, nil
}

// getRevision gets revsion id of a given revision
func getRevision(revision string) string {
	return getReleaseName(revision)
}

// ggetReleaseName gets the name of a given release
func getReleaseName(release string) string {
	releaseFields := strings.Split(string(release), "\t")
	if len(releaseFields) <= 0 {
		return ""
	}
	releaseName := strings.TrimSpace(releaseFields[0])
	return strings.TrimSpace(releaseName)
}
